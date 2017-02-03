// Code generated by go-bindata.
// sources:
// ../../migrations/0001_initial.sql
// ../../migrations/0002_join_accept_params.sql
// ../../migrations/0003_rx_window_and_rx2_dr.sql
// ../../migrations/0004_add_node_apps_nwks_key_name_devaddr.sql
// ../../migrations/0005_add_queue.sql
// ../../migrations/0006_remove_application_table.sql
// ../../migrations/0007_migrate_channels_to_channel_list.sql
// ../../migrations/0008_relax_fcnt.sql
// ../../migrations/0009_adr_interval_and_install_margin.sql
// ../../migrations/0010_recreate_application_table.sql
// DO NOT EDIT!

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

var __0001_initialSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x90\xc1\x4e\xc3\x30\x0c\x86\xcf\xcb\x53\xf8\xb8\x09\x26\x8d\x73\xaf\xbc\x02\xe7\xca\x4b\x7e\x46\xb4\xd4\x8e\xdc\xb4\x90\xb7\x47\x2d\x05\x1a\x09\x6e\x91\xed\x7c\xff\x67\x9f\xcf\xf4\x30\xc4\x9b\x71\x01\xbd\x64\xe7\x0d\xcb\xab\xf0\x35\x81\x38\xe7\x14\x3d\x97\xa8\x42\x47\x77\xe0\x9c\x7b\x4c\x91\xae\xb5\x80\x29\x5b\x1c\xd8\x2a\xdd\x51\x1f\xdd\x41\x78\x00\xf9\x37\x36\xf6\x05\x46\x33\x5b\x8d\x72\xa3\xe3\xd3\xe5\x72\x22\xd1\x42\x32\xa5\xe4\x4e\x9d\x6b\x13\x44\x03\x16\x74\xc0\xfc\x2f\xba\x8d\x35\xbc\xc2\x20\x1e\x63\xa3\xa7\x42\x01\x09\x05\xe4\x79\xf4\x1c\xf0\x13\xba\x11\xee\xa8\x1b\x61\xd7\x98\x46\x84\x7e\xc9\x16\x5d\x89\xeb\xc0\xde\x32\x4a\xc0\xc7\x6a\xd9\x7f\x6b\xa8\x6c\xd6\x5b\x61\x99\xde\x1f\xf1\x59\xdf\xc5\x05\xd3\xfc\xc7\xe7\xce\x7d\x75\x7e\x97\x6f\x2b\xbb\x8d\x3a\xf7\x19\x00\x00\xff\xff\x82\x1d\xc3\x75\x9a\x01\x00\x00")

func _0001_initialSqlBytes() ([]byte, error) {
	return bindataRead(
		__0001_initialSql,
		"0001_initial.sql",
	)
}

