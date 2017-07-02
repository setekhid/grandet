package http

import (
	"bytes"
	"errors"
	"os"
	"time"
)

type barnFile struct {
	reader *bytes.Reader
	info   *barnFileInfo
}

func newBarnFile(name string, content []byte, modtime time.Time) *barnFile {
	return &barnFile{
		reader: bytes.NewReader(content),
		info: &barnFileInfo{
			name:    name,
			size:    int64(len(content)),
			modtime: modtime,
		},
	}
}

func (bf *barnFile) Close() error { return nil }

func (bf *barnFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, errors.New(bf.info.name + " is not a directory")
}

func (bf *barnFile) Stat() (os.FileInfo, error) {
	return bf.info, nil
}

func (bf *barnFile) Seek(offset int64, whence int) (int64, error) {
	return bf.reader.Seek(offset, whence)
}

func (bf *barnFile) Read(b []byte) (int, error) {
	return bf.reader.Read(b)
}

type barnFileInfo struct {
	name    string
	size    int64
	modtime time.Time
}

func (info *barnFileInfo) Name() string       { return info.name }
func (info *barnFileInfo) Size() int64        { return info.size }
func (info *barnFileInfo) Mode() os.FileMode  { return os.ModePerm }
func (info *barnFileInfo) ModTime() time.Time { return info.modtime }
func (info *barnFileInfo) IsDir() bool        { return false }
func (info *barnFileInfo) Sys() interface{}   { return nil }
