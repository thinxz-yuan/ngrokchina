// Code generated by go-bindata.
// sources:
// assets/server/tls/snakeoil.crt
// assets/server/tls/snakeoil.key
// DO NOT EDIT!

// +build release

package assets

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

var _assetsServerTlsSnakeoilCrt = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x93\xb9\xd2\xab\x46\x10\x85\x73\x9e\xc2\x39\xe5\x12\x8b\x16\x08\x1c\x0c\x0c\x3b\x03\x0c\x20\x09\xc8\xd8\x34\xec\x42\x80\x40\xe2\xe9\xfd\xdf\x1b\x38\xb0\xdd\xe1\xd7\x5d\xa7\xba\xea\xd4\xf7\xe7\xaf\x91\x14\xcd\x70\xfe\x90\x15\x3f\x34\x54\x43\x06\xa1\xf2\x9b\x52\xc8\x30\xe4\x19\xca\x32\x48\x88\x2c\x63\xb8\xa7\xc7\x51\x70\xb3\x99\x8b\x43\xe0\x48\xa4\x7d\x55\x6d\xad\x89\x1b\x23\x01\x3c\xab\x00\x82\x14\xf9\x64\x53\x49\x0c\x6f\x18\x43\x08\xc4\x37\x95\xf0\xe6\x9a\x7e\x4f\x4c\xaa\x75\xef\x52\x7f\xbe\x63\xee\xb8\xe9\x55\xee\xa0\xb0\xdd\xdc\xd0\xe0\x50\x88\xbe\x28\x04\xdb\xfd\x17\xdb\xd1\xe6\x34\x80\xf9\x87\x35\x20\xa5\xfe\x9d\xf8\x7f\x81\x84\x28\x35\x02\x8c\x26\x07\x2f\x2d\x30\x32\x1e\x62\xe5\xe7\xa3\x2b\x00\x47\x43\xa2\xe0\x06\x7e\x1d\x58\xe0\x69\xfc\x40\x98\xee\x37\x91\x6c\x7e\xab\x22\x99\x0e\xc1\x9e\xb2\xd1\x20\x0b\x9d\x56\xcc\xe3\x99\x8d\x90\xd1\xd9\x92\x8c\x00\x69\x8b\x53\x36\x9e\xb1\xad\x74\x2d\x75\x00\xa4\x77\x03\x50\xf8\x3b\x2f\xae\xe6\xe3\xc9\x1f\xa2\x93\x24\xaa\x5c\xd0\x1e\x6c\xa1\xd5\x19\x9a\x11\xdf\x7a\x43\x5c\x5a\xb1\xf1\xab\xea\x1d\x1b\xec\x42\x2e\xec\x18\x7d\x9e\xc1\x57\x1d\x58\xea\xf2\x19\xeb\x74\x00\xef\x76\x9e\x71\xbc\x46\x3c\x4b\x78\xae\x4d\xdc\xea\x4c\x98\xf4\xf5\x2d\x5f\x65\xb8\x8e\x97\x13\x7d\x1c\xc9\x39\xd8\xe9\x74\x6c\x97\xf2\xa1\x66\x31\x62\x95\xce\xf3\x3f\x75\x48\xe9\x6f\x16\xfb\xfa\x43\xf5\x63\x33\x60\xa7\x18\xe6\x65\x6d\xd7\xd7\x77\xd6\xa0\xf2\xa4\x4f\xd3\x6b\xea\xeb\xe4\x3a\x7f\x0e\x2c\x8f\x05\x71\x16\x8d\x87\x2b\xc6\x72\xae\x34\x43\xff\x74\x62\x32\x29\x54\x5b\x67\x2e\xad\x6f\x9d\xf9\xb0\x64\x11\xa9\x9a\x2a\xae\xd3\x75\xea\x64\xda\x6b\x62\x09\x99\xcb\x05\xf9\x37\x2f\x98\x93\x3e\x54\x7d\x64\x44\x07\x78\x47\x9c\x50\x04\xf8\x38\x8c\xaf\xe5\xfb\x14\xa9\x17\x5b\xde\x27\x96\x3e\x48\xf4\xda\xe5\x8d\xa0\xd2\x33\x1f\xc1\x6b\xc7\xf8\xaf\x47\x9e\x59\x86\x30\x5d\x3c\x1d\x10\x24\x01\xa0\x6c\x10\xc7\xa6\xf5\x4c\x8c\x6a\xcd\x1d\x80\x15\x5b\xc2\x00\x52\x3f\x0d\xfc\x2c\x67\xd6\x1e\xd7\xb2\x5f\x1c\x71\xc4\xdb\xb1\xf6\x98\xa2\xcf\xe8\x5e\xe6\xbd\x63\x97\x0c\x9a\x2c\x3c\xb4\x6a\x43\xae\x12\xb3\x59\xa4\xb5\x74\xa1\x57\x97\x7a\x89\xfa\x2b\xbe\x1b\x54\x6f\x2d\x3e\xc7\x61\xe7\xae\x48\xfb\x69\x33\xaa\x8f\xb6\xa2\x3e\x83\xa1\xe9\xba\xae\x93\xdc\x0d\x5a\xeb\xb7\xee\xc6\x3e\xe3\x53\x4e\x68\x9f\xf6\x42\x26\xda\xad\x3c\xc8\xe6\x81\xed\xb8\xf2\x09\x28\xbe\xbc\x78\x9c\xd9\x7e\x1c\xe0\x29\x95\x9c\x3e\xc4\x73\x7f\xaa\xb5\xda\xa2\xf3\x8b\x5a\x84\x9e\x3f\x99\x23\xf4\xba\xdb\x58\x16\x66\x7f\xb8\xee\xcb\xe4\x1f\xda\x02\x3a\xc4\x78\x0a\x4e\xd1\x54\x1a\x05\x3f\xb5\xc5\x08\xa3\xda\xd7\xf7\x1b\x8b\x5d\xb0\x9e\xe3\x7b\xb2\xbb\x04\xa3\x4b\xb3\xe7\x1f\xfb\x68\xc3\x72\xb9\xb2\x5d\x9c\xa8\xd6\x50\x58\xad\x35\x64\x8d\xf3\x81\xa9\xf9\xed\x7c\xb1\xf1\x3d\x4a\x6e\xaa\x81\xaf\x8a\xbb\x42\xc0\x55\x64\x99\x61\x96\x92\x79\x04\xda\xde\xe1\x45\x8f\x64\x64\x8f\xb2\xae\x61\xc9\x21\xe2\xf9\x02\x16\x8c\x5b\xda\xc8\x8e\xd2\xcc\x07\x86\x7e\xd6\x3e\xdc\x1e\x50\x7d\xa7\xa8\xeb\xba\x32\x4b\xef\xae\x27\x93\x56\x53\xbc\xd6\xf5\x51\x62\xa3\x35\xff\x8b\xfa\x2d\xac\xe2\xc0\xff\x4a\xfc\x77\x00\x00\x00\xff\xff\xc9\x30\xf4\x28\xe1\x03\x00\x00"

