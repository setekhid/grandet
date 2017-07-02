package http

import (
	"net/http"

	"github.com/setekhid/grandet"
)

// FS presents the instance of http#FileSystem
var Fs http.FileSystem = barnFileSystem{}

type barnFileSystem struct{}

// Open implement Filesystem#Open
func (fs barnFileSystem) Open(name string) (http.File, error) {

	content := grandet.Asset(name)
	if content != nil {
		modtime := grandet.ModTime(name)
		return newBarnFile(name, content, modtime), nil
	}

	branches := grandet.Branches(name)
	files := grandet.Grandet(name).FoldlNames(
		[]string{},
		func(value interface{}, name string) interface{} {
			return append(value.([]string), name)
		},
	).([]string)

	return newBarnFolder(name, branches, files, grandet.ModBeginningTime), nil
}
