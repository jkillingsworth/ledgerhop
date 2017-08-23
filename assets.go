// Code generated by go-bindata.
// sources:
// assets/index.html
// assets/ripple/account-info.html
// assets/ripple/account-lines.html
// assets/ripple/ping.html
// assets/ripple/submit-trust-set.html
// assets/styles.css
// DO NOT EDIT!

package main

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

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x51\x3d\x4f\xf4\x30\x0c\xde\x4f\xea\x7f\xf0\x9b\xbd\x6f\x74\x37\x31\xb8\x91\x10\x1f\x02\xe9\x24\x4e\x5c\x19\x18\x73\xad\xdb\x44\xa4\x69\x95\xb8\xc3\xfd\x7b\xd4\xa4\x1c\xb0\x20\x98\x6c\x3f\x1f\x7e\x24\x1b\xff\xdd\x3e\xdd\xd4\xaf\x87\x3b\x30\x3c\x38\x55\x6c\x70\xad\x4b\x47\xba\x55\xc5\x06\x00\xd9\xb2\x23\xb5\xa7\xb6\xa7\xf0\x30\x4e\x28\x33\x90\xb8\x81\x58\x43\x63\x74\x88\xc4\x95\x78\xa9\xef\xcb\x2b\x91\x19\x67\xfd\x1b\x04\x72\x95\x88\x7c\x76\x14\x0d\x11\x0b\x30\x81\xba\x4a\xc8\x0c\xfd\x6f\x62\x5c\xd4\x28\xd7\xb0\x62\x83\xa7\xb1\x3d\xa7\x0e\x00\xcd\xf6\x6b\xaa\xd9\xe6\xc5\x26\x5c\xf8\x9d\x7a\xb6\xd3\xe4\x08\xa5\xd9\x65\x72\x76\xa9\xa6\x7c\x85\xfa\x23\x2f\x24\x99\x9c\xac\xef\x85\x3a\x58\xdf\xa3\xd4\x0a\xa5\xb3\x3f\xa9\x75\xd3\x8c\xb3\xe7\xd2\xfa\x6e\x14\xea\x3a\x4f\xf0\xe8\xbb\xf1\x2f\x6e\x67\x3d\xc5\x4f\xfb\x7e\x19\x7f\xe3\x8f\xf3\x69\xb0\x5c\x72\x98\x23\x97\x91\x58\xa8\x63\x42\xa0\x5e\x10\x38\x12\x7f\xdb\x82\x72\x5e\x1f\x27\x2f\x27\x44\x99\xde\xf9\x1e\x00\x00\xff\xff\x85\x59\x34\xf5\xe6\x01\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 486, mode: os.FileMode(438), modTime: time.Unix(1500851083, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rippleAccountInfoHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\x51\x6b\xdc\x30\x0c\x7e\x3f\xb8\xff\xa0\xe9\xbd\x67\xfa\x36\x86\x62\x18\xed\xc6\x0a\x83\x8d\xa3\x7b\xd8\xd3\xf0\xc5\xca\x62\xe6\xc4\xa9\xed\x1c\x3b\x4a\xff\xfb\x88\xed\xa4\x57\x76\x1b\x2d\xd7\x27\x4b\xb2\xbe\xcf\x9f\x14\x45\xf4\xe6\xfa\xcb\xd5\xed\xf7\xaf\x1f\xa0\x8d\x9d\x95\xeb\x15\x95\x73\xb2\x58\x69\xb9\x5e\x01\x50\x34\xd1\xb2\xfc\xcc\xfa\x27\xfb\x4f\x6e\x20\x91\x03\xe9\xae\xe3\xa8\xa0\x6e\x95\x0f\x1c\x2b\xfc\x76\xfb\xf1\xe2\x2d\xe6\x1b\x6b\xfa\x5f\xe0\xd9\x56\x18\xe2\xc1\x72\x68\x99\x23\x42\xeb\xb9\xa9\x70\xb3\x11\x39\xb8\xa9\x43\x98\xf2\x49\x94\xe7\xd6\x2b\xda\x39\x7d\x48\x16\x00\xb5\x97\x72\x6b\x86\xc1\x32\xbc\xaf\x6b\x37\xf6\x11\x6e\xfa\xc6\x91\x68\x2f\xf3\x23\xad\x9f\x33\x1b\xe7\x3b\x50\x75\x34\xae\xaf\x50\xf8\x04\x12\x2a\x83\x2e\x4c\xdf\x38\x84\x8e\x63\xeb\x74\x85\x83\x0b\x31\x8b\x9c\x8a\x53\xbb\x52\x4b\x76\xfd\x62\x4f\x9e\x86\xda\xaa\x10\x2a\xb4\x6a\xc7\x16\x25\xa5\x13\x1a\xe7\x2b\x34\x1a\xe5\xcd\x35\x89\x14\x92\x24\xa2\x3e\x0d\x6d\x0c\x5b\x8d\x92\x4c\x3f\x8c\x11\xe2\x61\xe0\x0a\x23\xff\x8e\x08\xbd\xea\x38\xf1\x80\xd1\x99\xef\x09\x0b\x89\x23\x35\x2f\x51\x56\xbb\xae\x53\xbd\x46\x79\x95\x8d\xf3\x35\xce\x8c\x49\xe8\xe2\xec\x95\x1d\xb9\xc2\xd2\xe5\x1f\xb9\xcb\x9e\x95\x76\xbd\x3d\xbc\x52\x2d\x36\x8d\x1d\x96\xf1\x7b\x76\x25\x8f\x77\x00\x14\xd8\x72\x1d\x4b\x29\x85\x30\x55\x32\x93\x1f\x67\x03\x90\x1b\xa6\x31\x9a\xcb\xdb\x2b\x6b\xb4\x8a\xac\x11\x32\x11\x6b\xb9\xc4\x48\xe4\xe4\xff\x53\xd4\xd6\x05\xd6\x28\xf3\xf9\x4c\xcc\xe8\x3d\xf7\x11\x65\x31\x4e\xa2\x48\x64\x49\xc7\xbd\x78\x95\xb6\x97\x6f\x8a\xb2\xfc\x77\xe7\x8f\xd0\xcc\x98\x1a\xbf\xd0\x9f\xa1\xf6\x45\x4a\xc2\xb8\xeb\xcc\xa2\xa5\x78\xff\x7c\x9d\xc4\xe3\x5a\x20\x31\xad\x96\x65\x21\x2d\x0b\x47\x9b\x7d\xc9\xde\xf2\xdd\xc8\x21\xbe\x2b\xd8\xc1\xb3\xbc\xbf\xdf\x6c\xf9\xee\xe1\x81\xc4\xe4\x65\x9a\x19\xf0\x04\x19\x06\xd7\x07\xfe\x0b\x1a\x4e\x41\xa7\x35\xb9\x2c\x47\x12\x69\x55\xff\x09\x00\x00\xff\xff\x6e\x53\x1f\x52\xc2\x05\x00\x00")

func rippleAccountInfoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_rippleAccountInfoHtml,
		"ripple/account-info.html",
	)
}

