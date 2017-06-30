package assets

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/setekhid/grandet"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBytesAsset(t *testing.T) {

	asset := BytesAsset([]byte{
		1, 2, 3, 4, 5, 6, 77,
	})
	asset_str := asset.String()

	assert.EqualValues(t, "1, 2, 3, 4, 5, 6, 77,", asset_str)

	t.Log("hash:", asset.UniqueName())
}

func TestAssetsResult(t *testing.T) {

	grandet_raw, err := ioutil.ReadFile("grandet.go.tmpl")
	require.NoError(t, err)
	asset_raw, err := ioutil.ReadFile("asset.go.tmpl")
	require.NoError(t, err)

	grandet_asset := grandet.Asset(
		"github.com/setekhid/grandet/assets/grandet.go.tmpl")
	asset_asset := grandet.Asset(
		"github.com/setekhid/grandet/assets/asset.go.tmpl")

	assert.EqualValues(t, grandet_raw, grandet_asset)
	assert.EqualValues(t, asset_raw, asset_asset)
}

func TestBarnResult(t *testing.T) {

	grandet_raw, err := ioutil.ReadFile("grandet.go.tmpl")
	require.NoError(t, err)
	asset_raw, err := ioutil.ReadFile("asset.go.tmpl")
	require.NoError(t, err)

	grandet_asset := Grandet.Asset("grandet.go.tmpl")
	asset_asset := Grandet.Asset("asset.go.tmpl")

	assert.EqualValues(t, grandet_raw, grandet_asset)
	assert.EqualValues(t, asset_raw, asset_asset)
}

func TestAssetTmplZipping(t *testing.T) {
	t.SkipNow()
	testTmplZipping(t, "./asset.go.tmpl")
}

func TestGrandetTmplZipping(t *testing.T) {
	t.SkipNow()
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
