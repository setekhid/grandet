package main

import (
	"encoding/json"
	"go/build"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPkgNames(t *testing.T) {

	pkg, err := build.ImportDir("../", 0)
	require.NoError(t, err)
	pkg_str, err := json.Marshal(pkg)
	require.NoError(t, err)

	t.Log(string(pkg_str))

	pkg_name, pkg_import, err := PkgNames("../")
	assert.NoError(t, err)

	assert.EqualValues(t, "grandet", pkg_name)

	t.Log(pkg_import)
}