func rippleAccountInfoHtml() (*asset, error) {
	bytes, err := rippleAccountInfoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ripple/account-info.html", size: 1474, mode: os.FileMode(438), modTime: time.Unix(1500847606, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rippleAccountLinesHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\x51\x6b\xdc\x30\x0c\x7e\x2f\xf4\x3f\x68\x7a\xef\x99\xbe\x8d\xa1\x18\x46\xbb\xb1\x41\x61\xe3\xe8\x1e\xf6\x34\x7c\xb1\x6e\x31\x73\xec\xd4\x76\x8e\x1d\xa5\xff\x7d\xc4\x76\xd2\x2b\xbb\x8d\x96\xeb\x93\x25\x59\xdf\xe7\x4f\x8a\x22\x7a\x73\xfd\xe5\xea\xf6\xfb\xd7\x0f\xd0\xa5\xde\xca\xf3\x33\xaa\xe7\x64\xb1\xd2\xf2\xfc\x0c\x80\x92\x49\x96\xe5\x0d\xeb\x9f\x1c\x3e\xf9\x81\x44\x09\xe4\xbb\x9e\x93\x82\xb6\x53\x21\x72\x6a\xf0\xdb\xed\xc7\x8b\xb7\x58\x6e\xac\x71\xbf\x20\xb0\x6d\x30\xa6\xbd\xe5\xd8\x31\x27\x84\x2e\xf0\xb6\xc1\xd5\x4a\x94\xe0\xaa\x8d\x71\xca\x27\x51\x9f\x3b\x3f\xa3\x8d\xd7\xfb\x6c\x01\x50\x77\x29\xd7\x66\x18\x2c\xc3\xfb\xb6\xf5\xa3\x4b\x70\x63\x1c\x47\x12\xdd\x65\x79\xa5\x0b\x73\xea\xd6\x87\x1e\x54\x9b\x8c\x77\x0d\x8a\x90\x51\x42\x15\xd4\x85\x9d\x50\x08\x3d\xa7\xce\xeb\x06\x07\x1f\x53\x91\x39\x95\xa7\x36\xb5\x9a\xe2\x86\xc5\x9e\x3c\x0d\xad\x55\x31\x36\x68\xd5\x86\x2d\x4a\xca\x27\x6c\x7d\x68\xd0\x68\x94\x9f\xaf\x49\xe4\x90\x24\x91\xf4\x71\xe8\xd6\xb0\xd5\x28\xc9\xb8\x61\x4c\x90\xf6\x03\x37\x98\xf8\x77\x42\x70\xaa\xe7\xcc\x03\x46\x17\xbe\x27\x2c\x24\x0e\xd4\xbc\x44\x59\xeb\xfb\x5e\x39\x8d\xf2\xaa\x18\xa7\x6b\x9c\x19\xb3\xd0\xc5\xd9\x29\x3b\x72\x83\xb5\xcd\x3f\x6a\x9b\x03\x2b\xed\x9d\xdd\xbf\x52\x31\x36\x4f\x1e\xd6\x09\x7c\x76\x29\x8f\x77\x00\x14\xd9\x72\x9b\x6a\x2d\x95\x30\x97\x32\x93\x1f\x66\x03\x90\x1f\xa6\x41\x9a\xeb\xdb\x29\x6b\xb4\x4a\xac\x11\x0a\x11\x6b\xb9\xc4\x48\x94\xe4\xff\x53\xb4\xd6\x47\xd6\x28\xcb\xf9\x4c\xcc\x18\x02\xbb\x84\xb2\x1a\x47\x51\x24\x8a\xa4\xc3\x5e\xbc\x4a\xdb\xeb\x47\x45\x59\x7f\xbd\xd3\x67\x68\x66\xcc\x8d\x5f\xe8\x4f\x50\xfb\x22\x25\x71\xdc\xf4\x66\xd1\x52\xbd\x7f\xbe\x4e\xe2\x71\x2f\x90\x98\x96\xcb\xb2\x93\x96\x95\xa3\xcd\xae\x66\xaf\xf9\x6e\xe4\x98\xde\x55\xec\x10\x58\xde\xdf\xaf\xd6\x7c\xf7\xf0\x40\x62\xf2\x0a\xcd\x0c\x78\x82\x8c\x83\x77\x91\xff\x82\xc6\x63\xd0\x69\x53\x2e\xfb\x91\x44\xde\xd6\x7f\x02\x00\x00\xff\xff\x07\x73\x24\x35\xc5\x05\x00\x00")

