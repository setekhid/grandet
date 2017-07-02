package grandet

import (
	"bytes"
	"compress/gzip"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func genGrandet(t *testing.T, pkg_import string) *AssetsImpl {

	raw_content := []byte("I'm an asset!")

	writer := &bytes.Buffer{}
	func() {
		encoder := gzip.NewWriter(writer)
		defer encoder.Close()
		_, err := io.Copy(encoder, bytes.NewBuffer(raw_content))
		assert.NoError(t, err)
		err = encoder.Flush()
		assert.NoError(t, err)
	}()

	zipped_content := writer.Bytes()

	ga := new(AssetsImpl)
	ga.Init(pkg_import)

	ga.RegistAsset("test.txt", zipped_content, time.Now())

	return ga
}

func TestGrandetAsset(t *testing.T) {

	ga := genGrandet(t, "github.com/setekhid/grandet00")

	asset_content := ga.Asset("test.txt")

	require.EqualValues(t, []byte("I'm an asset!"), asset_content)
}

func TestGrandetFoldl(t *testing.T) {

	ga := genGrandet(t, "github.com/setekhid/grandet01")

	result := ga.Foldl(

		[]byte("I'm an asset!"),

		func(value interface{}, name string, content []byte) interface{} {

			if found, ok := value.(bool); ok {
				return found
			}

			require := value.([]byte)
			if bytes.Equal(require, content) {
				return true
			}

			return require
		},
	)

	assert.EqualValues(t, true, result)
}

func TestGrandetFoldlNames(t *testing.T) {

	ga := genGrandet(t, "github.com/setekhid/grandet03")

	result := ga.FoldlNames(

		false,

		func(value interface{}, name string) interface{} {

			if found := value.(bool); found {
				return found
			}

			if "/test.txt" == name {
				return true
			}

			return value
		},
	)

	assert.EqualValues(t, true, result)
}