func _0001_initialSql() (*asset, error) {
	bytes, err := _0001_initialSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0001_initial.sql", size: 410, mode: os.FileMode(420), modTime: time.Unix(1472984060, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __0002_join_accept_paramsSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x52\x4d\x6e\xb3\x30\x10\x5d\xe3\x53\xcc\x32\xe8\x23\x12\xf9\xba\xe8\x82\x26\xab\x5e\xa1\x6b\x34\xb1\x07\x62\xc5\x8c\x93\xc1\xb4\xe5\xf6\x15\x8e\x20\x84\x22\xb5\xea\xce\x9a\x79\x6f\xde\x8f\xbc\xdd\xc2\xbf\xc6\xd6\x82\x81\xe0\xed\xa2\xb4\xd0\xf0\x0a\x78\x74\x04\xfa\x84\xcc\xe4\x4a\x67\xdb\x00\x1b\x95\x58\x03\x47\x5b\xb7\x24\x16\x1d\x5c\xc4\x36\x28\x3d\x9c\xa9\xcf\x54\xc2\xd8\x44\xbc\xa0\x0e\x24\xf0\x8e\xd2\x5b\xae\x61\xb3\xcb\xf3\x14\xd8\x07\xe0\xce\x39\x95\x16\x6a\x55\xe1\x87\xe3\x73\x1f\xe5\x0d\x67\x39\x80\x50\x45\x42\xac\xa9\x7d\x74\xea\x19\x0c\x39\x0a\x04\x1a\x5b\x8d\x86\x26\x03\xf7\x5b\x60\x39\x50\x4d\x32\x5f\x55\x42\xd7\x8e\x58\xf7\x6b\x4b\x7d\x22\x7d\x86\xcd\x48\x3f\xec\xe1\x09\x90\xcd\x94\xe0\x65\x0f\xcf\x71\x70\xbf\x72\x80\x3c\xcd\x54\xd2\xb1\xbd\x76\x34\x51\xc7\x14\xd9\x48\x4d\x63\x2d\xe8\x86\xde\x6e\xad\xb0\x37\xa4\x12\x34\x06\xb4\x77\x5d\xc3\x20\x9f\xa5\x21\x87\xd1\xd8\xff\xc9\x15\x18\xaa\xb0\x73\x01\xf2\x6c\x81\xde\x95\x46\x4a\x5f\x55\x2d\x85\xdf\x51\xfe\xdc\xf0\xa0\x30\x1c\x5e\x4f\x10\x27\xcb\x0c\x46\xfc\x65\xf4\x91\x7d\x03\xcd\xad\xcf\x91\x85\x52\xf3\xaf\xfa\xea\x3f\x78\x45\x31\x32\x16\x82\xd9\x72\x3c\x93\x58\xec\x16\x35\x14\x4a\xc5\xed\xc3\x5f\x5d\x1d\x46\x46\xa1\xbe\x02\x00\x00\xff\xff\x38\x36\x6f\x1a\x4d\x03\x00\x00")

func _0002_join_accept_paramsSqlBytes() ([]byte, error) {
	return bindataRead(
		__0002_join_accept_paramsSql,
		"0002_join_accept_params.sql",
	)
}

func _0002_join_accept_paramsSql() (*asset, error) {
	bytes, err := _0002_join_accept_paramsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0002_join_accept_params.sql", size: 845, mode: os.FileMode(420), modTime: time.Unix(1472974457, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __0003_rx_window_and_rx2_drSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\xcd\xb1\x0a\xc2\x40\x0c\x87\xf1\xd9\x7b\x8a\xff\x6e\x0b\xd2\xb5\xab\xaf\xe0\x5c\xa2\x89\x72\x90\x26\x25\xe4\x38\x1f\x5f\xdc\xa4\xea\xfa\xf1\xc1\x6f\x1c\x71\x5c\xeb\x23\x28\x05\x97\xad\x90\xa6\x04\x92\xae\x2a\x30\x67\x29\x07\x62\xc6\xcd\xb5\xad\x86\x78\x2e\xbd\x1a\x7b\x47\xb5\x9c\x60\x9e\xb0\xa6\x0a\x96\x3b\x35\x4d\x9c\x86\xdd\x3e\x2d\x1c\xff\xde\xb9\x94\x4f\xfb\xec\xdd\x7e\xe8\x1c\xbe\x7d\xf1\xc3\xbe\xbf\x9d\xb9\xbc\x02\x00\x00\xff\xff\xd0\x11\x0c\xf5\xcb\x00\x00\x00")

func _0003_rx_window_and_rx2_drSqlBytes() ([]byte, error) {
	return bindataRead(
		__0003_rx_window_and_rx2_drSql,
		"0003_rx_window_and_rx2_dr.sql",
	)
}