func rippleAccountLinesHtmlBytes() ([]byte, error) {
	return bindataRead(
		_rippleAccountLinesHtml,
		"ripple/account-lines.html",
	)
}

func rippleAccountLinesHtml() (*asset, error) {
	bytes, err := rippleAccountLinesHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ripple/account-lines.html", size: 1477, mode: os.FileMode(438), modTime: time.Unix(1500830165, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ripplePingHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x53\xc1\x6a\x1b\x31\x10\xbd\x1b\xfc\x0f\xd3\xb9\xd7\x22\xb7\x52\x46\xba\x24\x2d\x2d\x14\x1a\x4c\x7a\xe8\x51\x5e\x8d\xbd\xa2\x5a\x49\x91\xc6\xa1\x26\xe4\xdf\xcb\xae\xd6\xdb\x86\x26\x87\xe0\x93\xe6\x8d\xe6\x3d\x3d\xed\x3e\xd1\xbb\x9b\xef\xd7\x77\x3f\x6f\x3f\x41\x2f\x43\x30\xeb\x15\xcd\xeb\x58\xb1\x75\x66\xbd\x02\x20\xf1\x12\xd8\x7c\x63\x77\xe0\xf2\x25\x65\x52\xad\x31\xed\x0d\x2c\x16\xba\xde\x96\xca\xa2\xf1\xc7\xdd\xe7\xf7\x1f\xb0\xed\x04\x1f\x7f\x41\xe1\xa0\xb1\xca\x29\x70\xed\x99\x05\xa1\x2f\xbc\xd7\xb8\xd9\xa8\xd6\xdc\x74\xb5\x8e\xf3\xa4\xe6\xe3\xd6\x2b\xda\x25\x77\x9a\x2a\x00\xea\xaf\xcc\xd6\xe7\x1c\x18\x6e\x7d\x3c\x90\xea\xaf\x9a\x78\x5f\xce\x13\xfb\x54\x06\xb0\x9d\xf8\x14\x35\xaa\x32\x0d\xab\xec\xe3\x01\x61\x60\xe9\x93\xd3\x98\x53\x95\x66\x6a\xbc\x8c\xdd\xcd\xde\x1b\x2c\x4b\x3d\x22\x07\x5d\xb0\xb5\x6a\x0c\x76\xc7\x01\x0d\x4d\x2b\xec\x53\xd1\xe8\x1d\x9a\xaf\x37\xa4\xa6\x96\x21\x25\xee\x65\xea\xde\x73\x70\x68\xc8\xc7\x7c\x14\x90\x53\x66\x8d\xc2\xbf\x05\x21\xda\x81\x27\x1d\xf0\xae\xe9\x3d\x53\x21\xf5\x8f\x9b\xb7\x38\xeb\xd2\x30\xd8\xe8\xd0\x5c\xb7\xe2\x72\x8f\x67\xc5\xc9\xe8\x02\x1e\x6c\x38\xb2\xc6\xf6\x75\x0b\x5b\x97\x62\x38\x5d\x72\x87\x37\xf9\xab\xc7\xdd\xe0\x17\x87\x33\x7a\xf5\x74\x52\x7f\xff\x34\xa9\x31\x25\x4b\xa6\x96\xec\x38\xff\x30\x4f\x6f\xf9\xfe\xc8\x55\x3e\xce\xdc\x5c\xd8\x3c\x3e\x6e\xb6\x7c\xff\xf4\x44\x6a\x44\x4d\xe6\x4c\x78\xc6\xac\x39\xc5\xca\xff\x51\xeb\x4b\xd4\x31\xe9\x4b\xbe\x49\x4d\xaf\xed\x4f\x00\x00\x00\xff\xff\xff\x30\xc4\x9c\x85\x03\x00\x00")

