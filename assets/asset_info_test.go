package assets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytesAsset(t *testing.T) {

	asset := BytesAsset([]byte{
		1, 2, 3, 4, 5, 6, 77,
	})
	asset_str := asset.String()

	assert.EqualValues(t, "1, 2, 3, 4, 5, 6, 77,", asset_str)
}
