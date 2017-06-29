package assets

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAssetTmplZipping(t *testing.T) {
	testTmplZipping(t, "./asset.go.tmpl")
}

func TestGrandetTmplZipping(t *testing.T) {
	testTmplZipping(t, "./grandet.go.tmpl")
}

func testTmplZipping(t *testing.T, name string) {

	buff := &bytes.Buffer{}

	// encoding
	func() {

		file, err := os.Open(name)
		assert.NoError(t, err)
		defer file.Close()

		info, err := file.Stat()
		assert.NoError(t, err)
		size := info.Size()
		t.Log("file size: ", size)

		writer := gzip.NewWriter(buff)
		defer writer.Close()

		written, err := io.Copy(writer, file)
		assert.NoError(t, err)
		require.EqualValues(t, size, written)

		err = writer.Flush()
		assert.NoError(t, err)
	}()

	// decoding
	func() {

		bytes := BytesAsset(buff.Bytes())
		t.Log(bytes)

		raw, err := gzip.NewReader(buff)
		assert.NoError(t, err)
		str, err := ioutil.ReadAll(raw)
		assert.NoError(t, err)
		t.Log(string(str))
	}()
}
