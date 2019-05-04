// Code generated by go-bindata.
// sources:
// client_secret.json
// DO NOT EDIT!

package main

import (
	"github.com/elazarl/go-bindata-assetfs"
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

var _client_secretJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\xd0\xcd\x6a\xe3\x30\x14\x05\xe0\x77\xd1\x7a\x62\xc7\x9e\x24\xd6\x68\x39\x90\xae\x4a\x7f\xa0\xd0\x45\x29\x46\x91\x6f\x6d\xa5\x8a\xa4\xe8\x5e\xd9\x0d\x21\xef\x5e\x29\x3f\xd0\x6d\x37\x06\xc3\x39\x9f\x8f\xef\x91\x69\x8b\x24\x8d\x81\x8e\x89\x23\x53\x46\x83\xa5\x56\xa7\x17\xc6\xf9\xdf\x15\xaf\xf9\x62\xd1\xf0\xe5\x0c\x11\xa2\x95\xb2\xab\xba\x66\x4b\xbe\xde\x6b\x52\x7d\x33\x06\x53\xcf\xc3\xca\x6f\xab\x5d\x21\xbd\xc7\xa2\x77\xae\x37\x10\x11\x82\x72\x96\x92\x54\x28\xb7\x63\x7f\x98\x0f\x6e\x0b\xea\xea\xca\x1e\x0c\x20\xce\x42\xb4\x93\x3c\xcc\xaa\x15\xe7\x55\x93\x42\x32\xd2\xd0\xc6\xa0\x53\x64\x20\xf2\x28\xca\x52\x2a\xe5\xa2\xa5\x1b\x9c\xb5\xd2\x95\x2e\x27\xeb\x32\x3f\x53\x8d\xdc\x27\xd8\xdf\xf4\xce\x85\xdb\xf7\xd2\xb2\x51\x77\x10\xda\xaf\xe5\xfc\x5f\xab\x20\x50\xa2\xcc\x0f\x6a\x9a\xa6\xab\x22\xbd\xc6\x8b\x74\x71\xc6\xaa\xcc\x79\x4c\xd4\xf5\x6c\x08\x2a\x00\xa5\x72\xfb\xf8\xfc\x54\xdd\xad\xef\xff\x1f\x5e\x1e\x0e\xe3\x86\x5e\x17\xc3\x9a\xd3\x3e\xff\x65\x80\x4e\x87\x7c\x8b\xb4\x18\x99\x78\x63\x31\x58\xa1\x81\x3e\xc4\xd4\x8b\xb3\x2c\xea\x62\x2e\x9c\xdb\xa4\x70\x1e\x91\x36\x18\xa7\xa4\x19\x1c\x12\x7b\x3f\x9d\xbe\x03\x00\x00\xff\xff\x7a\x51\x3e\x00\xb2\x01\x00\x00")

func client_secretJsonBytes() ([]byte, error) {
	return bindataRead(
		_client_secretJson,
		"client_secret.json",
	)
}

func client_secretJson() (*asset, error) {
	bytes, err := client_secretJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client_secret.json", size: 434, mode: os.FileMode(420), modTime: time.Unix(1556967287, 0)}
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
	"client_secret.json": client_secretJson,
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
	"client_secret.json": &bintree{client_secretJson, map[string]*bintree{}},
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


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
