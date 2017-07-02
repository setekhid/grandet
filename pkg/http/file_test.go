package http

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOSReadDirFile(t *testing.T) {

	dir, err := os.Open("./testdata/dir/hello.txt")
	assert.NoError(t, err)

	infos, err := dir.Readdir(0)
	t.Log(infos == nil)
	t.Log(infos, err)
	t.Log(reflect.TypeOf(err))

	//perr := err.(*os.PathError)
	//t.Log(reflect.TypeOf(perr.Err))
}