func _0003_rx_window_and_rx2_drSql() (*asset, error) {
	bytes, err := _0003_rx_window_and_rx2_drSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0003_rx_window_and_rx2_dr.sql", size: 203, mode: os.FileMode(420), modTime: time.Unix(1472974457, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __0004_add_node_apps_nwks_key_name_devaddrSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x91\xbd\x4e\xc3\x40\x10\x84\x6b\xee\x29\xb6\x73\x22\x12\x64\x6a\x2b\x1d\x3c\x02\x5d\x24\x6b\xe3\x9d\x90\x53\xee\x4f\xe7\xbd\x98\xbc\x3d\xb2\xf9\x11\x71\x90\xa0\xc8\x75\x33\xba\x19\xad\xbe\x59\xaf\xe9\xde\xdb\xd7\xcc\x0a\x7a\x49\x86\x9d\x22\x93\xf2\xce\x81\x42\x14\x98\x3b\x16\xa1\x2e\xba\xe2\x03\x71\x4a\x6d\xdf\x1e\x71\xa6\xdd\x59\xc1\x14\xa2\x52\x28\xce\x91\x60\xcf\xc5\x29\x3d\x57\xdb\xed\x5b\xfd\xc7\xab\x56\x17\xa5\x61\x38\xde\xbe\x54\x70\x6a\x59\x24\xff\xab\x73\x7e\x10\x7b\xd0\x89\x73\x77\xe0\xbc\x78\xac\xeb\xe5\x75\xbc\xaa\x1a\x63\x4a\x92\x11\xda\x48\x89\x7a\xe8\x94\xdb\x2c\x7a\x38\x74\x1f\x82\xf6\x39\xfa\x11\x9a\xb3\x1d\xab\x8d\x81\x86\x03\x32\x26\x8c\x28\x96\x36\x53\xf6\xe1\x53\x2e\x1b\x63\x7e\x8e\xf1\x14\x87\xf0\xcb\x1c\x92\x63\xba\xda\x63\x75\xe9\x7f\x23\x9d\xf9\x5f\x54\xe6\xdf\xd9\xa3\x31\xef\x01\x00\x00\xff\xff\x0d\x4a\x1c\x07\x09\x02\x00\x00")

func _0004_add_node_apps_nwks_key_name_devaddrSqlBytes() ([]byte, error) {
	return bindataRead(
		__0004_add_node_apps_nwks_key_name_devaddrSql,
		"0004_add_node_apps_nwks_key_name_devaddr.sql",
	)
}

func _0004_add_node_apps_nwks_key_name_devaddrSql() (*asset, error) {
	bytes, err := _0004_add_node_apps_nwks_key_name_devaddrSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0004_add_node_apps_nwks_key_name_devaddr.sql", size: 521, mode: os.FileMode(420), modTime: time.Unix(1474796205, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __0005_add_queueSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x91\xc1\x4e\xc3\x30\x10\x44\xcf\xf8\x2b\xe6\x98\x8a\x56\x2a\xe7\x5c\xf9\x05\xce\xd1\xc6\x9e\x04\x8b\xcd\x3a\x38\x4e\x0b\x7f\x8f\x42\x9b\x0a\x22\x84\x7a\xb3\xf4\x66\x56\x33\xe3\xc3\x01\x8f\x43\xec\xb3\x14\xe2\x65\x74\x3e\x73\x79\x15\x69\x95\x08\xe9\x6c\x1a\xed\xad\x79\x9f\x39\x13\x95\x03\x80\x18\xd0\xc6\x7e\x62\x8e\xa2\x7b\xf7\x90\xd9\x31\xd3\x3c\x71\x92\xec\x5f\x25\x57\x4f\xc7\xe3\x0e\x96\x0a\x6c\x56\xdd\x7f\x7b\x02\x4f\x0d\xe7\x88\xf6\xb3\x50\x70\xb3\x4c\xb0\x14\x88\x64\x08\x54\x16\xc2\xcb\xe4\x25\x70\xe3\xf6\xc9\xba\x98\x07\x06\xb4\x29\x29\xc5\x6e\x1c\x81\x9d\xcc\x5a\xd0\x89\x4e\xbc\xa8\x47\x5a\x88\xd6\xdf\xa5\xed\xc6\x94\x0b\xa6\x41\x54\xa3\x95\x6d\x68\x29\x72\x4d\xbc\x02\xb7\xab\xdd\x3a\x51\xb4\xc0\x8f\xcd\x44\xcd\x5a\x74\xa9\xf4\x8b\x54\x57\xb2\x1c\xf8\x39\xf9\x73\x3a\x9b\x0b\x39\x8d\xff\xde\xab\xdd\x45\xf3\xd7\xb7\xd4\xee\x2b\x00\x00\xff\xff\xb7\xd4\x53\x7f\xc3\x01\x00\x00")

