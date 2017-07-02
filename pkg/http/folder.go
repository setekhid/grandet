package http

import (
	"errors"
	"math"
	"os"
	"time"

	"github.com/setekhid/grandet"
)

type barnFolder struct {
	info *barnFolderInfo

	branches []string
	files    []string
}

func newBarnFolder(
	name string, branches, files []string, modtime time.Time) *barnFolder {

	return &barnFolder{
		info: &barnFolderInfo{
			name:    name,
			modtime: modtime,
		},
		branches: branches,
		files:    files,
	}
}

func (bf *barnFolder) Close() error { return nil }

func (bf *barnFolder) Stat() (os.FileInfo, error) {
	return bf.info, nil
}

func (bf *barnFolder) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

func (bf *barnFolder) Read(b []byte) (int, error) {
	return 0, &os.PathError{
		Op:   "read",
		Path: bf.info.name,
		Err:  errors.New("is a directory"),
	}
}

func (bf *barnFolder) Readdir(count int) ([]os.FileInfo, error) {

	infos := []os.FileInfo{}

	for _, branch := range bf.branches {
		infos = append(infos, &barnFolderInfo{
			name:    branch,
			modtime: grandet.ModBeginningTime,
		})
	}

	for _, file := range bf.files {
		infos = append(infos, &barnFileInfo{
			name:    file,
			size:    -1,
			modtime: grandet.ModBeginningTime,
		})
	}

	if count > 0 {
		return infos[:int(
			math.Min(float64(count), float64(len(infos))),
		)], nil
	}

	return infos, nil
}

type barnFolderInfo struct {
	name    string
	modtime time.Time
}

func (info *barnFolderInfo) Name() string       { return info.name }
func (info *barnFolderInfo) Size() int64        { return 0 }
func (info *barnFolderInfo) Mode() os.FileMode  { return os.ModePerm }
func (info *barnFolderInfo) ModTime() time.Time { return info.modtime }
func (info *barnFolderInfo) IsDir() bool        { return true }
func (info *barnFolderInfo) Sys() interface{}   { return nil }
