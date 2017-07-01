package http

import (
	"net/http"
	"os"

	"github.com/setekhid/grandet"
)

// BarnFS return the instance of http#FileSystem
func BarnFS() http.FileSystem { return barnFileSystem{} }

type barnFileSystem struct{}

// Open implement Filesystem#Open
func (fs barnFileSystem) Open(name string) (http.File, error) {

	content := grandet.Asset(name)
	if content == nil {
		return nil, os.ErrNotExist
	}

	return newBarnFile(name, content), nil
}
