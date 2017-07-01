package grandet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBarn(t *testing.T) {

	genGrandet(t, "github.com/setekhid/grandet02")

	asset_content := Asset("github.com/setekhid/grandet02/test.txt")
	assert.EqualValues(t, []byte("I'm an asset!"), asset_content)

	result := Foldl(

		[]byte("I'm an asset!"),

		func(value interface{}, name string, content []byte) interface{} {

			if found, ok := value.(bool); ok {
				return found
			}

			require := value.([]byte)
			if "/github.com/setekhid/grandet02/test.txt" == name {
				return true
			}

			return require
		},
	)

	assert.EqualValues(t, true, result)

	assert.EqualValues(t, Branches("github.com"), []string{"/setekhid"})
	t.Log(Branches("github.com/setekhid"))
	assert.Contains(t, Branches("github.com/setekhid"), "/grandet02")
}