func assetsServerTlsSnakeoilCrtBytes() ([]byte, error) {
	return bindataRead(
		_assetsServerTlsSnakeoilCrt,
		"assets/server/tls/snakeoil.crt",
	)
}

func assetsServerTlsSnakeoilCrt() (*asset, error) {
	bytes, err := assetsServerTlsSnakeoilCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/server/tls/snakeoil.crt", size: 993, mode: os.FileMode(420), modTime: time.Unix(1569504060, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsServerTlsSnakeoilKey = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x95\xb7\x0e\xa4\x50\x12\x45\x73\xbe\x62\x72\xb4\x02\x1a\xd3\x10\xe2\x1a\xfb\x68\xbc\xcb\xf0\xbe\xf1\xf6\xeb\x77\x76\xe2\xad\xac\x54\xc9\x51\xe9\x5e\x9d\xff\xfc\x6f\x38\x51\x52\x8c\x3f\xb6\xc3\xfe\x31\x6d\xc5\x67\x5d\xf1\x8f\x26\x46\xff\x2e\x10\x50\x14\x71\x3c\x15\x8e\x65\x35\x9e\xb5\x44\xf6\xb5\x62\x65\x04\xc4\x98\x73\xd9\xa3\x3b\xc1\x6b\xf3\xc9\xf3\x50\xed\x5f\xaa\x15\x9b\xff\xf0\xaa\x73\x5a\x2d\xab\xa8\x72\x11\x50\x45\xc7\x5f\x2a\x14\x9b\xa7\xa2\xb6\x5d\x25\x7b\x19\x83\xa4\x77\x48\x19\x08\x06\x5b\xa5\x9d\x77\x93\x79\x23\x2a\xc7\x98\x9b\x99\xd6\x04\x21\x1c\x35\x8f\x6a\x63\xdc\x3a\x27\x40\x64\x13\xa0\x02\x48\xc4\xb5\x8e\x21\xa6\xa0\x93\x68\x98\x4e\x7d\xd2\x75\x51\xd2\x31\x86\x89\x0c\x66\x92\xdc\xb1\x18\x0d\x69\x5e\x7f\xf3\xaf\x7b\x53\x05\x5c\xee\x5a\xf4\xed\x56\x64\x98\x63\x3d\xfc\x5d\xc1\x4b\x30\x6c\xd5\x45\xb3\x08\xea\xae\xf7\xe6\x89\x36\x7e\x79\x01\xef\x6d\xc9\x8b\xc5\xe5\xb1\x21\x7a\x7d\x20\x1e\x79\xb7\x29\x8a\x9a\xc9\x71\xf0\x75\x50\x32\x39\x6a\x9a\xa9\xe9\xc8\xcf\x11\xb0\x3f\x4e\x21\xc9\x59\x08\x14\x0d\xba\x54\x65\x78\x8e\x9a\x56\x9d\xf0\xae\x0e\x97\xb3\xeb\x32\xa5\x30\x8d\xb4\x8e\x96\x78\xb1\x2e\x9f\x16\x8f\xe8\xbb\xe8\x22\xc5\xdd\xe5\xb9\xfc\x07\x39\xb1\xd6\x38\x64\xaf\xfb\xaa\x54\xb2\x64\x33\x64\x26\x5b\xd8\xcf\x4c\x79\x9c\xe5\x42\x86\x0a\xc2\x95\x8b\x81\x9d\x98\x9a\x8b\xc9\x0f\x97\xee\xc6\xd4\xe0\xe7\x3a\x15\x81\xb5\x58\x8e\x1d\xff\x3e\xdb\xe2\x2d\xb1\x71\xa3\x90\xc2\x66\xb2\x71\x72\xc8\xa6\x8d\xcb\x44\x01\x15\x98\x83\x47\x99\x88\x1d\x0f\xeb\xd9\x9d\x5f\x46\xb5\x76\x43\x31\x34\x7f\x68\xbe\xfd\x8a\xa8\xfd\xa9\xbf\xb6\x39\xa9\x0b\x98\xbb\xb5\xd3\x0f\xc4\x24\x99\xc8\xa1\x5b\x12\xe8\x98\x5d\xae\x80\x1f\xda\x17\x0a\x37\x60\x40\xc2\xe8\xd1\x18\xbf\x23\xdb\xfb\xce\xfb\x0e\x72\xdf\x03\xdd\xd3\xb1\xca\x5e\x07\x7d\xb3\xfa\xdf\x57\x16\x01\x29\xdb\x74\x94\x74\xaf\x6b\x64\x6e\x19\x12\xe0\x70\xca\x5c\x47\xb7\x84\x58\x4b\x3b\xef\xf4\xa5\x14\xd4\x24\xda\x51\x6d\x98\x78\x91\x62\x45\x3e\xa1\x14\x45\xb4\x08\x38\x23\xea\x4f\x31\x37\x40\x6d\x6a\x3d\x07\xb9\xb4\x3a\xc0\xef\x18\x28\x4d\x2a\x3b\x76\xb7\xcd\x05\x83\x89\xde\x87\x81\x21\x4e\x95\xc2\x1d\xa3\x68\xae\xf5\x59\xdb\xe5\x52\x7b\x85\x9e\xb4\xcd\x1c\xda\xa7\xe4\x18\xd2\x2f\x0a\xab\xf2\xf2\x73\x6e\x66\xea\x25\x21\x12\xb4\x1e\xd5\xfd\xa3\x25\xb4\xdf\xb2\xd5\x3a\x43\xd1\x68\x09\xae\xac\x06\xc7\x22\x6a\xda\xb0\xe0\xac\xab\xf5\x70\x5f\x6a\x89\x74\x86\x7d\x4a\x32\xb1\x5c\x75\x9d\x00\xb9\x21\x5b\xeb\xf4\x15\xdb\x0c\x42\xca\xb6\xd4\xa9\xd7\xc5\x8e\x12\xc7\x9a\x99\x0f\xae\x79\x98\xcb\x0c\x7b\xff\x84\x6e\xf8\x48\xcb\x50\xdd\xae\x1c\x2f\xc4\xb4\x5f\x9e\x2a\x33\xd9\xe9\x3d\x64\xef\xee\x9f\x30\xe5\xfa\xa9\x6b\x31\x88\xc9\x5d\x5b\xe5\xa7\xf7\xfb\x9b\xe1\xd4\x92\x56\xb8\x5f\x8c\xb7\xc8\xd9\x81\x87\x39\xed\x48\x30\x3c\xce\x34\xa6\x90\x58\xee\x53\x27\xb2\x2c\x93\xa7\xc6\x0a\xd1\x3e\xec\x7d\x59\x8b\x91\x5f\xa0\x10\xdb\x62\x21\x23\x48\x30\xdd\xc8\x58\xf0\xc3\x69\xb7\xf6\xe8\xd2\x14\x61\xf8\x97\xc6\xd8\xbb\x66\xc1\xe4\xcf\x88\x42\xc5\xfb\x1d\x5c\xe6\xe7\xf6\xcf\xa8\x10\xf8\x51\x3b\xfe\x11\x7f\xb5\x1b\x22\x70\x39\xcb\xd4\x69\xe4\x44\x41\xac\xe9\xd5\x12\x50\x49\xa8\x6a\x4d\x53\xb3\xd0\x13\x68\x37\xa5\xf6\x07\x07\xc5\x3b\xe8\xa5\xfb\xb5\x3f\x8b\xec\xd9\x20\x0a\x1c\xa2\xfb\xbe\xaf\xef\x43\xb4\x31\x14\x04\x45\x03\xa2\xd9\x72\x7a\xdc\xf5\xc8\x96\x13\xee\x7c\x23\xd2\x40\xfb\x04\x52\x8b\x01\x22\x63\x59\x6e\x13\xbf\xda\xb7\x19\x66\xa1\x38\x31\x65\x60\x42\x4c\x8b\x98\xa3\x57\x5c\x0f\x3c\x8a\x07\x15\xed\x32\x46\xc7\x27\xb7\xd2\xf3\x90\xe3\x41\x6f\xad\x7e\x4b\xbc\x66\xdb\x85\xf6\x89\xcb\x96\x2c\xb4\x98\x5a\xd9\x78\xfa\x8b\xfc\x37\xc5\x30\x31\xc3\xc9\xdf\x1a\x15\x0e\x32\xbf\xdd\xeb\xd8\x0b\xc8\x10\x59\xb3\xa2\xe1\xbd\xf6\x2f\x3e\x3b\x8a\x6e\x20\x77\x69\xac\x1c\xd8\x3f\x64\x4e\xcd\x3b\xb9\xd0\x93\x03\xc1\xb1\x8c\xd7\x84\x5e\x20\x32\xeb\x65\xcb\xc1\x40\xa5\x52\xe8\xb4\x0d\x88\x69\x02\x92\x86\x69\x1a\x10\x3a\x54\xfd\xed\x1d\xb5\xd7\x69\xef\x22\xe5\x31\xd7\x71\xdd\x26\x9a\xa4\xde\xd9\xd1\x22\x5f\xe7\xf8\xe7\x3a\xa8\x00\x96\x5f\xe6\x85\x6a\x62\x3a\x1b\xc5\x8c\x6f\x35\xc9\xa8\x0c\xd4\x45\xfe\xb7\xd2\x5b\x1a\xa9\x9a\x0b\x9b\xf2\xf5\x62\x6b\x63\x4b\x68\xbe\x8a\x38\x27\x87\x01\x7e\x13\xb3\x2c\x3a\xbf\x86\xca\x0f\x36\xe9\x94\xbb\xaa\x82\xbb\xd8\x45\xf0\xc2\x51\x4c\x38\xe1\x06\x72\x97\xe4\x2e\xad\x35\xec\x3d\x9c\xd1\x8f\xf7\x2b\xfe\x54\x1d\x55\xe6\x98\x2a\xab\x13\xe7\x72\x24\x4c\xe3\x0b\x85\xed\xb7\xc7\x6f\x40\x8a\x72\xed\x78\xfb\x93\xdf\x64\x6f\x71\x63\x92\xf4\x05\x56\xe8\x52\xb8\x15\x1e\xe1\x0d\xce\xf9\x15\x7f\x8d\x71\x3c\x6e\x80\x7a\xf7\xf3\x02\x36\x09\x21\x09\xbd\x29\xeb\x2c\xa7\x3e\x87\x61\xf2\xc5\xc4\x50\xf6\x21\x39\x79\x6f\xdc\x74\x77\xa6\x51\xf7\xe5\x56\x48\x25\xf7\x34\xb1\x34\xae\xe2\xb6\x26\x01\xf1\x74\x3e\xb9\x8c\xd0\xcd\x00\x63\x01\x4a\x35\x1c\x13\x7b\x70\xf5\x7d\xdf\xfd\xcf\x6d\xbb\xa8\x6e\x07\x58\x72\x0d\x63\xef\x10\xd2\x43\xb1\x20\x7a\x21\xd0\x85\xcf\xdb\x51\xe0\xf5\x35\x7c\xf6\x14\xe8\x52\x5e\x35\x6e\x66\xd2\xd6\x40\xef\x4b\xd1\x9f\xdf\x4d\x68\xc7\x4f\x04\x1e\xfe\xf6\xe5\x36\xa6\xeb\x08\x2b\x88\xbc\xc2\xef\x1d\xf9\x84\xf0\xca\x29\x90\xbc\xf4\xe0\xef\x72\xa2\xc3\x51\xfc\xe6\x5b\x78\xa2\x0f\x57\xd0\xda\x3e\x8c\x73\x3f\x98\x52\xff\xfe\xb0\xfd\x36\x6b\xcf\xb5\x4e\x0f\xee\x52\xdd\x2a\x26\x37\xf4\x4f\x29\xa2\x21\xfc\x7f\xd5\xfc\x37\x00\x00\xff\xff\x0e\x2e\x97\x54\x8b\x06\x00\x00"

func assetsServerTlsSnakeoilKeyBytes() ([]byte, error) {
	return bindataRead(
		_assetsServerTlsSnakeoilKey,
		"assets/server/tls/snakeoil.key",
	)
}

func assetsServerTlsSnakeoilKey() (*asset, error) {
	bytes, err := assetsServerTlsSnakeoilKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/server/tls/snakeoil.key", size: 1675, mode: os.FileMode(420), modTime: time.Unix(1569504060, 0)}
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
	"assets/server/tls/snakeoil.crt": assetsServerTlsSnakeoilCrt,
	"assets/server/tls/snakeoil.key": assetsServerTlsSnakeoilKey,
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
	"assets": &bintree{nil, map[string]*bintree{
		"server": &bintree{nil, map[string]*bintree{
			"tls": &bintree{nil, map[string]*bintree{
				"snakeoil.crt": &bintree{assetsServerTlsSnakeoilCrt, map[string]*bintree{}},
				"snakeoil.key": &bintree{assetsServerTlsSnakeoilKey, map[string]*bintree{}},
			}},
		}},
	}},
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