func ripplePingHtmlBytes() ([]byte, error) {
	return bindataRead(
		_ripplePingHtml,
		"ripple/ping.html",
	)
}

func ripplePingHtml() (*asset, error) {
	bytes, err := ripplePingHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ripple/ping.html", size: 901, mode: os.FileMode(438), modTime: time.Unix(1500828548, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rippleSubmitTrustSetHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x58\x4f\x4f\xe3\x38\x14\xbf\x57\xea\x77\xf0\xfa\x4e\xd3\xa2\x3d\xa0\x95\x63\x09\x95\xad\x16\x2d\xda\x1d\xd1\xce\x48\xa3\x8a\x83\x9b\xbc\x36\x11\x4e\x1c\x6c\x87\x21\x83\xf8\xee\xa3\xd8\xf9\xd7\x34\x03\x03\x75\x87\x4b\xec\x67\xbf\xf7\xfb\xc3\x53\xe2\x9a\xfc\x71\xf5\xff\x7c\xf5\xf5\xd3\xdf\x28\xd2\x09\xa7\xe3\x11\xa9\x9e\xe5\x08\x58\x48\xc7\x23\x84\x88\x8e\x35\x07\x7a\x03\xe1\x0e\xe4\x3f\x22\x23\x9e\x0d\x98\xb5\x04\x34\x43\x41\xc4\xa4\x02\xed\xe3\xcf\xab\xc5\xd9\x05\xb6\x2b\x3c\x4e\xef\x91\x04\xee\x63\xa5\x0b\x0e\x2a\x02\xd0\x18\x45\x12\xb6\x3e\x9e\x4c\x3c\x1b\x9c\x04\x4a\x95\xfb\x89\x57\xc1\x8d\x47\x64\x23\xc2\xc2\x8c\x10\x22\xd1\x8c\xde\xc6\x59\xc6\x01\x2d\xf3\x4d\x12\x6b\xb4\x92\xb9\xd2\x68\x09\x9a\x78\xd1\xcc\x02\x45\xb2\xde\xbd\x15\x32\x41\x2c\xd0\xb1\x48\x7d\xec\x49\x93\xe8\x29\x93\x78\xa6\xcb\xc4\x33\x55\x72\x48\x40\x47\x22\xf4\x71\x26\x94\xb6\x64\x4b\x91\x6c\x53\x69\xb2\x53\xd9\x8c\xcb\x59\x88\x02\xce\x94\xf2\x31\x67\x1b\xe0\x98\x12\xf3\x44\x5b\x21\x7d\x1c\x87\x98\x5e\x5f\x11\xcf\x84\x28\xf1\x74\x38\x9c\xba\x8d\x81\x87\x98\x92\x38\xcd\x72\x8d\x74\x91\x81\x8f\x35\x3c\x69\x8c\x52\x96\x80\xa9\x83\xe2\xd0\xd6\xdb\xab\x42\xbc\x0e\x9b\xf7\x30\x0b\x44\x92\xb0\x34\xc4\x74\x6e\x07\xc7\x73\xac\x2b\x1a\xa2\xcd\xe4\x91\xf1\x1c\x7c\x6c\x9d\xc6\x48\x02\x0b\x45\xca\x0b\x47\x2a\xb4\x64\xa9\xb2\xff\xd5\x55\x91\x01\xa6\xab\x36\x60\x28\x1e\x2f\xab\x0f\x61\xe4\x1d\x04\x2b\x99\xe7\x53\xe7\x12\x59\x10\x88\x3c\xd5\x98\x5e\xda\xc1\xf1\x8a\xea\x8a\x46\x49\x53\xde\x0d\x5b\x05\x81\x04\xfd\x2f\x14\x98\x2e\xcd\x10\xdd\x43\xf1\x31\xca\x19\x53\xea\x9b\x90\x61\x4d\xbb\x2d\x6d\x88\x77\x90\xdc\x50\xdf\x02\x60\xba\x80\x0f\xb6\x4c\x9a\x27\x1b\x90\x35\xd7\xb2\x96\x61\x69\x8a\xba\xb2\xf6\x21\x87\x34\x80\xd2\x59\x3b\x72\xc1\xb4\xa9\x5a\x99\x5a\x63\xb8\xe1\xcc\x99\xd2\xf6\xe3\xb0\x6c\x2a\xdf\x30\xa5\x11\x37\x41\xa4\x1c\x2a\x19\xc0\x32\x9a\x86\x38\x38\xea\x18\xce\x76\x8b\x9c\xf3\x62\xce\x52\x91\xc6\x01\xe3\xcb\x78\x87\xe9\x82\xb3\xdd\xc7\x04\x05\x11\x04\xf7\x1b\xf1\xd4\xbc\x4f\xa6\x4f\x17\x53\xfb\xd7\x74\xd6\x20\xa6\xed\xb5\xe1\x25\x53\x14\xba\x2c\xca\x2f\xf0\x9b\x2a\xd6\x2d\xf6\x1d\x3a\x58\xaf\xf5\x75\xa4\x39\xb3\x74\x09\xfa\x32\xd7\x91\x63\x23\xa7\xd3\xe9\xac\x6f\x64\x8d\xd4\xd8\xd7\x40\xbf\x66\x56\xb3\x69\xdd\x56\xbd\x43\x55\xf4\xc4\xc6\xfc\x27\xec\x99\xc7\xbd\x39\xe7\x03\xe6\x34\x68\x5d\x83\x5a\x0a\x6f\x98\xd4\x6e\x5c\xb7\x08\xc6\xa8\x7a\xe5\x94\x66\xcd\x39\x30\x79\x3a\xbb\xfe\xec\xdb\xb5\x8f\xd7\x18\xd6\xa3\xf1\x9a\x65\xbd\xad\xeb\x16\xe7\xb7\x99\xb6\x04\xbd\x90\x00\xdf\xdd\x1b\x36\x3b\x78\x8b\xb5\x58\xdd\xee\xaa\xe1\xdf\xe8\xad\x7a\xdb\xba\xad\x6d\x4c\xb2\xf1\x93\xf7\xd5\x89\x4c\x3a\x3f\x30\xa9\x8b\xb6\xdf\x53\xbf\x62\xd4\xde\xc6\x75\x8b\x70\x87\x3a\x2b\xa7\x32\x8b\xc7\x49\xac\x2f\x93\xf2\x74\x39\xcf\xa5\x84\x34\x28\x30\xbd\x29\x83\x88\x99\x28\x0a\xaa\xf0\xf1\x47\xda\x21\x2c\xfb\xf1\x1f\x22\xe1\x5c\xdf\xb5\x52\x39\xc8\x9e\xba\xd8\x04\x9d\x6a\xab\x70\xfa\xca\x6a\x78\xe7\xba\xbe\x94\xad\xd9\x93\x65\xda\xd5\xa9\x2a\x8b\xd2\x17\x55\x61\x1f\xa1\xe9\x5d\xdc\xea\x1f\xa8\xd5\xa9\xd8\xce\x7e\x8a\x4e\xbc\xf6\x4e\x80\x78\x5b\x21\x93\xe6\x56\xa2\xb9\x71\x08\xe3\xc7\x6a\xf7\x6d\x79\xee\x54\xfa\xaf\x2a\x37\x93\x40\x9f\x9f\x27\xb7\xf0\xf0\xf2\x42\xbc\x72\x66\xcb\xd4\x09\x7b\x99\x2a\x13\xa9\x82\x83\x54\x35\x94\x3a\x1e\x11\xaf\xb9\x21\x21\x9e\xb9\xaf\xf9\x11\x00\x00\xff\xff\xec\x1b\xfd\x5c\xc7\x11\x00\x00")

