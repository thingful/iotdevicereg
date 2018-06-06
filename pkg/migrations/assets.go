// Code generated by go-bindata. DO NOT EDIT.
// sources:
// sql/20180606143811_add_users.down.sql
// sql/20180606143811_add_users.up.sql
package migrations

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __20180606143811_add_usersDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\xf0\xf4\x73\x71\x8d\x50\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x2d\x4e\x2d\x2a\x8e\x2f\xcd\x4c\x89\xcf\x4c\x29\x56\x70\x76\x0c\x76\x76\x74\x71\xb5\xe6\xe2\x02\x2b\x0e\x71\x74\xf2\x71\x45\x57\x8c\xae\xc8\x35\x22\xc4\xd5\x2f\xd8\xd3\xdf\x0f\x49\x61\x41\x7a\x72\x51\x65\x41\x49\xbe\x35\x20\x00\x00\xff\xff\xac\xc3\x03\x6d\x74\x00\x00\x00")

func _20180606143811_add_usersDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__20180606143811_add_usersDownSql,
		"20180606143811_add_users.down.sql",
	)
}

func _20180606143811_add_usersDownSql() (*asset, error) {
	bytes, err := _20180606143811_add_usersDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20180606143811_add_users.down.sql", size: 116, mode: os.FileMode(420), modTime: time.Unix(1528296183, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __20180606143811_add_usersUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\x41\x4e\xc3\x30\x10\x45\xf7\x3e\xc5\x5f\x36\x12\x37\xe8\xca\xa5\x53\x31\xc2\x71\x4a\x3c\x51\x13\x36\x51\x88\x2d\x64\x81\x44\x94\xda\x88\xde\x1e\x25\x02\x21\xa0\xcb\x99\xff\xde\xe2\xdd\xd6\xa4\x85\x40\xad\x90\x75\x5c\x59\xf0\x01\xb6\x12\x50\xcb\x4e\x1c\xa6\xe7\x71\xbe\x4c\xe9\x6d\xab\xd4\x17\x29\x7a\x67\xe8\x0f\x95\xcf\x61\x3e\x63\xa3\x80\xe8\xe1\xa8\x66\x6d\x70\xac\xb9\xd4\x75\x87\x7b\xea\x6e\x14\x90\xa3\x87\x50\x2b\xab\x66\x1b\x63\x96\xe7\x34\xc7\xf7\x21\x85\xfe\x25\x5c\xb0\xeb\x84\xf4\xef\x35\x3f\xbd\xc6\x71\x1d\xff\x99\xe3\x1c\x86\x14\x7c\x3f\x24\x08\x97\xe4\x44\x97\x47\x9c\x58\xee\xd6\x13\x8f\x95\x25\xec\xe9\xa0\x1b\xb3\x78\xa7\x4d\xa1\x8a\x9f\x84\xc6\xf2\x43\x43\x60\xbb\xa7\xf6\x5a\x49\x9f\xa3\xef\xa3\xff\x50\x40\x65\xbf\xe3\x72\xf4\xc5\xf6\x33\x00\x00\xff\xff\x22\x68\x0f\x28\x2e\x01\x00\x00")

func _20180606143811_add_usersUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__20180606143811_add_usersUpSql,
		"20180606143811_add_users.up.sql",
	)
}

func _20180606143811_add_usersUpSql() (*asset, error) {
	bytes, err := _20180606143811_add_usersUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20180606143811_add_users.up.sql", size: 302, mode: os.FileMode(420), modTime: time.Unix(1528296951, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"20180606143811_add_users.down.sql": _20180606143811_add_usersDownSql,
	"20180606143811_add_users.up.sql": _20180606143811_add_usersUpSql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"20180606143811_add_users.down.sql": &bintree{_20180606143811_add_usersDownSql, map[string]*bintree{}},
	"20180606143811_add_users.up.sql": &bintree{_20180606143811_add_usersUpSql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

