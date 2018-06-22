// Code generated by go-bindata.
// sources:
// schema/1.0.0/seed.manifest.example.json
// schema/1.0.0/seed.manifest.schema.json
// schema/1.0.0/seed.metadata.schema.json
// DO NOT EDIT!

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

var _schema100SeedManifestExampleJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\xdf\x6f\x22\x37\x10\x7e\x2e\x7f\xc5\xc8\xba\xa7\x8a\x9f\xb9\x03\x55\x3c\xb5\x3d\x82\x42\x55\x92\xd3\x85\xe4\xa5\x8a\x90\xe3\x9d\x5d\x4c\xd6\xf6\xd6\xf6\x86\x50\xb4\xff\x7b\x65\x2f\xb0\xde\xb0\x34\xb4\x52\xf3\x14\xcf\x7c\x33\xf3\xf9\x9b\x19\x2f\xbb\x16\x00\x31\x88\xd1\x23\x6a\xc3\x95\x24\x63\x20\x83\x6e\xbf\xdb\x27\x6d\xe7\x59\xab\x67\x32\x06\x07\x02\x20\x92\x0a\x74\x7e\xb1\xed\x38\x7b\xbb\xb4\xae\xd5\x73\x63\x2c\x00\xc9\x28\x7b\xa1\x09\x9e\x73\x5b\x6e\x53\x9f\x70\xbe\x85\x98\x6b\x63\x21\x48\x1b\xa1\x61\x9a\x67\x76\x1f\xf8\x1d\x69\x64\x80\x4a\xb8\x99\x4c\x87\x10\xf3\x14\x81\xca\x08\x54\x6e\xb3\xdc\x1a\xb0\x1b\x05\x8b\xd9\x74\x0a\x5c\xd0\x04\x4d\x1b\x28\x7c\xbd\x7f\xf4\x10\x41\x25\x8f\xd1\x58\x60\x4a\x5a\xca\x25\x97\x09\x30\x4c\xd3\x25\x53\xb9\xb4\x47\x2e\x34\x31\x64\x0c\x7f\xf8\x13\x00\x59\x45\xf1\x70\xef\xf3\x4c\xe3\xb8\x3a\x31\xf3\x5a\x1d\x7c\x41\xc8\xb4\x62\x68\x0c\x97\x09\xf1\x8e\xa7\x7d\x5a\x41\xb9\x2f\x8a\xfa\x28\x63\x20\xe4\x6f\x6a\x25\x61\xa2\xb0\xca\xa6\x74\x42\x25\xff\x8b\x1e\xae\x7d\xdd\x61\x4a\x67\x95\x1f\x05\xe5\xa9\x73\xac\x23\x85\x3f\xe3\x1b\x15\x59\x8a\x5d\xa6\x44\x05\xc9\xb5\x07\xac\xac\xcd\xc6\xbd\xde\x66\xb3\xe9\x36\xc2\xb2\x95\x92\x9e\xc4\x68\x34\xea\x0c\x87\xc3\xce\x97\xcf\x57\x83\x92\x7c\x71\xec\x8f\x40\x95\x5b\x32\x86\xcf\xa3\x7e\x7f\x6f\xe4\xd2\xa2\x8e\x29\xc3\xf0\x42\x4c\x09\x41\x65\xe4\xd2\x7d\xda\xcd\x6e\xbf\x3d\x2c\x96\xd3\xd9\xef\xd7\x05\x7c\xda\xdd\x3d\x2c\xdc\x71\x32\xfb\x5e\x04\xa2\x49\xd7\xb5\x20\x03\x00\x71\x3d\x0d\x5b\xe0\xfe\x76\xc1\xff\x81\x6e\x55\x05\xd2\x6e\xfd\xb0\x1f\x35\x6d\x39\x75\x37\xb7\x3a\xc7\x76\x3d\x4e\xe3\x9f\x39\xd7\x18\x35\x7b\x05\x46\x9c\x2e\xb6\xd9\x49\xf5\xaa\xbf\xbd\xb7\x8e\x1b\x88\x8e\x3f\x90\x1a\xe4\x29\x38\x15\xad\xf7\xd6\xa2\xea\x6c\x39\xa9\xff\xf9\xce\x65\xfc\xd2\x45\x2c\xdd\x38\x1a\x72\xee\x1a\x0e\x5d\xb2\xae\x8d\xed\x01\x96\xa7\x96\x67\x7e\xef\x1a\xb4\xc8\xa8\xb5\xa8\xe5\xbe\xa0\x2b\xf6\x63\xd7\xf2\x38\xbc\x72\xd1\xfe\x77\x64\xc3\x6d\x69\xa2\x6a\xf1\xcd\xf6\x1a\x40\x4d\x54\x1c\xac\x59\xef\x2a\x98\xac\x8d\xdf\x9d\x0b\x34\x3d\x79\x06\x8e\x88\x17\xdc\x1e\x00\x5f\x9b\xfc\xf6\xa0\xb3\xb4\x98\xa0\x3e\x43\xea\x64\x08\x84\xcb\x55\x6f\x78\x48\xed\x48\x6c\x7e\xf7\x70\xbb\x58\x7e\xfb\x65\x71\x53\x2b\xec\x34\x59\x39\x7f\xcf\xae\xb0\xb7\x7f\xcf\x50\xf7\xbc\xb9\x06\x14\x2a\xf2\x89\xb4\xaa\xa8\x1d\x88\x1d\xb5\x22\x06\xad\xe5\x32\xb9\x80\xd0\xe4\xd7\xe5\xcd\xdd\xfd\xa2\x5e\xc4\x20\xd3\xe8\x9e\x87\x98\xa6\x06\x4f\xeb\xd4\x1e\x13\x8d\x46\xe5\x9a\x61\xb8\x01\xc4\x30\x9a\x52\xfd\x71\x79\x96\xe5\xf5\x79\x27\xaf\x34\xcd\x9d\x6b\x50\x95\x6d\xff\x73\x0e\x81\xe2\x4c\x8a\xab\x9f\x2e\x4e\x62\x56\x54\x63\x34\x3f\x97\xaa\x7f\x71\xa2\x88\x9b\x97\x33\x74\xfa\x35\xb3\x7f\x2a\xe7\xe5\xd2\x72\xff\x19\xf9\xf2\x81\xd2\xa8\xb5\xd2\x61\x4f\x83\x07\x87\x95\x73\x31\x08\xd6\xe5\xf8\x15\xbe\x76\x71\x70\xeb\x08\x06\xee\x77\xdf\xe1\x12\x34\x09\x8c\x01\x96\x51\x8b\x89\xd2\x7e\x75\x22\x6a\x29\x79\xbf\x00\xa7\x4c\xae\xfe\x7f\x26\xee\x47\x45\x2b\x14\xcc\xc9\x55\xb4\x8a\xd6\xdf\x01\x00\x00\xff\xff\xad\x4d\xb2\xc2\xfc\x08\x00\x00")

