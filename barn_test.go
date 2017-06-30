package grandet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBarn(t *testing.T) {

	grandet := genGrandet(t)
	grandet.barnRegist("github.com/setekhid/grandet")

	asset_content := Asset("github.com/setekhid/grandet/test.txt")
	assert.EqualValues(t, []byte("I'm an asset!"), asset_content)

	result := Foldl(

		[]byte("I'm an asset!"),

		func(value interface{}, name string, content []byte) interface{} {

			if found, ok := value.(bool); ok {
				return found
			}

			require := value.([]byte)
			if "github.com/setekhid/grandet/test.txt" == name {
				return true
			} else {
				return require
			}
		},
	)

	assert.EqualValues(t, true, result)
}