func _0005_add_queueSqlBytes() ([]byte, error) {
	return bindataRead(
		__0005_add_queueSql,
		"0005_add_queue.sql",
	)
}

func _0005_add_queueSql() (*asset, error) {
	bytes, err := _0005_add_queueSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0005_add_queue.sql", size: 451, mode: os.FileMode(420), modTime: time.Unix(1474128106, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __0006_remove_application_tableSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x90\xc1\x6a\xf3\x30\x10\x84\xef\x7a\x8a\xb9\x25\xe1\x4f\x20\xff\xd9\xd7\xbe\x42\xcf\x66\x23\xad\xdd\x25\xf2\x4a\xac\x36\x2d\x7e\xfb\x62\x25\x94\x94\x96\xde\x04\x33\x9a\x6f\x67\x4e\x27\xfc\x5b\x64\x36\x72\xc6\x6b\x0d\x94\x9d\x0d\x4e\x97\xcc\xd0\x92\x38\x00\x40\xb2\x52\x11\x8b\x36\x37\x12\xf5\x2e\x8c\x54\xeb\xc8\x37\x19\xa7\x2b\xaf\x43\x08\xdd\x73\xff\x47\xb5\x66\x89\xe4\x52\x74\x08\xe1\x19\xf0\x52\x3e\x34\x44\xe3\xed\xfd\xc3\x8b\x7d\x87\x3d\x72\x71\x59\x9d\x09\xd5\x64\x21\x5b\x71\xe5\xf5\xd8\x65\xa5\x85\x11\xdf\xc8\x28\x6e\x97\xbe\x93\xad\xa2\x33\xf6\xff\xcf\xe7\x03\xb4\x38\xf4\x96\x73\x38\x0c\x21\x88\x36\x36\x87\xa8\x97\x67\x4c\x4f\x69\x9c\x39\x3a\x92\x34\x17\x8d\xbe\x7f\x40\x0f\x47\xec\xb4\x74\xc6\x0e\xd4\xee\xb0\xc9\xca\xd2\x2b\x0f\xe1\xf7\x79\x28\xa5\x3f\xd7\xc1\x54\x8c\x65\xd6\xad\xc4\x17\x09\xc6\x13\x1b\x6b\xe4\xf6\x6d\x84\xa2\x48\x9c\xd9\x19\x91\x5a\xa4\xc4\xc3\x67\x00\x00\x00\xff\xff\xb4\xbf\x2b\xef\xa2\x01\x00\x00")

func _0006_remove_application_tableSqlBytes() ([]byte, error) {
	return bindataRead(
		__0006_remove_application_tableSql,
		"0006_remove_application_table.sql",
	)
}