func schema100SeedManifestExampleJsonBytes() ([]byte, error) {
	return bindataRead(
		_schema100SeedManifestExampleJson,
		"schema/1.0.0/seed.manifest.example.json",
	)
}

func schema100SeedManifestExampleJson() (*asset, error) {
	bytes, err := schema100SeedManifestExampleJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/1.0.0/seed.manifest.example.json", size: 2300, mode: os.FileMode(448), modTime: time.Unix(1523280816, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schema100SeedManifestSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\xcb\x6e\xeb\x36\x10\xdd\xeb\x2b\x08\xf5\x2e\xec\xa4\x8a\x1d\xa0\x9b\x78\x53\xe4\x03\x0a\x74\x51\x74\x11\xdb\x2d\xc6\xd2\xd8\xa6\x23\x91\x2a\x45\xb5\x70\x1a\xff\x7b\xa1\x87\x13\x99\xe2\x43\xf2\x23\x4d\x1b\x5f\xe0\xe2\x1a\xa3\xe1\xcc\x21\x79\xe6\x70\xae\xc4\xbf\x3d\x42\xfc\x6f\x59\xb8\xc6\x04\xfc\x09\xf1\xd7\x52\xa6\x93\xd1\x68\x93\x71\x16\x54\xd6\x3b\x2e\x56\xa3\x48\xc0\x52\x06\xe3\x1f\x46\x95\xed\x3b\xff\xfb\x62\x9c\xdc\xa6\x58\x0c\xe2\x8b\x0d\x86\xb2\xb2\x41\x14\x51\x49\x39\x83\xf8\x67\xc1\x53\x14\x92\x62\xe6\x4f\xc8\x12\xe2\x0c\x4b\x87\xb4\x69\x2e\xd2\x13\xe2\x67\x88\xd1\xaf\x28\x32\xca\xd9\x9b\xb1\x11\x3f\x93\x82\xb2\x55\x19\xbf\xb4\xa7\x20\x25\x8a\xc2\xd5\xff\xed\x7e\x36\xbb\x1b\x17\x7f\xbf\xf9\xe5\xe3\x5d\xe5\xe5\x6f\xf8\x42\x17\xaa\x01\xb5\xb4\xbb\xe0\x56\xf9\xda\x90\x4b\x3b\x83\x04\x0f\x2c\x66\xcc\x2d\xdc\x53\x08\x5e\xc6\xc1\xc3\xef\xc1\xfc\xb6\x46\xde\x40\xbf\x9f\x41\x7b\x4d\x7a\xe5\x18\x8c\x5f\xa7\xf7\xc1\xc3\x7c\x3a\x0e\x1e\xe6\x37\xc3\xd9\xec\xce\x69\x19\x04\x87\x86\xd7\xea\x9f\x02\xed\x63\xf0\x14\x94\xd6\xfd\xef\x9b\xe1\xa0\x15\xc0\xee\x3f\xbc\x19\xfe\x38\x98\xcd\x6e\x9b\xd6\xdb\x22\xc8\x81\xa1\xf0\x32\xac\x49\x0a\xe1\x33\xac\xf0\xba\x2e\xca\xba\x48\x2a\x63\x17\x15\xb5\x03\x23\xcc\x42\x41\x53\xe9\x5e\x4d\x7d\x5e\x58\x65\xa6\x71\x20\x04\x6c\x0f\x37\x81\x4a\x4c\x54\x7f\x4b\x26\x42\x76\xda\xac\x09\x50\x26\x81\x32\x14\xa6\xdc\x4a\x99\x93\xae\xa5\x4e\x2c\xe5\x4e\x0c\x25\xef\x98\xc2\x01\xf4\xd2\x97\x8b\x15\x30\xfa\x02\x9a\x45\xef\x1d\x0b\x13\xa0\xf1\xa9\x41\x72\x71\x72\x88\x74\xcd\x59\xff\x85\xf1\x0c\x01\x7d\x81\x7f\xe4\x54\x60\xe4\x4f\xc8\x54\xb3\x01\xda\x45\x68\xd8\xe6\x86\x22\x49\x90\xe7\xd2\xc4\x19\xca\x24\xae\x50\xe8\x89\x2e\x30\xe3\xb9\x08\x5b\x94\xb8\x3c\xe3\xb2\x10\x62\x50\x89\x4e\xec\x85\x46\xcc\xc5\xe6\x02\xdd\x13\xba\x7b\x02\xb5\x87\xa1\x70\x14\x40\x1a\xe5\x6e\x66\x51\xce\xcf\xc7\xe0\x49\x39\x3f\xdf\xff\xec\x74\x31\xfc\x3f\x21\xce\x3b\xc0\x60\x79\xb2\x68\x32\xc1\x15\x96\xb2\x34\x97\x3f\xe5\xb1\xa4\x69\x4c\x5b\xaa\xd4\x2f\x41\xcb\xa6\x49\x69\xae\x8f\xfa\xb9\xa6\x4a\xea\x27\xd5\x0a\xb4\x9e\xcc\x3d\x47\x52\x7b\xca\x3d\x4b\x3d\x5b\xd4\x9d\x53\xd8\x8b\x12\x14\x4b\x08\x8d\x47\xd9\xc5\xaa\x2c\xe4\x49\x02\x2c\x3a\x55\x06\x4b\x22\xe8\x6a\xc0\x55\x72\xbd\x0a\xce\x55\x6e\xfe\x92\xc6\xe6\x4a\xb4\xcb\x06\xb1\x4b\x47\x97\xb9\x1c\x31\xa3\x6e\xf3\xaa\xbd\xac\x52\x42\x3a\xcb\x09\xe9\x2b\x29\xc4\x54\xff\x44\xa9\x8f\x0e\xc8\x16\x9c\xc7\x08\xcc\x06\x2d\xc2\x25\xe4\x71\x71\x5c\x49\x91\x63\x6f\x3c\x09\x46\x14\x7e\xd9\xa6\xd6\xb5\xec\x44\x87\xda\xd1\x4e\x0a\x25\x9a\xae\x58\x0e\x70\xf7\x9f\x4f\x25\xaf\x1d\x77\xbe\xd7\xfa\x96\x94\xec\x0d\x28\x05\x21\x29\xe8\x3a\xa7\x8b\xe2\xd1\xda\x0d\x28\x5d\xe7\x04\x79\x2b\x27\xed\x43\xf5\x50\xd0\xe7\xd7\x9d\x4f\x9b\x4c\xdb\xda\x92\xab\xfe\xa8\x19\xff\xb7\xfa\x53\x67\x3a\xcb\x2a\x21\xcb\x13\x0b\x85\x4b\x1f\x97\x7e\x91\x2e\x93\x26\xef\xff\x0d\xb0\x3a\xd5\xfd\x9b\xd5\xc7\xca\xce\xda\xc7\x25\x93\xba\x0a\x24\x97\x53\x01\xfb\x66\x9e\xa2\x11\x6a\x7b\x69\xed\xa2\x78\x2e\xaf\x6d\xd4\x31\x33\xea\x36\xaf\xda\xeb\x73\xca\xd8\x5b\xdb\xd2\x0b\xda\x11\x87\xf7\x1e\xf0\x05\x93\x7c\xba\x96\xe5\x5f\x3a\x23\x3e\x58\xac\xf6\x7b\x7b\xed\x69\xfe\xeb\x62\xf0\x8c\xdb\xcb\x56\xe8\xb5\x49\xb9\x64\x93\xf2\xc5\x75\xe8\x63\x9b\xa6\x84\xe7\xcc\xde\x33\x7d\x95\xf7\xc4\xe4\x3c\xef\x8a\x53\x90\xeb\xce\x48\x3a\x47\x4d\x78\x64\x99\x9f\x43\x43\x7c\xc1\xcd\x74\x13\x7f\x19\xc8\x66\x20\xf8\x7b\xcd\x14\x61\x8f\x3e\x16\x8f\x7f\x15\x5d\x2e\x70\xeb\x41\xeb\x4d\xb4\x95\xf5\x19\x4a\x49\x59\xeb\x23\x24\xf9\x8a\xbc\x3f\x03\xe7\x33\x0c\x05\xaa\x5f\xc8\x34\x38\xec\x92\xdc\xa1\x4f\x3d\x23\xb7\xfa\x32\xc8\xd3\xfd\x6e\x7e\x85\x40\x21\xb8\x38\xe3\x67\x6d\x2d\x8b\x7a\x30\xc8\xce\x1e\x3f\x34\x29\x8a\xf9\xcb\x66\x6b\xce\xfb\x01\x9a\x6b\x04\x4a\x2c\xbd\xde\xb5\x43\x99\x2f\x16\x1c\x19\x30\x04\x89\x2b\x2e\xf4\x0d\xa1\xbb\x40\x9a\x72\xb7\xe1\x0b\x9d\x87\x45\x7c\x0d\x43\x8a\xb0\x20\xe1\x54\x11\xb3\xd0\xbc\xda\x5d\xc3\x46\x79\xa6\x94\x0d\x5e\x7b\x4a\x4a\x7d\x32\x55\xa5\x9b\x97\x8f\xcc\xd7\x6f\x5a\x17\x50\x0c\x17\x4b\xf4\x57\x37\x34\x5f\xe6\xbd\xe6\x54\x0a\xe8\x25\xec\x36\xe4\x83\x0b\x63\x55\x9c\x72\x87\xbc\x62\xec\xce\xfb\x27\x00\x00\xff\xff\x0b\x94\xb7\x98\xd5\x26\x00\x00")