func rippleSubmitTrustSetHtmlBytes() ([]byte, error) {
	return bindataRead(
		_rippleSubmitTrustSetHtml,
		"ripple/submit-trust-set.html",
	)
}

func rippleSubmitTrustSetHtml() (*asset, error) {
	bytes, err := rippleSubmitTrustSetHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ripple/submit-trust-set.html", size: 4551, mode: os.FileMode(438), modTime: time.Unix(1502136335, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _stylesCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x91\xcf\x6a\xc3\x30\x0c\x87\xef\x06\xbf\x83\x60\xec\x32\x48\xf1\x18\x83\xe1\xb2\x27\x19\x39\xf8\x8f\xd2\x9a\x2a\x91\x71\x14\x96\x6e\xec\xdd\x47\xdb\x25\x64\xa1\x2d\xba\x18\xfc\xfd\x3e\x09\xa9\xe1\xd2\x82\x38\x4f\xa8\xd5\xb7\x56\x00\x00\x9e\x4b\xc4\x52\x05\x26\x72\xb9\x47\x0b\xd3\x6b\xab\xd5\x8f\x56\x5a\x5d\x32\x11\x9e\x16\x91\xb1\xea\xd3\x57\xea\x76\x76\x8a\x7b\x1e\xb7\x97\xcf\xcf\x14\x65\x6f\xe1\xd9\x98\xc7\x95\x61\x43\xce\x23\xcd\x96\x19\xc4\xf6\x2a\x08\xff\xf1\x86\xd8\x89\x05\xc2\x46\xd6\x78\x93\x90\xe2\xda\xfb\xb2\xf6\xa6\x2e\x0f\xf2\x51\xd0\x45\xee\xe8\x58\xcf\x7c\x60\xe2\x62\xe1\xe1\xcd\x9c\xea\x4a\x44\x8e\x19\xdf\xc3\x1e\xc3\xc1\xf3\x58\xdf\x18\x68\xd1\xd9\x0d\xc2\xb7\x34\xfd\xe0\xdb\x24\xf5\xbd\x25\xe4\xb2\x38\x8e\x0b\x87\x5d\xe1\xa1\x8b\xd5\x34\x26\x9a\x53\xfd\x75\xcc\x2e\xc6\xf3\x19\xcc\xe6\x15\xdb\xb3\xe0\x37\x00\x00\xff\xff\xf3\xad\x7d\xf9\xe3\x01\x00\x00")

func stylesCssBytes() ([]byte, error) {
	return bindataRead(
		_stylesCss,
		"styles.css",
	)
}

func stylesCss() (*asset, error) {
	bytes, err := stylesCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "styles.css", size: 483, mode: os.FileMode(438), modTime: time.Unix(1502209241, 0)}
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
	"index.html": indexHtml,
	"ripple/account-info.html": rippleAccountInfoHtml,
	"ripple/account-lines.html": rippleAccountLinesHtml,
	"ripple/ping.html": ripplePingHtml,
	"ripple/submit-trust-set.html": rippleSubmitTrustSetHtml,
	"styles.css": stylesCss,
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
	"index.html": &bintree{indexHtml, map[string]*bintree{}},
	"ripple": &bintree{nil, map[string]*bintree{
		"account-info.html": &bintree{rippleAccountInfoHtml, map[string]*bintree{}},
		"account-lines.html": &bintree{rippleAccountLinesHtml, map[string]*bintree{}},
		"ping.html": &bintree{ripplePingHtml, map[string]*bintree{}},
		"submit-trust-set.html": &bintree{rippleSubmitTrustSetHtml, map[string]*bintree{}},
	}},
	"styles.css": &bintree{stylesCss, map[string]*bintree{}},
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