func _0006_remove_application_tableSql() (*asset, error) {
	bytes, err := _0006_remove_application_tableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0006_remove_application_table.sql", size: 418, mode: os.FileMode(420), modTime: time.Unix(1474703337, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __0007_migrate_channels_to_channel_listSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x91\xc1\x6e\xc2\x30\x10\x44\xcf\xd9\xaf\xd8\x23\x51\xa1\xaa\xd4\x43\x0f\x69\x38\xf5\x17\x7a\xaa\x2a\x64\xec\x01\x2c\x9c\x4d\x58\x3b\x42\xf9\xfb\xca\xa8\x09\x09\xaa\x7a\x4a\x3c\x3b\x6f\xad\x19\x6f\x36\xfc\xd4\xf8\xa3\x9a\x04\xfe\xec\xc8\x84\x04\xe5\x64\xf6\x01\x6c\x4f\x46\x04\x61\x17\x7c\x4c\x54\x18\xe7\xd8\xb6\xa1\x6f\x64\x1c\x44\xf6\x92\x70\x84\x7e\x7d\x57\x44\x7d\xe7\xf2\x8e\x25\x14\x91\x26\x77\xbd\xa2\xa2\x88\x08\xb0\x89\x8d\xaa\x19\xf2\x79\x14\x0e\x8a\x4b\x0f\xb1\x43\xd6\x0e\xda\x36\xf9\x5b\xfc\xa2\xf9\xff\x7a\x82\x62\x2e\xde\x6e\xd8\x79\x57\xcf\xcf\xcf\xde\x65\x4f\xab\x0e\xca\xfb\x61\xee\x2f\xcb\x8a\xc8\x69\xdb\x2d\xd3\x55\x44\xf3\x0a\x3e\xda\xab\xfc\x53\xc2\x8d\x7f\x68\xa1\x22\xb2\x8a\x0c\x2f\x08\x5e\x51\xe1\x1d\xef\xfd\x31\x42\xbd\x09\xdc\xa9\x6f\x8c\x0e\x7c\xc6\xb0\xa6\xc7\x14\xd9\xe7\x25\xb1\xe2\x00\x85\x58\xc4\xc5\xcd\xdc\x0a\x3b\x04\xe4\x82\x4d\xb4\xc6\x81\xa5\x4d\x2c\x7d\x08\xf7\x5d\xe3\x73\xcc\x47\x53\xaf\x7f\x0d\xed\x09\xf6\xcc\xab\x11\xdf\xd6\xfc\xca\x46\xdc\x94\xe0\xbd\xe6\xb7\x9b\x70\xdf\xb2\xe5\x97\x72\x4d\x45\x2f\xfe\xd2\x63\x42\xc7\x14\xeb\x11\x2d\xa9\xac\xe8\x27\x00\x00\xff\xff\xb3\x84\xa3\xcc\x5b\x02\x00\x00")

func _0007_migrate_channels_to_channel_listSqlBytes() ([]byte, error) {
	return bindataRead(
		__0007_migrate_channels_to_channel_listSql,
		"0007_migrate_channels_to_channel_list.sql",
	)
}

func _0007_migrate_channels_to_channel_listSql() (*asset, error) {
	bytes, err := _0007_migrate_channels_to_channel_listSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0007_migrate_channels_to_channel_list.sql", size: 603, mode: os.FileMode(420), modTime: time.Unix(1474740058, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __0008_relax_fcntSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\xcc\xb1\x0d\xc2\x40\x0c\x05\xd0\x1a\x4f\xf1\x7b\x94\x09\xd2\xb2\x02\x35\x72\x62\x1f\x42\xfa\xb1\x4f\x87\x4f\x30\x3e\x2d\x12\x2c\xf0\x96\x05\xe7\xe3\x71\x1f\x5a\x8e\x6b\x17\x65\xf9\x40\xe9\x46\x47\xa4\xb9\x9c\xd4\x0c\x7b\x72\x1e\x81\xe1\xd4\xf7\xad\xed\x51\xd8\x32\xe9\x1a\x88\x2c\xc4\x24\x61\xde\x74\xb2\xd0\x94\x4f\x5f\x45\xbe\xdd\x4b\xbe\xe2\x8f\x6c\x23\xfb\x2f\xbd\xca\x27\x00\x00\xff\xff\xe7\x7c\x74\x3b\x93\x00\x00\x00")

func _0008_relax_fcntSqlBytes() ([]byte, error) {
	return bindataRead(
		__0008_relax_fcntSql,
		"0008_relax_fcnt.sql",
	)
}