func schema100SeedManifestSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_schema100SeedManifestSchemaJson,
		"schema/1.0.0/seed.manifest.schema.json",
	)
}

func schema100SeedManifestSchemaJson() (*asset, error) {
	bytes, err := schema100SeedManifestSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/1.0.0/seed.manifest.schema.json", size: 9941, mode: os.FileMode(448), modTime: time.Unix(1510255815, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schema100SeedMetadataSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x4f\x6f\xdb\x3e\x0c\xbd\xe7\x53\x10\x6a\x8f\x6e\xfd\xc3\x6f\xc5\x86\xe5\x56\x0c\xe8\xb0\x61\x5b\x8b\xf5\x58\xe4\xa0\x3a\x74\xa2\x42\x96\x52\x49\x41\x17\x14\xfe\xee\x83\x15\xcb\x96\x6c\x25\x75\xda\xac\xf0\xcd\xa6\x48\xf1\xf1\x3d\xfd\xa1\x9e\x27\x00\xe4\x54\x67\x4b\x2c\x28\x99\x02\x59\x1a\xb3\x9a\xa6\xe9\x83\x96\xe2\x6c\x6b\x3d\x97\x6a\x91\xce\x15\xcd\xcd\xd9\x7f\x17\xe9\xd6\x76\x42\x92\x2a\x8e\xcd\x5d\x88\x9e\xa6\xa9\x91\x92\xeb\x73\x86\x26\xb7\x21\x4b\x53\xf0\x54\xe5\xd9\xa7\xcf\x17\x1f\x6b\x7f\xc3\x0c\xc7\x2a\xe4\x2b\x4a\xf8\x7e\x7b\xfd\x0b\xe4\xfd\x03\x66\x66\x3b\x3a\x47\x9d\x29\xb6\x32\x4c\x8a\xca\xe7\xd6\x66\x82\x5c\x2a\xa0\x10\x0d\x30\x9b\x95\x9d\xcd\xb7\x29\x7c\x5c\x33\x85\x15\xb0\xbb\xda\x03\x66\x76\x44\x0a\xbc\xce\x2b\xf3\x04\x00\xe0\x19\xc8\xa9\xc2\xea\x9f\x9c\xa4\x73\xcc\x99\x60\x55\x62\x9d\x2e\x50\x16\x68\xd4\x86\x40\x99\x0c\xf3\xfc\x22\x39\xc7\xcc\xc2\x7e\x31\x26\x47\x6a\xd6\x0a\x07\x3b\x06\x73\x4f\xa0\x2e\xc5\x73\x24\x53\x78\xb6\x33\x91\x06\xb8\xb3\xc4\x19\xb2\x76\x9f\xa5\xda\xe6\xbc\x93\xf6\x3f\x93\x52\xcd\x99\xa0\x06\x35\xa9\xad\xb3\x66\x8a\x90\x4e\x5b\x4b\xf3\xe5\x4b\x7d\x23\x99\x30\xde\xac\x00\x84\xce\xe7\x16\x3c\xe5\x37\x4a\xae\x50\x19\x86\x55\x1d\x39\xe5\x1a\x03\xc7\x95\x3f\xec\x4f\xdf\x96\x16\x5a\x01\x08\x8a\x75\x11\xc0\x6a\x46\xb6\x48\x3a\xf6\x59\xf0\x5f\x26\x61\x12\x9f\x81\x48\xae\xb8\x74\x2b\xa9\xed\x57\x98\xaa\x9c\xc4\xbe\xbd\x84\x71\xfa\x7e\xae\xb9\x61\xe3\xe1\xd0\x83\xf3\x7e\x44\x5e\x2a\x45\x37\x47\x61\xf3\x07\x13\x78\x6b\x14\x13\x8b\x51\xb0\xe9\xc1\xf9\xf7\x6c\xf2\x1d\xc9\xde\xb2\x30\x47\xc6\x67\x17\xd3\x31\x49\x75\x47\x29\xb5\x8b\x31\xe9\x0e\x33\x83\x45\x2c\xee\x60\x39\x42\x41\x5e\x2d\xcf\x8d\xe4\x9b\x85\x14\xa3\x90\xc5\x61\x79\x8f\x13\x23\x92\xe9\x6d\x27\xef\x78\x68\x0c\x00\x8d\x79\x69\x47\x55\x18\xb8\xae\x27\x7e\x3d\x75\x1d\x24\xd2\x69\xf9\x4d\x8e\xd7\x55\x76\xdd\x9a\x76\xa5\xd3\x5c\x5e\x42\xd6\x38\x81\xcc\xc1\x65\xa8\xfb\x4b\x1d\x6f\x95\x1a\x24\x95\xa2\x5e\x2f\xb4\x43\xe7\x46\xe3\x56\xd5\x28\x48\x98\xf9\x8a\xf9\x39\xa6\x9d\xb5\xb9\x53\xa8\x56\xa4\x01\x7d\x6d\x8f\xed\x32\x60\xdb\xf5\xa8\x31\x8a\xaf\xea\xb1\xdd\xbc\x36\x9d\x7a\x3d\x0d\xec\xeb\x3d\xbd\xae\x35\x09\x68\x7c\x25\xb9\x0e\x5e\x9c\xd1\x4d\x97\xcf\x7e\x0f\x0b\xdb\x9e\xdc\x11\x2d\xd6\x9c\x93\xee\x76\x3a\x88\x62\x7f\x6b\x96\x7b\x0e\x87\x26\xe7\x5d\xdb\xad\xbb\xfc\xb3\x4e\xa4\x7d\x76\x05\x11\xba\xbe\x83\xab\x88\xe2\x1e\x95\x8d\xd9\x27\xee\x0b\x3b\xe9\xaa\xe7\x75\x80\xe0\x59\x24\x2a\x14\xbd\xf6\x7c\xab\xcc\xbb\xb7\x50\x93\xe0\x88\x1b\xa8\x7d\xba\xed\xdf\x3f\x55\x2d\xee\x05\x10\x54\xd2\x63\x4e\x33\xb1\xe0\x08\x8d\x77\xd2\xab\xba\x07\x94\x14\x4c\x7c\xab\x91\xfe\xef\x9b\xe9\x1f\x67\xfe\xe0\x99\x5d\x51\xfe\x02\x8f\xdf\x43\x6e\xe1\xf8\xa7\x72\xf2\x8e\x51\xcd\xf7\xcc\x83\xdf\x5e\xb4\xae\x38\x7b\xc7\x4e\x3a\x99\x48\xf8\x50\xd8\x4b\xba\x00\x4b\x69\x75\xe6\xbb\x28\x3d\x88\xf8\xd8\x2d\x38\xf8\x01\x58\xf6\x20\x7b\xed\xdf\x50\xbc\xe6\x49\x82\x54\x50\x48\x85\x71\xec\x94\xf3\xde\x71\xd6\x51\xe0\xc0\x87\xd6\x3e\x31\xfd\x95\x18\x57\x32\x5a\x34\x55\xbf\x0f\x29\x3a\x97\x6b\xd5\x56\x0b\x4f\x4b\x54\x08\x66\x89\x90\x33\xa5\x0d\xe0\xe3\x9a\x72\x6d\x0d\x9c\x6a\x33\x0a\x36\x2e\x06\xb2\xe1\xda\xa4\xa1\x54\x6c\xd9\x83\x6a\xcd\x1c\x7d\xc9\x7a\xca\x44\x4f\xb7\x72\x52\xfe\x0d\x00\x00\xff\xff\xff\xd6\xdf\x2f\x2a\x14\x00\x00")

func schema100SeedMetadataSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_schema100SeedMetadataSchemaJson,
		"schema/1.0.0/seed.metadata.schema.json",
	)
}

func schema100SeedMetadataSchemaJson() (*asset, error) {
	bytes, err := schema100SeedMetadataSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/1.0.0/seed.metadata.schema.json", size: 5162, mode: os.FileMode(448), modTime: time.Unix(1525726087, 0)}
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
	"schema/1.0.0/seed.manifest.example.json": schema100SeedManifestExampleJson,
	"schema/1.0.0/seed.manifest.schema.json": schema100SeedManifestSchemaJson,
	"schema/1.0.0/seed.metadata.schema.json": schema100SeedMetadataSchemaJson,
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
	"schema": &bintree{nil, map[string]*bintree{
		"1.0.0": &bintree{nil, map[string]*bintree{
			"seed.manifest.example.json": &bintree{schema100SeedManifestExampleJson, map[string]*bintree{}},
			"seed.manifest.schema.json": &bintree{schema100SeedManifestSchemaJson, map[string]*bintree{}},
			"seed.metadata.schema.json": &bintree{schema100SeedMetadataSchemaJson, map[string]*bintree{}},
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

