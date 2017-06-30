package grandet

import (
	"bytes"
	"compress/gzip"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func genGrandet(t *testing.T) *AssetsImpl {

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
	ga.Init()

	ga.RegistAsset("test.txt", zipped_content)

	return ga
}

func TestGrandetAsset(t *testing.T) {

	ga := genGrandet(t)

	asset_content := ga.Asset("test.txt")

	require.EqualValues(t, []byte("I'm an asset!"), asset_content)
}

func TestGrandetFoldl(t *testing.T) {

	ga := genGrandet(t)

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