func _0008_relax_fcntSql() (*asset, error) {
	bytes, err := _0008_relax_fcntSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0008_relax_fcnt.sql", size: 147, mode: os.FileMode(420), modTime: time.Unix(1475694687, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __0009_adr_interval_and_install_marginSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\xce\xc1\xca\xc2\x40\x0c\x04\xe0\xf3\xbf\x4f\x31\xc7\x5f\x6c\x41\x04\x4f\xbd\xfa\x0a\x9e\x4b\x6c\x62\x59\xc8\x66\x4b\xcc\xea\xeb\x8b\x37\x0b\x7b\x0b\x84\x99\xf9\xc6\x11\xc7\x92\x57\xa7\x10\xdc\xb6\x44\x1a\xe2\x08\xba\xab\xc0\x2a\x4b\xfa\x23\x66\x2c\x55\x5b\x31\x10\xfb\x9c\x2d\xc4\x5f\xa4\xf8\x1e\xab\x38\xac\x06\xac\xa9\x82\xe5\x41\x4d\x03\xa7\x61\x17\xca\xf6\x0c\x52\xa5\xc8\xd5\xe6\x42\xbe\x66\x03\xcb\x92\x0b\xe9\xff\x65\x38\x1f\x3a\x05\x53\x4a\xbf\xac\x6b\x7d\x5b\x07\xc6\x5e\xb7\x9e\x6c\xd8\xbf\x3a\xfb\x53\xfa\x04\x00\x00\xff\xff\x49\xf2\xe1\xba\xf6\x00\x00\x00")

func _0009_adr_interval_and_install_marginSqlBytes() ([]byte, error) {
	return bindataRead(
		__0009_adr_interval_and_install_marginSql,
		"0009_adr_interval_and_install_margin.sql",
	)
}

func _0009_adr_interval_and_install_marginSql() (*asset, error) {
	bytes, err := _0009_adr_interval_and_install_marginSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0009_adr_interval_and_install_margin.sql", size: 246, mode: os.FileMode(420), modTime: time.Unix(1482345058, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __0010_recreate_application_tableSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x92\xc1\x92\xdb\x20\x0c\x86\xcf\xe6\x29\x74\xb3\x3d\x75\x3a\xe9\xd9\x93\x5b\x5f\xa1\x67\x0f\x01\x25\xd1\x14\x0b\x0a\x72\xe3\xbc\xfd\x0e\x24\xd9\x71\x1c\xef\xde\x18\x81\x3e\xfd\xff\x8f\x76\x3b\xf8\x31\xd2\x39\x6a\x41\xf8\x13\x94\x89\x98\x4f\xa2\x8f\x0e\x41\x87\xe0\xc8\x68\x21\xcf\xd0\xa8\x8a\x2c\x1c\xe9\x9c\x30\x92\x76\x10\x22\x8d\x3a\xde\xe0\x2f\xde\x3a\x55\xb1\x1e\x11\xfe\xeb\x68\x2e\x3a\x36\xbf\xf6\xfb\x16\x26\xa6\x7f\x13\x02\x7b\x01\x9e\x9c\xeb\x54\x65\x31\x99\x48\xa1\xd0\x04\x67\xf9\xbc\x03\xd5\xf6\xea\x39\x99\xd8\xe2\x0c\x64\xe7\x61\x31\x7d\x28\x7c\xcf\x4b\x45\x4d\xae\xe5\x46\xe2\x84\x51\x80\x58\xfc\xf2\x5e\x55\xe5\x45\x07\x8b\xb9\xad\xaa\x12\x3a\x34\x02\x96\x92\x10\x1b\x69\x90\x8d\xb7\xd8\xe8\x10\x06\x9c\xa8\x83\xfa\x82\x73\xdd\xb6\xa0\x13\xdc\xdb\xeb\x3a\x9f\x97\xe2\x4f\xd1\x8f\xc0\xde\x62\xaf\x94\x76\x82\xf1\x11\x57\x2e\x81\xaa\xb4\xb5\x60\xbc\x9b\xc6\x17\xb9\xc3\x3d\x3d\x62\x81\x88\x27\x8c\xc8\x06\xd3\x4b\xc2\x9e\xc1\xa2\x43\x41\x30\x3a\x19\x5d\xf0\x53\xb0\x39\x94\x42\x4e\x28\x6b\xe0\x01\x9a\x87\x1f\xb2\x77\x59\x4b\xe0\xf5\x82\x11\x8b\x0b\x38\xc0\xc3\x67\x26\xfd\x5c\x9b\xfd\xc2\x47\x29\x6d\x3b\xc9\x5a\x9e\xdf\xb7\xf5\x77\x19\x31\xac\x7a\x3c\x17\x72\xf3\x5a\xce\xc3\x97\x3b\xf8\xdb\x5f\x59\x29\x1b\x7d\xf8\x1e\xb7\xa1\x59\x55\xa5\x6d\x53\x71\xff\xc6\x5c\xaf\xd7\xf3\xc5\xdb\xea\xf7\xea\x23\x00\x00\xff\xff\xb7\x93\x74\x6a\x24\x03\x00\x00")

