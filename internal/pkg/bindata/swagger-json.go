// Code generated by go-bindata. DO NOT EDIT.
// sources:
// api/api.swagger.json (2.638kB)

package bindata

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _apiApiSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x55\x4d\x73\xd3\x30\x10\xbd\xfb\x57\xec\x08\x8e\x9d\x3a\x29\x81\x43\x6e\x30\x43\xa1\x17\x0e\xa5\x70\x61\x3a\x8c\x62\xad\x13\x75\x6c\x49\x5d\xad\x03\x81\xf1\x7f\x67\xe4\x38\xf1\x47\x9c\x3a\x43\x19\xe2\x9b\xa5\xd5\xdb\xb7\x4f\xab\x7d\xbf\x23\x00\xe1\x7f\xc8\xe5\x12\x49\xcc\x41\x5c\x5d\x4e\xc4\x45\x58\xd3\x26\xb5\x62\x0e\x61\x1f\x40\xb0\xe6\x0c\xc3\xfe\x07\x0b\x1e\x69\xad\x13\x04\xc6\xdc\x65\x92\xb1\x8a\x07\x10\x0a\x7d\x42\xda\xb1\xb6\x26\x44\x7a\x96\xc4\x48\x07\x61\x6b\x24\x5f\x87\x4c\x2f\x27\x22\x02\x28\xab\x84\x2c\x97\x5e\xcc\xe1\x5b\x15\xb5\x4d\x0b\x20\x8c\xcc\xeb\xbc\x9f\xb7\x69\xef\x6a\xb8\xfa\xf7\xeb\x54\x54\xa1\x65\x04\x70\x5f\xe1\x24\xd6\xf8\x22\xc7\x06\x4b\x48\xe7\x32\x9d\xc8\x40\x2c\x7e\xf0\xd6\x88\x7d\xac\x23\xab\x8a\xe4\xc4\x58\xc9\x2b\xdf\x28\x12\x4b\xa7\xe3\xf5\x34\x46\xc3\x9a\x37\x31\xa3\xe7\xfd\x66\x88\xb6\x9d\x7f\x00\x61\x1d\x52\x85\x7b\xa3\x9e\x2e\xe8\xfb\x1d\x7a\xfe\x28\x8d\xca\x90\x6a\xd1\x2a\x00\x42\xef\xac\xf1\xe8\x3b\xb8\x00\xe2\x6a\x32\xe9\x2d\x1d\xde\xc6\x5b\xf0\x45\x92\xa0\xf7\x69\x91\xc1\x0e\xe9\xb2\x05\x5f\x1d\xf2\xc9\x0a\x73\x79\x00\x06\x20\x5e\x12\xa6\x01\xe7\x45\xac\x30\xd5\x46\x07\x5c\x1f\xaf\xa7\x2d\xae\xb7\x35\xaa\xe8\x9c\x2d\x5b\x7f\x65\x3b\x9d\x98\x9d\x40\xfb\x9d\x54\x70\x8b\x8f\x45\x50\xf7\xd9\x5c\xdf\x13\xd9\xbf\x60\x39\x1d\x65\xf9\xc5\xc8\x82\x57\x96\xf4\x2f\x54\x67\xa3\xf9\x6a\x94\xe6\xb5\xa5\x85\x56\x0a\xcd\xd9\x38\xce\x46\x39\x7e\xb2\x0c\xd7\xb6\x30\xe7\xd2\xf1\xf5\x09\x4d\x79\x63\x18\xc9\xc8\x0c\xc2\x93\x45\x82\x2a\xd1\x99\xf8\x2a\x4c\x65\x91\xf1\x09\x2d\x8a\x3f\x1d\x26\x8c\xea\x7f\xd3\x8d\x06\x88\x0b\x27\x49\xe6\xc8\x48\xcd\xe8\xdd\x7e\xbd\x2a\x76\xe3\x7f\x61\xd5\xa6\x4f\x59\x9b\x63\x3b\x84\x8f\x85\x26\x0c\x73\x96\xa9\xc0\x7f\x3b\xe4\xb6\xc3\xe8\x84\x7a\xef\x5b\xf5\x76\xcc\xad\x5e\x1b\xb3\xb4\x0a\x23\x6a\xa3\x96\x7b\xb3\x6c\xf1\x6b\x2c\xa9\xbe\x93\xb6\x0d\xf1\xc6\x55\xfa\xd9\xc5\x03\x26\xcd\x08\x0d\xc6\xe7\x90\x58\xf7\xcc\x44\x24\x56\x61\xdf\x5e\x76\x18\x9e\x49\x9b\x65\x47\x6b\x91\x5a\xca\x65\xe8\x3e\xa1\x0d\xbf\x99\x89\xc1\xcb\xce\xd1\x7b\xb9\x1c\xc3\x1d\x3c\xaa\x90\xa5\xce\x0e\x1c\xaf\x57\x56\xd4\x17\xbf\x96\xeb\xa2\xa3\xcb\xbe\x57\x9f\xa1\x8f\x92\xdc\xef\x9c\x63\x18\x03\x38\xc3\x6f\x01\x7b\x97\x06\xa3\xaf\x6e\xbc\xe0\x81\x8e\x3d\x5e\xf5\x53\x67\x47\x25\x6b\x1a\x33\x2a\xa3\x3f\x01\x00\x00\xff\xff\x3b\x6e\x32\x98\x4e\x0a\x00\x00")

func apiApiSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_apiApiSwaggerJson,
		"api/api.swagger.json",
	)
}

func apiApiSwaggerJson() (*asset, error) {
	bytes, err := apiApiSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api/api.swagger.json", size: 2638, mode: os.FileMode(0644), modTime: time.Unix(1729661436, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x86, 0xa1, 0xd0, 0x81, 0xcc, 0x1a, 0x2, 0xcc, 0x7a, 0x12, 0xe0, 0x5f, 0xf0, 0xdc, 0xa3, 0xc3, 0xbe, 0xb0, 0xe4, 0xf2, 0xb3, 0x4c, 0x6c, 0x79, 0xac, 0x3a, 0x54, 0x4b, 0x1c, 0x6, 0x44, 0x34}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"api/api.swagger.json": apiApiSwaggerJson,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"api": {nil, map[string]*bintree{
		"api.swagger.json": {apiApiSwaggerJson, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	err = os.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}