package rpc

import (
	"context"
	"strings"

	"github.com/twitchtv/twirp"

	kitlog "github.com/go-kit/kit/log"
	devicereg "github.com/thingful/twirp-devicereg-go"

	"github.com/thingful/iotdevicereg/pkg/postgres"
)

type deviceRegImpl struct {
	logger  kitlog.Logger
	db      postgres.DB
	verbose bool
}

// NewDeviceReg constructs a new DeviceRegistration instance. We pass in the
// components needed for this component to operate.
func NewDeviceReg(db postgres.DB, logger kitlog.Logger) devicereg.DeviceRegistration {
	logger = kitlog.With(logger, "module", "rpc")
	logger.Log("msg", "creating devicereg")

	return &deviceRegImpl{
		db:      db,
		logger:  logger,
		verbose: true,
	}
}

// Start starts the service, doing any work that needs to be done
func (d *deviceRegImpl) Start() error {
	d.logger.Log("msg", "starting devicereg")
	return nil
}

// Stop stops the service, cleaning up anything that needs to be cleaned up.
func (d *deviceRegImpl) Stop() error {
	d.logger.Log("msg", "stopping devicereg")
	return nil
}

// ClaimDevice is our implementation of the ClaimDevice method defined for our
// Twirp service. As a result of this call the system as a whole should have
// created some key pairs for the user and the device, and created a new
// encrypted stream for the device on the encoder. THis service will maintain a
// store of the generated keys.
func (d *deviceRegImpl) ClaimDevice(ctx context.Context, req *devicereg.ClaimDeviceRequest) (*devicereg.ClaimDeviceResponse, error) {
	device, err := createValidDevice(req)
	if err != nil {
		return nil, err
	}

	device, err = d.db.RegisterDevice(device)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return &devicereg.ClaimDeviceResponse{
		UserPrivateKey:  device.User.PrivateKey,
		UserPublicKey:   device.User.PublicKey,
		DevicePublicKey: device.PublicKey,
	}, nil
}

// RevokeDevie is our implementation of the method defined on the
// DeviceRegistration servie interface. Calling this causes the specified device
// to be deleted, and the system will also call down to the encoder in order to
// delete associated streams.
func (d *deviceRegImpl) RevokeDevice(ctx context.Context, req *devicereg.RevokeDeviceRequest) (*devicereg.RevokeDeviceResponse, error) {
	err := validateRevokeRequest(req)
	if err != nil {
		return nil, err
	}

	err = d.db.DeleteDevice(req.DeviceToken, req.UserPublicKey)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return &devicereg.RevokeDeviceResponse{}, nil
}

// createValidDevice both validates the incoming request, and returns an
// instantiated Device object ready for saving.
func createValidDevice(req *devicereg.ClaimDeviceRequest) (*postgres.Device, error) {
	if req.DeviceToken == "" {
		return nil, twirp.RequiredArgumentError("device_token")
	}

	if req.UserUid == "" {
		return nil, twirp.RequiredArgumentError("user_uid")
	}

	if req.Location == nil {
		return nil, twirp.RequiredArgumentError("location")
	}

	if req.Location.Longitude == 0 {
		return nil, twirp.RequiredArgumentError("longitude")
	}

	if req.Location.Latitude == 0 {
		return nil, twirp.RequiredArgumentError("latitude")
	}

	if req.Location.Longitude < -180 || req.Location.Longitude > 180 {
		return nil, twirp.InternalError("longitude must be between -180 and 180 degrees")
	}

	if req.Location.Latitude < -90 || req.Location.Latitude > 90 {
		return nil, twirp.InternalError("latitude must be between -90 and 90 degrees")
	}

	return &postgres.Device{
		Token:       req.DeviceToken,
		Longitude:   req.Location.Longitude,
		Latitude:    req.Location.Latitude,
		Disposition: strings.ToLower(req.Disposition.String()),
		User: &postgres.User{
			UID: req.UserUid,
		},
	}, nil
}

// validateRevokeRequest validates the incoming request, returning an error if
// any required fields are missing.
func validateRevokeRequest(req *devicereg.RevokeDeviceRequest) error {
	if req.DeviceToken == "" {
		return twirp.RequiredArgumentError("device_token")
	}

	if req.UserPublicKey == "" {
		return twirp.RequiredArgumentError("user_public_key")
	}

	return nil
}
