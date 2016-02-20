// Code generated by go-bindata.
// sources:
// templates/base.html
// templates/config-base.html
// templates/login.html
// templates/register.html
// templates/settings.html
// templates/shitbucket-import.html
// templates/tag-index.html
// templates/url-edit.html
// templates/url-index.html
// templates/url-new.html
// templates/url-view.html
// DO NOT EDIT!

package app

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

var _templatesBaseHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x58\xef\x8e\x1b\xb7\x11\xff\xee\xa7\x60\xd6\x41\xee\x43\xbd\xda\x5c\x6a\x14\xa8\x23\x29\x70\x93\x38\x35\x6a\xb4\x40\xec\xe4\x6b\x41\x2d\x67\x25\xfa\xb8\xcb\x35\xc9\x95\x4e\x30\xee\x39\xfa\x40\x7d\xb1\xce\x90\xdc\x15\xf7\x8f\x7c\x87\x16\xfd\x70\x27\xfe\x19\xce\x0c\xe7\xcf\x8f\x33\xfb\xf9\x33\x13\x50\xc9\x06\x58\xd6\x19\x95\xb7\xdc\x38\xc9\x55\xc6\x1e\x1e\x9e\xad\x85\x3c\xb2\x52\x71\x6b\x37\x99\xd1\xa7\x6c\xfb\x8c\xb1\x74\xad\xd4\x2a\x57\xfb\xfc\xf6\x3b\xbf\x33\xde\x23\x5e\xa5\x6e\x1c\x47\xce\x06\xf7\x3d\x01\x92\x1c\x5e\x6e\xe3\x10\x27\x9c\x1d\x0c\x54\x9b\x0c\x75\x58\xfd\xf6\xeb\x3b\xfa\x43\xb9\x19\x73\xdc\xec\xc1\x6d\xb2\x7f\xee\x14\x6f\xee\xb2\x6d\xbf\xff\x41\x3a\x05\x48\xb1\x2e\x78\xcf\x65\x5d\x10\xc7\x7e\xd2\xf6\xe2\x1d\xdc\xbb\xbc\xee\x1c\x08\xc6\xd5\x89\x9f\x6d\x7e\x32\xbc\xcd\xb6\x6b\x5b\x73\xa5\xb6\xa9\x64\x03\x47\x30\x36\x5e\xff\x28\xe1\x94\xb1\x4c\x8a\x2c\x48\x7c\xfb\x13\x29\x34\x28\xf0\x46\x9b\x9a\x3b\xe4\xfa\xa3\x01\x8e\x3f\xaf\x5d\xd4\x86\x7d\x53\x4b\x21\xb4\xfb\x9e\x8d\xef\x42\x53\x59\x31\x69\xcf\xba\x73\xdd\x0e\xd2\xbd\xe1\xc8\x5a\xf6\x6a\x57\x9c\x55\x3c\xaf\xa4\xaa\x51\xd5\x42\x92\x5c\x68\x84\x97\x11\x15\x2f\xda\xe1\xb6\x81\xf5\xea\x3d\x38\x27\x9b\xbd\x5d\xfd\x5c\xef\x40\xfc\x2e\x05\x68\x4b\xde\x4b\x89\x16\xe5\xf7\x46\x4b\xbc\x06\xc4\x22\x37\x60\x5b\xdd\x58\x79\x04\x36\x5d\xc8\x6f\xff\xb4\x3b\xff\x39\x4b\x7c\x28\x2b\xc3\x6b\xb8\xc6\x20\x97\x0e\xea\x8c\x59\x53\x6e\xb2\xa2\x38\x9d\x4e\xab\xa8\xc8\xaa\xd4\x75\xe1\xa9\x0b\x54\x31\x2e\x1e\xa5\x18\x07\x02\x5e\x59\x9f\xaa\x4e\x29\x5b\x1a\x80\x86\x6c\xe2\xc5\x5d\xbc\x8f\xca\x6f\x2f\x57\x0d\xc6\x9a\xcd\xc7\x06\x23\x01\x7f\xe5\xf6\x03\xdf\x27\x76\x9a\xfa\xc0\xf1\x7d\x70\x01\x9d\x32\xbc\xd9\x03\xfb\x1a\xd7\xd8\xab\x4d\x60\xf0\x0b\xb8\xc8\x60\x31\x96\x90\x34\x8d\x25\x3a\x9a\xc4\x92\x9f\xfe\x9d\xcc\x16\xa3\xe7\x51\xd5\xa7\xa9\xc5\x4b\x27\xd1\xc2\xa9\x23\x6c\xcb\x9b\x9e\xe4\x80\x91\x05\x4d\xb2\xeb\x93\x2d\x6e\xee\x5c\x93\x5d\x09\x7f\x10\xd2\x2d\x84\xff\xcf\xb8\x9c\xe4\xdc\x8c\x1d\xf3\xe9\x26\xc8\x4a\xa6\x67\xfd\x3c\x63\x82\x3b\x9e\x2f\x0b\x12\xa0\xc0\xc1\x5c\x54\x38\xe3\xf4\x7e\xaf\x60\x93\xd5\x5a\x10\x14\x85\xb5\x88\x09\xcf\x11\x55\x2a\x69\xea\x9e\xc5\x76\xe6\x39\xc3\xed\x21\xd7\xc1\x7b\x23\xad\x31\x89\xd0\x46\x0b\xb1\x33\x0c\xe3\x20\xfe\x24\x5e\x48\x50\x32\x2a\x35\xc1\x47\xbf\x8a\xe2\x05\x5e\x4a\x0a\xc2\xc6\x91\x96\x08\x69\x3b\xd9\x08\xb8\xdf\x64\xf9\x6d\xc6\x8c\xa6\xeb\x09\x04\x5a\xbd\x9f\xc1\xaa\x67\x95\x87\x4d\x16\x26\xb6\x5e\x80\xd8\xb0\x45\x20\x0b\x8d\x1b\x7c\x3d\xa7\x38\x00\x6a\x65\xae\x85\x4a\xea\xba\xed\x4f\xa4\x2d\x82\x09\x43\x8f\x5c\xb5\xd6\x92\x8c\x9d\x16\xe7\x54\x42\xbb\x7d\x6d\x80\xd2\x9a\xd9\x2e\x0e\x4e\xbc\x71\xcc\x69\x16\x2c\xc2\xdc\x41\x5a\x12\xf3\x83\x87\xb4\x27\x09\xa9\xb4\x76\xe3\x8b\xec\x3a\xe7\x34\x86\xdf\xb9\x45\x73\x86\x49\x96\xc6\x25\xfe\xa1\x07\x2a\xde\x29\x17\xc3\x48\x48\x5b\xcb\x81\x65\xb6\xfd\x91\x37\x25\xa8\x75\x11\x0e\x8f\xde\xa6\x29\x1f\x6f\x23\x3f\xd4\x77\xd1\x54\x30\x7a\x89\x1e\x8f\x27\x0c\xa7\x49\x34\x55\x28\xe4\xe0\xa3\xe9\x82\x33\x77\x70\x7e\xc1\xbe\x3e\x72\xd5\x81\xf5\x80\xf3\x86\x88\xc0\x8e\xa9\x6a\xb0\x96\xe3\x00\x09\x7a\xda\x49\x4c\x72\x05\xc6\x31\xff\x3f\x27\xd4\x41\xbe\x3e\x9d\x9f\x5d\x2c\xd7\xbf\xe4\x4a\x5b\x98\x5a\xc8\x1f\xcc\xb6\xff\xfe\x57\x6a\x1d\xe2\xd3\x4b\x26\x71\xd3\x4c\x59\x1a\xa5\xf7\x75\xf4\x84\xd3\x7d\x97\xb7\x77\xdc\x42\xc8\xad\xaf\x84\x2e\xc9\xaf\xec\xe0\x6a\x85\x26\x0c\x3f\x54\x41\x60\x38\x47\x33\x7b\x66\xdb\xf7\xbf\xbd\xf9\xf5\x15\x29\x86\x4f\x4d\xab\xf0\x59\x1e\xa4\xac\x3c\xbe\x06\xaa\x70\xa2\x06\xc7\x59\x83\xc8\xbb\xc9\x08\x9e\x5b\x8d\x37\x64\x31\x85\x36\xd9\x49\x0a\x77\xd8\x08\x7c\x87\x4a\xc8\xfd\xe4\x05\x93\x8d\xa4\x82\x28\xb7\x25\xda\x63\x73\xdb\xe7\xa1\x92\xcd\x5d\x40\xba\x9b\x83\x73\xed\xab\xa2\xa8\x90\x8b\x5d\xed\xb5\x46\xe8\xe2\xad\xb4\xfe\x85\x2b\xad\xfd\xa1\xe2\xb5\x54\xe7\xcd\x3f\x5a\x68\xfe\xf0\x9e\x37\xf6\xd5\xcb\x6f\xbf\x7d\xf1\xc7\xf0\x27\x1d\x57\xb2\x7c\xf1\xb2\x1f\xdd\x20\x4c\xaa\xcd\x8d\x75\x67\x05\xe8\x72\x70\x37\x21\xb8\x6f\x28\x4d\x89\xdb\x4d\x2a\x3f\xc4\x7d\xbf\x95\xf9\xb3\xd9\xe5\x6c\x0f\xc5\x85\x75\xdc\xc9\x92\x68\x8a\x1d\xe6\x90\x45\x88\x6c\x57\xb5\x6c\x56\x74\xea\x7f\x64\xc8\xdb\x76\xc6\xe6\x09\xc7\xc8\x5a\x39\x3f\x81\xd5\x35\x8c\x54\xc1\x82\x2e\x3a\x78\x4d\xa0\x12\xd9\x7e\x95\xe7\xec\x8d\xbc\xc7\x72\xae\xe1\xc7\x1d\x37\x2c\xcf\x13\x40\x24\xc0\x75\xba\x1d\x52\x3f\xd2\x84\x9f\x1e\x00\xfa\x69\x45\x6c\x72\x22\x5f\x42\x9a\xb4\x66\x1d\xb0\x20\xd9\x8f\x4c\x66\x98\xfa\x08\x18\xc5\x63\xe1\x61\x9b\xbc\x72\x58\x46\x2b\xde\x0e\xf9\xd7\x3f\x74\xab\x78\x66\xd8\x4e\x85\x4d\x40\x5c\xa2\xda\x39\x12\xd3\xa3\x97\xa2\xf6\xff\x8f\x74\x8e\x98\x29\x66\x46\xd5\x77\x08\x55\xe2\x5a\xad\xe1\x9f\xc3\xcc\x03\x12\x65\xf0\xe4\xa5\x4e\x1e\x82\x59\xd3\xe1\xcd\xc1\xbe\x64\x9e\x75\xa7\x12\x55\x7a\x52\xfc\x99\x18\x91\xd0\x8c\x4a\xa9\xa3\x87\xd1\xd5\x6b\x3f\xfc\xc0\x77\x97\x4a\x2c\xf2\xc3\xf8\x1b\x2d\xf8\x48\xbf\xb4\x11\xcf\xe7\x65\x88\x40\xc8\xde\x69\x6e\xc4\xa5\x10\x59\x17\x78\x66\xc2\xa6\xc1\x07\xb1\xa7\x64\x67\x70\x43\x60\x8f\x24\xa1\xa2\x58\xb9\xc2\xa7\x41\x5b\xb2\xa0\xcd\x1e\x1e\x7a\xa4\xf7\xab\x59\x7c\x60\x2e\x7a\x15\x73\xbd\x94\xb4\x98\x78\xf8\x1e\x7e\x51\xad\x45\x99\x58\xb5\x3e\x41\xe6\xb4\x0c\x4e\xfc\xbc\x50\x64\xdb\x58\x65\x53\x31\xbd\xa4\xcc\xba\xe8\xd4\x93\x3c\xdb\x0f\x8d\xdc\x1f\xdc\xdc\xcd\x54\xf8\xbf\xc3\x84\x03\xf1\xb6\x99\x79\x37\x75\xe5\x34\x4c\x1b\x2a\xe2\x97\x94\x6f\x55\xd7\x2b\xff\x5a\x88\x50\x31\x2d\x19\x93\x4c\x19\x0f\x0a\xa3\x5b\xa1\x4f\x4d\x36\x33\xf7\x25\x90\xa6\xb4\xcb\x98\x31\x70\x8a\xa5\x64\x0f\x37\xdc\x48\x2c\xb9\x39\xb6\x5f\x6d\xd7\x22\x28\x9a\x0e\xe2\x22\xdc\x63\x06\x0b\x10\xa4\xbe\xa2\x74\xf1\x8d\xaa\x05\x83\x7d\x23\x97\x8a\x7a\xd5\x51\xf2\x97\xdc\x80\x1b\x32\x7f\x52\xfd\x4f\x3c\x31\xe8\x5a\x43\xd3\xcd\x2e\xf7\x05\x03\xdb\xd8\xbb\x46\x18\x88\xb3\x6b\x41\x19\x6c\x19\x2e\x6c\xa1\xe5\x86\x3b\x6d\x2e\x06\x93\xd8\x3e\x82\x07\xab\x2b\x47\x97\x75\xc0\x2a\x1b\xbb\xcf\xa0\xc1\x3b\x3f\xbe\x9a\x14\x93\x58\xa4\x95\x29\xd5\xb4\x93\x4b\x03\xb0\xd1\xee\xbf\x08\x42\xd4\x4f\x36\x83\x7a\xb2\x59\xd6\x6e\x49\xee\x58\xdf\x80\xa9\x84\x64\x05\x3d\x2d\x03\x70\x26\x88\xb3\x54\xc2\x5e\x7b\x1e\xfb\x2e\xe7\x6a\xf3\x41\x5f\x87\xee\xed\xe5\xeb\xd0\xc5\x0e\xc3\x77\x9c\xa4\x5f\xe8\x4f\x61\xed\x24\xd8\xb5\xaf\x37\x9f\x3f\x87\xa3\x54\xd5\xb5\x23\xae\x93\xbb\x8f\x8a\xc1\x58\x62\xaf\xae\x12\xf4\xb7\x48\x49\xae\x98\x22\xb4\x20\x03\x10\x8c\x1b\x92\x47\xeb\x88\xa5\xcf\x53\x6d\xa7\x54\xae\xa0\x1a\x63\xd6\x37\xa5\x6e\xcf\xdf\xb3\xbf\x61\x09\xc5\x3e\x80\x31\xe7\xcb\x07\xa6\x21\x48\xa8\xee\xb4\x58\x78\xee\xa5\x3b\x74\x3b\x5f\x6e\xde\x9d\xa9\x29\x41\xf2\xc2\x76\x95\x99\x63\x56\x20\x8d\xa8\xf5\x8b\x9f\x8c\xbf\x5e\x3d\x9d\x79\x81\x3d\x02\xf6\x1c\xd9\xf6\xad\xff\xbd\xc2\xe6\xb9\xaf\xb3\xfe\xc2\xcb\x3b\x6a\xfd\x70\x32\xa6\xfb\x1d\x43\x5c\xea\x86\xaa\xf7\xcb\x77\xac\xb8\x98\x86\xc7\x62\x8b\xb8\x2e\x82\xfd\xc3\x6c\xe4\xd1\xd8\xa3\x0f\xfe\x5c\xdb\xd2\xc8\xd6\x85\xef\x50\xfd\xd5\xf8\x47\x7e\x3f\xad\xd7\x69\x0d\x73\x6b\x67\x8b\x8f\x9f\x3a\xc0\xab\xde\xae\x6e\x6f\x57\xdf\xc5\x99\x2f\x54\x3f\x7a\xd4\x0f\x0c\xb7\x73\xee\x7d\x91\xfb\x71\x5a\x6b\xcf\xcf\xcd\x9f\xa5\x6b\x9c\xa8\xc8\x1e\x9d\x1f\x87\x7c\x1a\x78\xb1\x7b\xb9\xb7\xec\x28\xad\xdc\x29\x1a\xd2\xc1\xc5\x74\x8e\xc4\xb6\x1e\x88\x6d\xfd\x18\x71\x2d\x06\xe2\x5a\x3c\x46\xac\xf6\x03\xb1\xda\x27\xc4\x58\x3e\xfa\xe2\x1e\xab\x7d\xdf\xd7\x5d\xee\xf3\x9f\x00\x00\x00\xff\xff\xfe\x2d\xed\xc6\x90\x16\x00\x00")

func templatesBaseHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesBaseHtml,
		"templates/base.html",
	)
}

func templatesBaseHtml() (*asset, error) {
	bytes, err := templatesBaseHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/base.html", size: 7608, mode: os.FileMode(420), modTime: time.Unix(1455911903, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesConfigBaseHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x54\x5d\x6e\xdb\x38\x10\x7e\xf7\x29\x66\xb9\x01\xfc\xb0\x96\x18\x67\xf3\x14\x48\x5a\x60\x1f\xb2\x58\xf4\xa1\x40\xd3\x1e\x60\x2c\x8d\x6c\x3a\xd4\x4f\xc5\xb1\x13\x23\xc8\x39\x7a\x86\x9e\xa0\xef\xed\xc5\x3a\xa4\x2c\x2b\x76\x52\x34\x40\x00\x5b\x1a\x92\xdf\x7c\x33\xfa\xf8\x91\x0f\x0f\x50\x50\x69\x6a\x02\x55\x5a\x74\x2b\x05\x8f\x8f\x13\x00\x99\xee\xb0\x5e\x12\x9c\xdd\xd2\x6e\x06\x67\x5b\xb4\x1b\x72\x70\x95\x42\x7c\xed\x61\x12\x9f\xe0\x2a\x72\x0e\x25\x10\xc8\x80\x0e\x08\x80\xa4\x30\x5b\xc8\x25\xcb\xa5\x0a\x2d\x75\x0c\xe1\x19\x49\xae\x67\x17\x98\xca\x02\x50\xa0\x8b\x0d\x73\x53\x0f\xe8\xdc\x36\x8e\x14\x14\xc8\x18\x15\xc6\x55\xe6\x40\xa1\xb2\x1f\x5f\x12\xdd\xa3\x87\x64\xcf\x37\x74\x31\x94\xd6\x52\x3b\xeb\xfb\xa4\xba\x38\xf4\xbc\x8f\xc7\x68\xf2\x44\x07\x36\x6c\xc9\xeb\xf0\xf2\xf2\x02\x5d\x58\x9d\x24\x7f\x14\x4d\xce\xbb\x96\x60\xc5\x95\xcd\x26\x49\xff\x92\xaa\x2b\xc2\xa2\xef\x2a\x09\x64\xd9\xcd\xa7\xeb\x0f\xbe\x2e\x53\xd5\x5a\xe4\xb1\x48\x2c\x44\x89\xee\x41\x7d\x42\x45\x8c\x50\x63\x45\xa9\xda\x1a\xba\x6b\x1b\xf9\x58\xc8\x9b\x9a\xa9\xe6\x54\xdd\x99\x82\x57\x69\x41\x5b\x93\x53\x14\x06\x33\x30\xb5\x61\x83\x36\x72\xb9\x48\x93\xce\xf7\x62\x26\xd6\xd4\xb7\xb0\xea\xa8\x4c\xa7\x2b\xe6\xf6\x4a\xeb\x52\x58\x5c\xbc\x6c\x9a\xa5\x25\x6c\x8d\x8b\xf3\xa6\xd2\xb9\x73\xff\x94\x58\x19\xbb\x4b\xdf\xb7\x54\xff\x75\x83\xb5\xbb\xba\x3c\x3f\x9f\xfd\xdd\xff\x0d\xa3\x35\xf9\xec\x72\x88\xa6\xd0\x91\x4d\xa7\x8e\x77\x96\xc4\x07\xc4\x53\xf0\x1a\xa4\x53\xa6\x7b\xf6\x6c\xd3\xa7\xf5\xc3\x92\x1a\x96\x54\xc8\x55\x63\xae\xea\x1b\x54\xda\x31\xb2\xc9\x3d\x46\x2f\x9a\x86\x1d\x77\xd8\xc6\x95\xa9\x63\x9f\xf5\x46\x42\x6c\xdb\x67\x34\xaf\x48\xf3\x6a\x45\x78\x47\xae\xa9\xe8\xa8\x95\x44\x0f\xfb\x9b\x2c\x9a\x62\x97\x3d\x33\xb9\xdf\x2d\x14\xaf\x74\x0a\x4c\xd1\x0f\x65\xf3\x46\x93\x1f\x41\x6d\x74\xef\xa2\xf9\xc5\x61\x15\x8e\x7d\xb2\x3f\x94\xf1\xe0\xe8\x67\x80\x81\xfd\x29\xe4\x60\xfb\x43\xd8\xc7\xa5\x48\x4b\xdd\x50\xbb\x1f\xfd\xa2\xad\xe1\x0b\xc6\xb6\x92\x76\x58\xf5\xea\x47\xd5\x86\xa9\x80\x10\xe6\x52\xff\x08\xfa\xfd\x2b\xbc\x13\x69\xe1\x23\x75\xdd\x6e\x9c\xfd\x36\x72\xe1\x5e\x70\xef\x4c\x27\xd6\x5c\x1a\x5e\x6d\x16\xc1\x90\xb7\x92\xc9\x3e\x51\xbb\x4d\x29\xac\x89\x39\x34\x8c\x50\x62\xd4\x43\x65\x5e\x9b\x0c\xfe\x0b\x83\x44\x63\xf6\xb6\x32\x5a\xae\x16\xb9\xb2\x54\xf6\x7f\x78\xff\x96\xf0\x4f\x6e\x5a\x95\xfd\x8b\xb9\x38\xb2\x91\x5f\x7b\x94\x91\xe8\x36\x7b\x69\x2b\x7a\xc5\xf7\x23\x97\x77\xa6\x65\x70\x5d\x3e\xf6\x87\x6b\xbc\x3f\x3d\xa0\x7e\x4e\x5b\xb3\x70\x7a\xfd\x79\x43\xd2\xef\x3c\x9e\xcf\xe3\x8b\xfd\x28\x38\x73\xed\xbc\x1c\x3d\xe1\x0b\xec\x83\xab\xd7\xa7\x87\xeb\xf5\x79\xfe\x0c\x9d\xa2\xe5\xfe\x0d\xf6\x97\xf3\x10\x2e\xbe\xf1\xba\xfc\x19\x00\x00\xff\xff\xac\x4a\x10\x82\x55\x06\x00\x00")

func templatesConfigBaseHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesConfigBaseHtml,
		"templates/config-base.html",
	)
}

func templatesConfigBaseHtml() (*asset, error) {
	bytes, err := templatesConfigBaseHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/config-base.html", size: 1621, mode: os.FileMode(420), modTime: time.Unix(1455843986, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesLoginHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x53\xcb\x8e\xdb\x30\x0c\xbc\xe7\x2b\x04\xdd\xb5\x46\x5b\xa0\x27\x3b\xe8\xa5\xb7\x02\x5d\xa0\xfd\x01\x45\xa2\xbd\x42\x29\x51\x90\xe8\x6c\xb6\x8b\xfc\x7b\xe9\x47\xe2\x04\x45\x2e\xed\xc1\xb0\x28\xce\x90\x33\x84\xf8\xfe\xae\x3c\xf4\x21\x81\xd2\x1c\x18\x41\xab\xf3\xf9\x1b\x0d\x21\x49\x02\x92\x97\x68\x77\x03\x71\x94\x18\x12\x4f\xa0\x5d\xeb\xc3\x51\x39\xb4\xb5\x76\xba\xd0\xab\xde\xef\x94\x6a\xf3\xe5\x06\xc1\x7a\xc5\x70\x62\xe3\x84\x00\x65\x39\xc7\x91\xc1\xeb\x7d\x1b\xe2\xa0\x6a\x71\x9d\x6e\x2a\x5b\x0e\xae\x09\xd1\x0e\x50\x9b\x3a\xf6\xc5\x20\x0d\xf4\x54\x8f\x83\xe0\x9a\x3c\x57\xbd\xe9\xe4\x08\xcd\xa9\x9a\x0f\x1f\xd5\x74\x8a\xde\x7c\xbe\x1c\xa8\xef\x2b\xb0\xf9\x34\x0b\x11\x52\x4f\x25\x2a\xeb\x38\x50\xea\xb4\x78\x28\x70\x84\x52\xc5\x04\x4e\xf6\x26\x0b\x5a\x45\xe0\x17\xf2\x9d\x7e\xfe\xfe\xe3\xe7\xca\x9b\x98\x01\xd0\x4b\xad\xcb\xc5\xbd\x82\xa9\xae\x19\x0a\x8d\x59\x6f\x00\x81\xa0\x3d\x00\x2a\xc9\x76\x1a\xa2\x0d\xa8\x37\xc9\x89\x8b\x68\x9c\x01\x7a\xff\x75\x4a\xb6\xcd\x1c\xdd\x15\x08\x29\x8f\x7c\xd7\x65\x65\xaa\x39\x63\x70\xd0\x2a\xf8\x6b\x75\x7e\xcb\x70\x0d\x92\x8d\x5b\x90\xd1\x3a\x78\x21\xf4\x20\x5a\x7e\xbd\x21\x7c\x81\x93\x8d\x19\xe1\xc9\x51\xd4\xca\x8e\x4c\xce\xe6\xc0\x16\xc3\x6f\x61\x25\x4a\x70\x63\xa5\x6d\xc4\xec\x3f\x5b\xcf\x02\x7c\xa5\xe2\x1f\xb9\x7f\x5e\xf3\x8f\x07\x30\x59\xfc\xab\xca\x83\x71\x2c\x33\xd8\xd0\xcb\x18\xae\xf1\x7f\x99\x3a\x8c\xcc\x94\xd6\x16\x75\x3c\xc4\xc0\x57\x39\x07\x4e\x4a\x3e\x23\x7b\x61\x47\x64\xbd\x9f\x57\xa6\x6d\x16\xce\x83\xb6\x6d\x73\xff\xb0\x24\x96\xd6\xf3\x03\x5f\x60\xeb\x6f\xdb\xbc\x3f\x01\x00\x00\xff\xff\xb2\x60\x15\xab\x9d\x03\x00\x00")

func templatesLoginHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesLoginHtml,
		"templates/login.html",
	)
}

func templatesLoginHtml() (*asset, error) {
	bytes, err := templatesLoginHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/login.html", size: 925, mode: os.FileMode(420), modTime: time.Unix(1455844022, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesRegisterHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x56\x4d\x93\xdb\x36\x0c\xbd\xe7\x57\x60\x74\xae\xd6\xdd\x9e\xd7\x9e\x4e\x3b\xdb\x63\xeb\xe9\xa6\xb9\x53\x22\x24\xa1\xa6\x48\x0d\x09\xd9\xeb\x64\xf2\xdf\x0b\x52\x1f\xfe\x4c\xe2\x6d\x9d\x5e\x34\x22\xf9\x08\x02\x0f\x78\x24\x3e\x7d\x02\x8d\x15\x59\x84\x8c\x89\x0d\x66\xf0\xf9\xf3\x0b\x72\xdf\xc9\x02\x5a\x2d\xa3\x77\x47\x90\xd2\x59\x46\xcb\x11\xf4\x0e\xe0\x49\xd3\x16\x4a\xa3\x42\x58\x66\xde\xed\xb2\x95\xcc\x9d\xce\x96\xce\xe4\xaf\x21\x7f\xfc\x69\x5c\x93\xd5\x6e\x5a\x33\xa8\x34\x30\xbe\x72\x5e\x8a\x49\xf4\xc3\x7f\xdb\x33\xea\x6c\xf5\x44\x6d\x0d\xc1\x97\xcb\x6c\x11\x58\x31\x95\x0b\x6a\x55\x8d\x61\x11\xfa\xca\xe7\xc6\xd5\xee\x21\x6c\x6b\xc1\x2d\xba\x4b\xcb\xca\xa0\x67\x48\xdf\x9c\x6c\xe5\x20\x1e\x95\xad\xde\x37\x14\xc0\x38\xb7\x91\x2f\x6d\x10\x14\x58\xdc\x01\x59\x39\xc0\x96\x08\xae\x82\x97\xbf\x7e\xfb\xf3\x01\xd6\x02\x0f\x08\x1e\x6b\x0a\xd1\xaf\xbd\xeb\x3d\x28\xdd\x92\x85\x3e\xc8\x58\x09\x2d\x42\x44\x45\x75\xef\x11\xb8\x11\x4b\x5d\x67\xa8\x14\x37\x9d\x85\x02\x8d\xdb\x3d\x1c\xfb\x55\x39\xdf\x4e\xae\xc5\xff\xbc\x71\x9e\x3e\x0a\x93\xca\x64\xa0\xca\xb8\x6b\x99\x09\xc9\x1e\xb7\xe8\xc3\xc0\xb2\x18\x8f\x24\x67\xd0\x22\x37\x4e\x2f\xb3\xf5\x1f\x2f\xef\x67\x12\xa3\x51\x42\xa3\x03\xf2\x61\xea\x94\xf9\x74\x50\xed\x5d\xdf\x65\xc7\x10\x01\x19\x25\x2e\x82\xac\x2f\x33\x6c\x15\x89\x0f\xe7\xc9\x82\xf8\xd7\xea\x3c\xfe\x58\xf6\x32\x48\x7b\xb2\xd5\x73\xc4\x3f\x2d\xd2\xe8\xcc\xea\xb5\xa4\x4f\x76\x1e\x7f\x3c\xf3\x41\xf0\x64\xbb\x9e\x4f\x9c\x1d\xcf\xca\x80\xf4\xec\x19\xef\x3b\x9c\x07\x56\xb5\x87\x41\x67\x54\x89\x8d\x33\x1a\x25\x8e\xcd\xde\xe0\xcf\xf8\xaa\xda\xce\xe0\x43\xe9\x5a\xa1\xb5\x67\x57\xaa\x8e\x84\x63\xfa\x28\xbb\xac\xb3\x78\x4e\xc4\x42\x7c\x3e\xa1\xef\x62\xe2\xad\x7c\x76\x02\xdd\x39\xaf\xdf\x40\xe9\x7a\xdc\x72\x4f\x56\x23\x81\x17\xae\x9c\x72\x3c\x10\x7b\x00\x0d\xdc\xce\xe3\x6f\x30\x25\x53\xa1\x53\x36\x1d\xd4\xa0\xe9\x7e\x31\xae\xdc\x7c\x39\x68\x57\x55\x52\xa9\x12\x7b\xc4\xe6\x45\x02\xaf\x7e\x77\x10\xe4\x9a\xa1\xa4\xa5\xc0\x5e\x91\xe5\xf0\x03\xfc\xdd\x07\x8e\x32\x13\x75\xd6\xce\x69\x08\xf1\x6e\x88\xf2\x14\x15\x42\x29\x67\x7a\x6c\xb1\x2d\xd0\x8b\xc4\xa2\x0f\x77\x4b\xe0\xcd\x19\xfb\x40\x81\x0a\x32\xc4\xfb\xfb\xe4\xec\xf8\x12\x55\x9a\xdc\x05\x62\x72\xf1\x72\xfe\x24\xe3\xdb\xd9\xaf\xbc\xeb\x0b\xb9\x8f\xa6\x2c\x0f\x56\xc7\x14\x1f\x50\x19\x6c\x95\xe9\x63\xd6\x47\x74\xd9\x60\xb9\x41\x7d\xed\x98\x75\x82\x5c\xfa\x75\x8d\x80\xab\xf5\xf2\x9d\xe2\xf4\xb4\x55\x8c\x37\x07\x3a\xc2\xaf\x46\x38\xac\xfd\xb7\x10\xbf\x22\x8a\xe3\xd2\x7f\xb6\xaa\x30\x08\xa3\x3b\xf3\x0b\x14\x80\xaa\x54\xe6\x3b\x65\x19\xd8\x41\x43\x1a\x21\x3e\x0a\x7b\x6e\xc8\xd6\x50\x79\xd7\xa6\xf7\x66\xc8\xd8\x15\x05\xdc\xfb\x56\xbb\x59\x14\xbf\x0e\x8d\x01\x3c\x8b\x34\xb5\x16\x67\x6f\xd6\xc6\xe3\x8d\xda\x48\xd5\x59\xb8\xd7\x7f\x5b\x36\xf1\xd2\xd0\x5d\xe3\xd8\x85\xa9\x5e\x66\x93\xf3\xd3\x72\x80\x5c\xb3\x96\x82\x83\x75\x02\xdc\x4d\x0b\x77\x89\x6b\x2b\x95\xf2\x8d\xb8\x46\xc8\x97\xe3\xfa\x90\x00\xff\x8f\x00\xa6\x2a\x81\xb1\xa1\x84\x1d\x19\x23\x1d\x9b\xb4\x4e\xd0\x51\xc9\xd2\x58\x85\xd4\x67\x0d\x5e\x47\x31\x14\x08\x49\xd2\x26\x0a\x66\xe8\xc9\x62\xbb\x06\x15\xa2\xfe\xae\x4a\xf8\xda\x7d\x3e\xbd\x6c\x97\x95\x5b\xf4\xcc\xd2\x0d\x0e\x09\x09\x7d\xd1\x12\xcf\x54\x14\x2c\x6d\x22\xdb\x5c\xba\x6a\xd5\x1b\x4e\xea\x19\xba\xc9\xa7\xc5\xb0\xef\x6d\x91\x3c\x2d\xce\xbb\x41\x99\x91\x78\xc6\x7e\x7c\x02\x8f\x3f\x87\xc6\xfe\x9f\x00\x00\x00\xff\xff\xb1\xbd\x81\x87\xfc\x0b\x00\x00")

func templatesRegisterHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesRegisterHtml,
		"templates/register.html",
	)
}

func templatesRegisterHtml() (*asset, error) {
	bytes, err := templatesRegisterHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/register.html", size: 3068, mode: os.FileMode(420), modTime: time.Unix(1455844119, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSettingsHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x57\x4f\x73\xeb\x26\x10\xbf\xbf\x4f\xb1\xa3\xe9\x55\x4e\xd3\x77\xeb\x38\x3e\xa4\x79\x6d\xdf\x29\x99\x38\xcd\x1d\xc4\xda\xa2\x41\xa0\x02\xb2\xe3\x7a\xfc\xdd\xbb\x20\xc9\x96\x64\xc7\x7f\x9a\xbc\xc9\x21\x0e\xec\xfe\x80\xdf\xfe\x15\xac\xd7\x20\x70\x26\x35\x42\xe2\xa5\x57\x98\xc0\x66\xb3\x5e\x8f\x9e\xc2\x38\x8c\x00\xb5\x20\xd1\x97\x0e\x2e\x33\xda\xa3\xf6\x49\x23\xfe\x69\xae\x0c\x67\x0a\x7e\xbd\x81\xd1\x14\xbd\x97\x7a\xee\x5a\x95\x6b\xe7\x5d\xe5\x3d\xff\x1b\x33\x1f\x20\x63\x21\x17\x90\x29\xe6\xdc\x4d\x62\xcd\x32\x99\x7c\x01\xe8\xca\x32\xa3\xd2\x57\x97\x5e\xff\x12\x35\x7d\x5d\x8b\x1f\xca\x87\x6b\x86\xfa\x92\x69\x54\x10\x7f\x53\xa9\x67\xa6\x03\x3b\x00\x4c\x73\x64\x82\x48\xf7\x50\x84\xcb\xbf\xf6\x61\xb5\xef\x26\xb7\x95\x54\x02\xbe\xd3\xb6\xb6\x60\x5e\x1a\x3d\xbe\xca\xbf\xf6\x0e\xb8\xa2\x13\x8e\x9f\xc8\x8d\x58\x0d\x8f\xf3\x8c\x2b\x6c\x61\xf5\x24\xfe\xa6\x14\x0b\x81\xda\xa1\x18\xac\x08\x6b\xec\x50\x14\x84\x62\x32\x76\xde\x1a\x3d\x9f\x3c\xa3\x75\x91\x61\x33\x1f\x5f\x91\xf2\xe0\x8a\xf5\xba\x89\xf1\xa8\x59\xb3\xd9\x1c\x02\x93\x6c\xef\xc8\x93\x2c\x1e\x71\x21\xcf\xa3\xd1\x9a\xcf\xd4\x92\xad\x5c\xba\xb4\xac\x4c\x3a\xd4\xa2\xeb\xff\x90\xfe\x4f\xe6\xf2\x0f\xe4\xf7\x24\x0b\xbc\xc8\x45\x91\x47\x58\xf5\x81\x24\xee\x98\x67\x70\x27\x2d\x95\x8d\xb1\xab\x8b\xe8\x84\xa5\xb4\xf2\x5c\x32\x24\x09\x79\x75\x24\x65\x7b\xd3\xce\xa4\x3b\x0c\xe9\xdf\x86\x2b\x8c\xd3\xdc\x58\xf9\x2f\xb5\x0d\xa6\x12\x60\x59\x28\x8c\x9b\x84\xfa\x83\xc5\x05\x25\x14\xb5\x94\xb6\x4f\x84\x9e\x92\x40\x81\x3e\x37\xe2\x26\x79\xb8\x9f\x3e\x1d\xac\xf1\xb8\xe7\xdc\x9a\xaa\xec\x16\xb9\x62\x9c\x2a\x7b\xd8\x06\x20\x8c\x0a\x91\x86\x81\x26\xbf\xa9\x34\xe2\x92\xc9\x33\xe5\x1d\x97\x4a\x7a\x72\x68\x14\x1d\xee\x17\x7b\x1b\x5d\xff\xfc\x66\xcb\xb0\xd4\x2b\xcc\xb0\x76\x07\x7b\x37\x52\xa9\xcb\xca\x83\x24\x2b\x17\x5b\x1e\x69\x59\x71\x25\xb3\x04\xfc\xaa\xc4\x76\x37\xd0\xac\xc0\x2e\x2a\x81\x05\x53\x15\x89\x5a\x34\x39\x52\xce\x00\xff\xd9\xb5\xdb\xd1\xce\x36\xd8\xc2\x36\x9b\x2c\xc7\xec\x05\xc5\x7a\x4d\x3d\x7d\xb3\x99\xc0\x80\xd3\x43\x04\x0e\xd2\x61\x8f\xfd\xd1\x0e\xf6\x6e\x07\x58\xb9\x60\x1e\xcf\xf6\x40\x0b\x3f\xe5\x82\x16\xb7\xe7\x83\xa1\x0b\x6a\xe0\xe5\x3e\x70\xd4\xbb\x5b\x27\xe4\xa8\xca\x94\x2b\x93\xbd\x24\x93\x6f\x3a\x36\xeb\x86\x00\x48\xed\x3c\xd3\x19\xba\x40\x77\x65\x2a\x58\x32\xed\xc1\x1b\xc8\xa5\x40\x08\xd5\xb0\xf2\x39\xd1\x87\x99\x35\x05\xf8\x9c\x56\xc6\xa0\x8c\xa8\xe6\xe9\x84\xd3\x55\xf8\x03\xea\xe4\xb7\xfa\x73\x0f\xdf\x0a\x8e\x22\x7c\x0b\xcf\x2a\x97\xeb\x33\xca\x25\xc6\x82\x9b\xd7\x4b\x13\x06\x03\x93\x32\x37\xde\xb8\x36\x53\xb6\x5b\x35\xc9\xd2\x83\xd4\xd9\xb1\x4b\x8d\x68\xc9\x43\x54\x9e\xcc\x88\x88\x85\x1a\xfc\xae\xd2\x78\x97\xb1\x0b\x4a\x8f\x13\xc6\xb6\x90\x83\xc6\x3e\x47\xe5\x99\xc6\xd6\xe0\x0f\xac\x81\x36\x71\xa0\xb9\x39\xc2\x52\x2a\x05\x4c\x29\xb3\x84\x52\x66\xbe\xb2\x54\x10\x8c\xee\x99\xb5\x0d\xa1\x1e\x38\x42\xac\x77\x15\x6a\x26\x54\x8a\x85\xe9\x5f\xbf\x3f\xc2\x0c\x51\x7c\x66\x31\x7c\x2f\x4a\x63\xfd\xfb\x2a\xa0\xec\x31\x6a\x0f\xa0\xc6\xe0\xa9\x4d\x0f\xb2\x83\x41\x6e\x71\x36\xf8\x58\xe6\xd2\xf3\x8a\xc2\xe8\x53\x19\xe9\xc4\xaf\x66\x43\xad\xee\x1c\xd3\x2d\x64\x7c\xc5\xfa\x51\x2b\x3f\xc5\x6f\xb7\x2c\x7b\xa9\x4a\xf7\xb9\x8e\x13\x74\x17\xe2\xcc\x61\xca\x23\x9b\xda\x6d\x35\x33\xb8\x6b\x74\xc7\xfc\x75\x24\xc7\x1f\xd1\xd1\xd5\x0c\x77\x7d\xfb\xee\x16\xea\x53\x80\xaf\xc8\x9e\x72\x15\xf2\x3f\x28\x66\x32\xe6\x34\xe5\xf8\x38\x33\x02\x0f\x5f\xd5\xa2\xe6\x7f\xe6\xf9\xfe\x03\xe8\xc2\xd0\xbe\x79\xf9\x71\x45\x27\xc4\x66\x36\xa3\x1e\x93\xd6\x73\x35\xdf\xce\xfb\x21\xe3\x95\xf7\x46\x37\x7d\xcb\x55\xbc\x90\x94\xad\xcd\xe6\xdc\x6b\xa0\xbf\x94\xde\x95\xac\x52\x3e\x99\x4c\xd9\x82\xfc\x5f\x2f\x99\x50\x04\x77\xb8\xe4\x40\x34\x2b\x1b\x1e\x70\x02\x5f\xeb\x38\x66\xe1\xbb\xaa\x7a\xe1\x7b\xfb\xaa\x1a\x0c\x8f\x8f\xcd\x5a\xd8\xfc\xdb\xbd\x76\xff\x0b\x00\x00\xff\xff\x75\x5f\x3f\x21\x16\x0f\x00\x00")

func templatesSettingsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesSettingsHtml,
		"templates/settings.html",
	)
}

func templatesSettingsHtml() (*asset, error) {
	bytes, err := templatesSettingsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/settings.html", size: 3862, mode: os.FileMode(420), modTime: time.Unix(1455435507, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesShitbucketImportHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x52\xc1\x8e\x1b\x21\x0c\xbd\xe7\x2b\x2c\x2e\x3d\x4d\x50\x7b\xac\x32\x39\x54\xbd\x54\xaa\xd4\xaa\xbb\xfd\x00\x06\xbc\x3b\x68\x19\x8c\xc0\x6c\xb3\x8d\xf6\xdf\x6b\xc8\x74\x92\x1c\x46\x63\xfc\xcc\x7b\x7e\xc6\xe7\x33\x38\x7c\xf2\x11\x41\xb1\xe7\x80\x0a\xde\xdf\xcf\xe7\xfd\x63\x8b\x5b\x04\x18\x9d\xa4\x76\x37\x75\x96\x22\x63\xe4\x56\xb9\x3b\x38\xff\x0a\x36\x98\x52\x46\x95\xe9\x8f\x3a\xee\x00\x6e\x73\x96\xc2\x70\x2a\xc3\xc7\x4f\x1d\x11\xec\x89\xf2\x02\xc6\xb2\xa7\x38\x2a\x21\xcd\xf8\x8a\xb9\x08\x6b\x99\x3d\x4f\xd5\xbe\x20\x0f\x7e\x49\x94\x3b\xbf\x82\x05\x79\x26\x37\xaa\x9f\x3f\x1e\x1e\x57\x8e\x7b\x85\x46\x38\x3c\x67\xaa\x69\x83\xa5\x20\x98\x09\x03\x08\x36\xaa\x9a\x83\xba\xb6\x13\x39\x4b\x4b\x1d\x56\xc7\x87\x4d\x13\x7e\xff\xfa\x7e\xd0\x3d\x7d\xc3\xe2\x63\xaa\x7c\x27\xb4\x12\x28\xf0\xae\x33\x0f\xa5\x4e\x8b\x97\x5e\xf9\x2d\xe1\xa8\x18\x4f\x12\x47\xb3\xe0\xaa\x9b\x82\xb1\x38\x53\x70\x28\x9d\xcc\xcc\xa9\x7c\xd6\xfa\x6a\x75\x8f\x27\xb3\xa4\x80\x7b\x4b\x8b\x36\xc9\x6b\xb9\x54\x14\x98\xca\x64\xe5\xc8\x26\xf8\xbf\x42\x15\x29\xe2\xad\xbb\x92\x4c\xec\x1d\xcc\x18\xd2\x97\x40\xf6\x65\x73\xd8\x32\xc3\xd4\x53\xc7\xaf\x14\x3f\x70\x1b\xc2\xb3\x18\x64\x02\xe3\x1c\x98\xf8\x06\x93\x29\xde\x36\x91\x19\x6c\x46\x57\x1a\xc6\x33\x5e\x66\xd0\xb8\xb7\x39\x6b\x19\xf4\x76\x98\x2a\x33\xc5\xd5\xe9\x7f\xdf\xab\xec\xc4\x11\xe4\x1b\x64\x47\x4c\x0d\xac\x8e\xdf\xfa\x1b\x1e\xf4\xe5\xd2\xfa\xf8\xba\xcd\xb0\xaf\xc8\x85\x78\xfd\x5d\xb7\xec\x5f\x00\x00\x00\xff\xff\x99\x01\x3c\x28\x8e\x02\x00\x00")

func templatesShitbucketImportHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesShitbucketImportHtml,
		"templates/shitbucket-import.html",
	)
}

func templatesShitbucketImportHtml() (*asset, error) {
	bytes, err := templatesShitbucketImportHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/shitbucket-import.html", size: 654, mode: os.FileMode(420), modTime: time.Unix(1455259370, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTagIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x50\xbb\x4e\xc4\x30\x10\xec\xf3\x15\xab\x15\x05\x14\xc9\x09\x4a\x94\xa4\x81\x06\x09\x5d\x81\xe0\x03\x4c\xbc\xe7\xb3\x64\xd6\x92\xbd\xe4\x0a\x2b\xff\xce\xda\xba\x88\xab\xac\x79\x78\x66\xec\x52\xc0\xd2\xc9\x33\x01\x8a\x97\x40\x08\xdb\xf6\x69\x5c\x56\x9e\xd8\x2a\xe8\x6e\x1c\x4b\x64\x21\x96\xea\xe9\x46\xeb\x57\x58\x82\xc9\x79\xc2\x14\x2f\x38\x77\x00\xb7\xdc\x12\x43\x1f\x5c\xff\xf8\xd4\x14\x80\x52\x92\x61\x47\x70\x27\xc6\xc1\xf3\x04\x43\x6d\xd1\x9c\xaa\x8d\x66\xbf\xf5\x2d\x8c\x70\x4e\x74\x9a\x50\x7b\x13\xad\x94\x72\x9d\x66\x5c\xbf\x7a\xba\x20\xa0\xb7\xd8\x32\x86\xb7\x57\x9d\x81\xb3\xda\x1a\x3c\x9a\x1f\x52\x02\xee\x4b\x69\xf8\xeb\xe3\xfd\x25\xfe\xb2\x6c\xdb\xc3\x78\x30\xfb\x06\x0a\x99\xae\xa5\xc7\x28\x67\xcf\x6e\xd8\x15\xb6\x4d\x18\x0f\xfa\x88\xb9\xbb\x1e\xff\xdf\xf0\x17\x00\x00\xff\xff\x1d\xd1\x06\x3d\x29\x01\x00\x00")

func templatesTagIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesTagIndexHtml,
		"templates/tag-index.html",
	)
}

func templatesTagIndexHtml() (*asset, error) {
	bytes, err := templatesTagIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/tag-index.html", size: 297, mode: os.FileMode(420), modTime: time.Unix(1455259161, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesUrlEditHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x54\xc1\x6e\xdb\x30\x0c\xbd\xf7\x2b\x08\x61\x57\x27\xe8\x8e\x43\xe2\xc3\xd0\x6e\x18\x30\x60\xc3\xda\x7d\x00\x2d\xd1\xb1\x30\x59\x32\x24\x3a\x4b\x5a\xf4\xdf\x47\x29\xae\xe7\x36\x1d\x8a\xa1\x87\x20\x12\xf9\xa4\xf7\xf8\x44\xfa\xfe\x1e\x0c\xb5\xd6\x13\x28\xb6\xec\x48\xc1\xc3\xc3\xb5\xb1\x6c\xfd\x0e\x24\xb7\xfa\xf9\xe3\xeb\xea\x36\x27\x24\x2e\x7b\xf2\x46\x16\x17\x8b\x53\x3a\x78\x26\xcf\x6a\x0a\xbf\x1b\xa3\x83\x0f\xdb\x72\x30\x87\x36\xc6\xee\x41\x3b\x4c\x69\xab\x62\xf8\xad\xea\x0b\x80\x65\x4c\x07\x57\x1d\x52\x75\xf9\xbe\x64\x24\xd7\x86\xd8\x3f\x26\xf3\xba\xea\x42\xb4\x77\x42\x82\x4e\x01\x6a\xb6\xc1\x6f\x95\x10\x45\xda\x53\x4c\x22\x40\x08\xab\x84\x7b\x51\xae\xac\x51\x45\xc0\xea\xcb\x95\x70\x2b\xe8\x89\xbb\x60\xb6\xea\xfb\xb7\x9b\xdb\xe9\xfe\xa7\xec\x85\x60\x17\xc3\x38\xcc\x69\x01\x38\x6c\xc8\x81\xe4\xb6\x8f\x9e\x3c\x13\x0b\xb9\xe8\x28\x9b\x82\x54\x75\x31\x68\xb3\x2e\xbb\xc5\x3d\x2f\x94\x79\xb9\xe0\x11\x84\xf5\xc3\xc8\x60\xcd\x73\xa2\xa2\x6b\x22\x51\xc0\xc7\x81\x04\x41\x07\x71\xd9\x63\x4f\x33\x7a\x8f\x6e\xa4\xec\x46\x29\xba\xa8\xc8\x65\x0f\x0e\x35\x75\xc1\x19\x92\x0a\xae\x0f\xd8\x0f\xf2\x7c\xf9\x3d\x0a\x62\x59\xe9\x5a\x24\xce\xbe\x3c\xd9\xfc\x97\x49\xb8\x4b\xaf\x7b\x24\xa0\x37\x5a\xb4\xa4\x79\xd5\xa1\x02\x9e\x0d\x3a\xb5\xc5\x67\xe2\x2c\xe3\x53\x88\x57\x36\x89\x4b\x47\x38\xb3\x4b\x87\x5e\xf8\x28\x56\x49\x5b\xf2\x9a\xc0\x09\x12\x72\xd4\x3a\xe9\x37\x69\xc1\x91\x83\xc6\xc1\x4a\x3f\xda\x3b\xb9\xdb\x07\xff\x4f\x4b\x65\x9b\x06\xf4\x45\x7d\x47\x6e\xf8\xe8\x82\xfe\x75\xe6\x54\xb6\xaa\xac\x42\xdb\x26\x62\xb1\x2e\x63\xab\xa6\x80\x8b\x6f\xd0\x8f\x89\xa1\x21\x48\x34\x60\x44\x26\x03\xcd\x11\x10\xe4\x72\x2d\x7d\x97\x39\x5e\x7e\xc5\x66\x64\x0e\x7e\xb2\x27\x8d\x4d\x6f\x79\xa6\x6f\xd8\x83\xfc\x2a\x19\x64\x1c\x1d\xab\xfa\x46\x86\x68\xb3\x3e\x1d\xa9\x37\xb8\xc0\x29\xe8\x22\xb5\xe7\x63\x67\xbd\xa1\x43\x9e\x7c\x55\x6b\x14\xb3\xdc\x66\x8d\xd3\x18\xaf\xf3\x03\x95\x61\x3f\x09\x9a\xfe\xfe\x7e\x42\xfe\x04\x00\x00\xff\xff\xf9\xb1\xfb\x6f\x79\x04\x00\x00")

func templatesUrlEditHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesUrlEditHtml,
		"templates/url-edit.html",
	)
}

func templatesUrlEditHtml() (*asset, error) {
	bytes, err := templatesUrlEditHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/url-edit.html", size: 1145, mode: os.FileMode(420), modTime: time.Unix(1455259223, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesUrlIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x51\x4d\xab\xdb\x30\x10\xbc\xfb\x57\x2c\xe2\xd1\x9b\x1d\xda\x63\xeb\xf8\xd0\x42\xa1\xd0\xf6\xd0\xbe\x16\x7a\x54\xe5\x75\xbc\x44\x96\x8c\xb4\x4e\x02\xc6\xff\xfd\xad\x3f\x63\x42\x4e\xf6\xcc\x7e\xcc\xcc\xaa\xef\xa1\xc4\x8a\x1c\x82\x62\x62\x8b\x0a\x86\x41\x38\xaa\x20\xfb\x16\x5f\xf5\xe9\x2f\xe1\x75\xa6\x32\x41\xd9\x4f\xdd\xe0\x0c\xd1\x46\x1c\x86\x1f\xe4\xa8\xd1\x96\x22\x93\x81\xcf\xde\x9f\x1b\x1d\xce\xe4\x4e\x63\x83\x2b\x97\xce\xe9\x27\xd9\x29\x19\xef\x18\x1d\x8f\x5a\x49\x5e\xd2\x05\x8c\xd5\x31\x1e\x55\xf0\x57\x55\x24\x00\x7b\xce\x78\x9b\xde\x62\xfa\xfe\xc3\x54\x01\xe8\xfb\xd1\xdb\x17\xdf\x39\x96\xe9\x99\x81\x97\x88\xcc\x22\x1b\xe1\xe3\x11\xb2\xdf\x2b\xd8\x1a\x82\x76\x27\x84\x97\x2e\xd8\xa9\xe1\xcf\xaf\xef\x5b\x71\x9a\x67\x6c\x5a\xab\x59\xac\x49\x4b\xda\xea\xc0\xa4\xad\x02\x87\xd7\xc9\xea\x8d\x41\xc9\x8c\x9a\x37\xa8\x75\xbf\xda\xe9\xde\xbd\x2c\x71\x47\xf4\x2c\xdb\x23\xff\x98\x6f\xaa\xb7\x6b\x75\xd4\x4e\x9b\x8e\xb1\x54\xc5\x57\xc9\x5c\x8e\x0a\x73\x7a\x11\x79\x57\xa3\xb5\xd4\x7e\xca\x0f\xed\xb6\xfa\x20\xbb\x67\xb0\xfb\xed\xfb\xf9\xb9\x92\x27\xdb\x8d\xbc\x04\x06\x55\xbc\xd6\x18\x10\x28\x82\xf3\x5c\x4b\x26\x98\xf0\xff\x8e\x21\x92\x45\x67\x30\x83\x7f\xbe\x03\xa3\x1d\x50\xd3\xfa\xc0\x50\x6a\xd6\x50\x05\xdf\x80\x4c\x60\x80\x88\xe1\x42\x06\x23\x90\x03\x21\x20\xd7\x50\x07\xac\x8e\x4a\x3c\x07\xbc\x60\x88\x72\xe0\xb8\x1d\x6f\x18\x54\xb1\xa2\xfc\xa0\x8b\x6c\x8b\xb1\xbf\xe2\x92\x62\xf9\xdc\x2b\x6f\x01\x00\x00\xff\xff\x90\xb0\x4c\xc4\xb9\x02\x00\x00")

func templatesUrlIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesUrlIndexHtml,
		"templates/url-index.html",
	)
}

func templatesUrlIndexHtml() (*asset, error) {
	bytes, err := templatesUrlIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/url-index.html", size: 697, mode: os.FileMode(420), modTime: time.Unix(1455259072, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesUrlNewHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x54\xc1\x6e\xdb\x30\x0c\xbd\xe7\x2b\x08\x5d\x76\x72\x8d\xf6\x38\xd8\x39\xec\x3c\x60\xc3\x96\x1d\x76\x94\x65\x3a\x16\x26\x4b\x82\x44\x75\x69\x83\xfe\xfb\x28\xc5\x49\xed\xa5\xc0\x90\xa1\x07\xc3\xa4\x48\x3d\x3e\x3e\x89\x3a\x1e\xa1\xc7\x41\x5b\x04\x41\x9a\x0c\x0a\x78\x79\x39\x1e\xef\x76\xd9\xce\x16\xa0\xed\x79\x69\xb3\xc8\x53\xce\x12\x5a\xca\x99\x9b\xa6\xd7\x8f\xa0\x8c\x8c\xb1\x15\xc1\xfd\x16\xdb\x0d\xc0\x72\x4d\x39\x53\x1d\x62\x75\xff\x50\x22\x1c\x1b\x5c\x98\xce\xc1\x6c\x57\xa3\x0b\xfa\x99\x11\xa5\x11\x20\x15\x69\x67\x5b\xc1\xc5\x02\x3e\x62\x88\x5c\x2d\x05\x53\xc5\xd4\x4d\xba\x14\x14\x30\x21\x8d\xae\x6f\xc5\xd7\x2f\xdf\x77\x33\xe8\xba\x64\x41\xdd\x07\x97\xfc\x25\xcc\x09\x46\x76\x68\x80\x63\x6d\x46\x14\x4b\x7e\x53\x5f\xdd\xc3\x99\x68\xb6\x2c\x05\xf6\xca\x0e\xb1\xfd\xf1\xed\x73\x53\x17\x7b\x81\xf6\x57\x87\x19\xe1\x15\xe2\x61\x51\x97\x73\xb5\xf5\x89\x56\xe4\xe6\x0a\x02\x74\xdf\xae\xfa\xa3\x27\x8f\xad\x20\x3c\xb0\x6d\xe5\x84\x33\x57\x6f\xa4\xc2\xd1\x99\x1e\x99\xfd\x48\xe4\xe3\xc7\xba\xc6\x83\x9c\xbc\xc1\x3b\xe5\x26\x16\x2e\x91\x53\xd2\x6b\x56\x51\x3f\xf3\x36\xeb\x2c\x2e\xbb\xaf\x99\xf0\x45\xab\x95\x73\x8b\x70\x24\xf7\xf1\x06\xe5\x76\x9c\xfe\x1e\xd2\x65\x95\x56\xa5\xd7\x22\x5e\x8b\x76\x4a\x5e\xa9\xc6\x2a\x31\x14\x86\x2a\x2a\x8d\x56\x21\x18\x1d\x3d\xe4\x55\x6d\xf8\x9e\xdd\xa6\x20\xbb\xd1\x4b\x5b\x88\x8d\x68\xfc\x27\xe3\xd4\x2f\x71\x75\xe5\x61\xee\xcf\x0d\x43\x44\x62\xa5\x72\x6e\xd5\x95\xe4\x22\x0e\x4c\x29\x12\x74\x08\x11\xbd\x0c\x92\xb0\x87\xee\x09\x24\x30\xb8\xc2\xa6\xce\x35\xfe\xf3\xd0\xde\x9a\xc0\x42\x27\x4e\x6f\x11\xcb\xbe\xd9\x5f\xfc\xf5\x21\x74\x89\xc8\xd9\x59\xe6\xf3\x4d\x9d\xc1\x3b\xb2\xc0\x5f\xc5\x4f\x83\x4c\x86\xc4\xf6\xa7\x4b\x1f\x02\x37\x34\x6a\x22\x6d\xf7\x3c\xad\x4d\x7d\x02\xd8\x36\x72\xb1\x4b\xc0\x18\x70\xb8\x9e\x74\x6d\x7b\x3c\x94\x41\xdf\x2a\xc9\xe7\x64\x9a\x5a\xfe\xfb\x1e\x37\x75\x96\xa1\xbc\x3d\xa7\xc5\xf9\xf7\xfa\x7c\xfd\x09\x00\x00\xff\xff\x90\xaf\x8c\x9d\xe7\x04\x00\x00")

func templatesUrlNewHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesUrlNewHtml,
		"templates/url-new.html",
	)
}

func templatesUrlNewHtml() (*asset, error) {
	bytes, err := templatesUrlNewHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/url-new.html", size: 1255, mode: os.FileMode(420), modTime: time.Unix(1455259279, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesUrlViewHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x8d\x4b\x0a\x02\x41\x0c\x44\xf7\x9e\x22\x64\xef\x9c\xc4\x95\x9f\x03\x34\x4e\x29\x81\x18\x87\xb1\x44\x61\xe8\xbb\x1b\xbb\x15\x66\xd5\x55\xaf\x8b\x97\x65\x91\x11\x17\x0b\x88\xd2\xe8\x50\xa9\x35\xd9\x70\xda\xef\x86\xe3\x17\xf4\x8e\x18\x33\x6c\x56\xeb\xf3\x3d\x88\xa0\xfe\x30\x71\x9b\xbc\x30\x3f\x9e\xb3\x6f\xa7\x32\xd3\x8a\xab\x04\x5e\x6d\xf8\xa6\x68\x2a\xb5\x89\x45\x0f\x20\x2d\xae\x8f\xec\xff\xd8\x3d\x79\x27\xdf\x4f\x00\x00\x00\xff\xff\xd0\x57\xb0\x99\x95\x00\x00\x00")

func templatesUrlViewHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesUrlViewHtml,
		"templates/url-view.html",
	)
}

func templatesUrlViewHtml() (*asset, error) {
	bytes, err := templatesUrlViewHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/url-view.html", size: 149, mode: os.FileMode(420), modTime: time.Unix(1455258875, 0)}
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
	"templates/404.html": templates404Html,
	"templates/base.html": templatesBaseHtml,
	"templates/config-base.html": templatesConfigBaseHtml,
	"templates/login.html": templatesLoginHtml,
	"templates/register.html": templatesRegisterHtml,
	"templates/settings.html": templatesSettingsHtml,
	"templates/shitbucket-import.html": templatesShitbucketImportHtml,
	"templates/tag-index.html": templatesTagIndexHtml,
	"templates/url-edit.html": templatesUrlEditHtml,
	"templates/url-index.html": templatesUrlIndexHtml,
	"templates/url-new.html": templatesUrlNewHtml,
	"templates/url-view.html": templatesUrlViewHtml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"base.html": &bintree{templatesBaseHtml, map[string]*bintree{}},
		"config-base.html": &bintree{templatesConfigBaseHtml, map[string]*bintree{}},
		"login.html": &bintree{templatesLoginHtml, map[string]*bintree{}},
		"register.html": &bintree{templatesRegisterHtml, map[string]*bintree{}},
		"settings.html": &bintree{templatesSettingsHtml, map[string]*bintree{}},
		"shitbucket-import.html": &bintree{templatesShitbucketImportHtml, map[string]*bintree{}},
		"tag-index.html": &bintree{templatesTagIndexHtml, map[string]*bintree{}},
		"url-edit.html": &bintree{templatesUrlEditHtml, map[string]*bintree{}},
		"url-index.html": &bintree{templatesUrlIndexHtml, map[string]*bintree{}},
		"url-new.html": &bintree{templatesUrlNewHtml, map[string]*bintree{}},
		"url-view.html": &bintree{templatesUrlViewHtml, map[string]*bintree{}},
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