func _0010_recreate_application_tableSqlBytes() ([]byte, error) {
	return bindataRead(
		__0010_recreate_application_tableSql,
		"0010_recreate_application_table.sql",
	)
}

func _0010_recreate_application_tableSql() (*asset, error) {
	bytes, err := _0010_recreate_application_tableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0010_recreate_application_table.sql", size: 804, mode: os.FileMode(420), modTime: time.Unix(1486123701, 0)}
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
	"0001_initial.sql": _0001_initialSql,
	"0002_join_accept_params.sql": _0002_join_accept_paramsSql,
	"0003_rx_window_and_rx2_dr.sql": _0003_rx_window_and_rx2_drSql,
	"0004_add_node_apps_nwks_key_name_devaddr.sql": _0004_add_node_apps_nwks_key_name_devaddrSql,
	"0005_add_queue.sql": _0005_add_queueSql,
	"0006_remove_application_table.sql": _0006_remove_application_tableSql,
	"0007_migrate_channels_to_channel_list.sql": _0007_migrate_channels_to_channel_listSql,
	"0008_relax_fcnt.sql": _0008_relax_fcntSql,
	"0009_adr_interval_and_install_margin.sql": _0009_adr_interval_and_install_marginSql,
	"0010_recreate_application_table.sql": _0010_recreate_application_tableSql,
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
	"0001_initial.sql": &bintree{_0001_initialSql, map[string]*bintree{}},
	"0002_join_accept_params.sql": &bintree{_0002_join_accept_paramsSql, map[string]*bintree{}},
	"0003_rx_window_and_rx2_dr.sql": &bintree{_0003_rx_window_and_rx2_drSql, map[string]*bintree{}},
	"0004_add_node_apps_nwks_key_name_devaddr.sql": &bintree{_0004_add_node_apps_nwks_key_name_devaddrSql, map[string]*bintree{}},
	"0005_add_queue.sql": &bintree{_0005_add_queueSql, map[string]*bintree{}},
	"0006_remove_application_table.sql": &bintree{_0006_remove_application_tableSql, map[string]*bintree{}},
	"0007_migrate_channels_to_channel_list.sql": &bintree{_0007_migrate_channels_to_channel_listSql, map[string]*bintree{}},
	"0008_relax_fcnt.sql": &bintree{_0008_relax_fcntSql, map[string]*bintree{}},
	"0009_adr_interval_and_install_margin.sql": &bintree{_0009_adr_interval_and_install_marginSql, map[string]*bintree{}},
	"0010_recreate_application_table.sql": &bintree{_0010_recreate_application_tableSql, map[string]*bintree{}},
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

