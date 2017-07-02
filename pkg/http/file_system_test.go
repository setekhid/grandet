package http

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	_ "github.com/setekhid/grandet/assets"

	"github.com/stretchr/testify/assert"
)

func TestFileSystem(t *testing.T) {

	file, err := FS.Open("github.com/setekhid/grandet/assets/grandet.go.tmpl")
	assert.NoError(t, err)
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	assert.NoError(t, err)

	osfile, err := os.Open("../../assets/grandet.go.tmpl")
	assert.NoError(t, err)
	defer osfile.Close()

	oscontent, err := ioutil.ReadAll(osfile)
	assert.NoError(t, err)

	// t.Log(string(content))
	assert.EqualValues(t, oscontent, content)
}

func TestHTTPFS(t *testing.T) {

	handler := http.FileServer(FS)

	req := httptest.NewRequest(
		http.MethodGet,
		"http://example.com/github.com/setekhid/grandet/assets/asset.go.tmpl",
		nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	resp := rec.Result()
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	defer resp.Body.Close()

	t.Log(resp.Status, resp.StatusCode)
	t.Log(resp.Header.Get("Content-Type"))
	t.Log(string(body))

	file, err := os.Open("../../assets/asset.go.tmpl")
	assert.NoError(t, err)
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	assert.NoError(t, err)

	assert.EqualValues(t, content, body)
}
