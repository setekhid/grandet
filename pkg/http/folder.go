package http

import (
	"io"
	"os"
	"time"
)

type barnFolder struct {
	info *barnFolderInfo
}

func newBarnFolder(name string) *barnFolder {
	return &barnFolder{
		info: &barnFolderInfo{
			name: name,
			// FIXME assign time
		},
	}
}

func (bf *barnFolder) Close() error { return nil }

func (bf *barnFolder) Stat() (os.FileInfo, error) {
	return bf.info, nil
}

func (bf *barnFolder) Seek(offset int64, whence int) (int64, error) {
	return 0, os.ErrPermission
}

func (bf *barnFolder) Read(b []byte) (int, error) {
	return 0, os.ErrPermission
}

func (bf *barnFolder) Readdir(count int) ([]os.FileInfo, error) {
	return nil, io.EOF
}

type barnFolderInfo struct {
	name string
	time time.Time
}

func (info *barnFolderInfo) Name() string       { return info.name }
func (info *barnFolderInfo) Size() int64        { return 0 }
func (info *barnFolderInfo) Mode() os.FileMode  { return os.ModePerm }
func (info *barnFolderInfo) ModTime() time.Time { return info.time }
func (info *barnFolderInfo) IsDir() bool        { return true }
func (info *barnFolderInfo) Sys() interface{}   { return nil }
