package http

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOSReadDir(t *testing.T) {

	dir, err := os.Open("./testdata/dir")
	assert.NoError(t, err)

	len, err := dir.Read(make([]byte, 1024))
	t.Log(len, err)
	t.Log(reflect.TypeOf(err))

	perr := err.(*os.PathError)
	t.Log(reflect.TypeOf(perr.Err))
}

func TestOSSeekDir(t *testing.T) {

	dir, err := os.Open("./testdata/dir")
	assert.NoError(t, err)

	ret, err := dir.Seek(12, 2)
	t.Log(ret, err)
	t.Log(reflect.TypeOf(err))

	//perr := err.(*os.PathError)
	//t.Log(reflect.TypeOf(perr.Err))
}
