package grandet

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

// Grandet is the interface for directly asset accessing, the return value may
// not be your original content.
type Grandet interface {
	Asset(name string) (content []byte)
}

type GrandetAssets struct {
	zipped map[string][]byte
	raw    map[string][]byte
}

func (ga *GrandetAssets) Init() {
	ga.zipped = map[string][]byte{}
	ga.raw = map[string][]byte{}
}

// Grandet#Asset
func (ga *GrandetAssets) Asset(name string) []byte {

	if asset, ok := ga.raw[name]; ok {
		return asset
	}

	if zipped, ok := ga.zipped[name]; ok {

		reader := bytes.NewBuffer(zipped)
		decoder, err := gzip.NewReader(reader)
		if err != nil {
			panic(err)
		}
		defer decoder.Close()

		raw, err := ioutil.ReadAll(decoder)
		if err != nil {
			panic(err)
		}

		ga.raw[name] = raw
		delete(ga.zipped, name)

		return raw
	}

	return nil
}

// RegistAsset register an asset into grandet
func (ga *GrandetAssets) RegistAsset(name string, content []byte) {
	ga.zipped[name] = content
}
