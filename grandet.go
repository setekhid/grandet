package grandet

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

// Grandet is the interface for directly asset accessing, the return value may
// not be your original content.
type Grandet interface {

	// Asset get an asset by name, if doesn't exists, return nil
	Asset(name string) (content []byte)

	// Foldl fold all assets in this Grandet with value
	Foldl(
		value interface{},
		process func(
			value interface{}, name string, content []byte,
		) (result interface{}),
	) (result interface{})
}

// GrandetAssets provide the implementation of Grandet
type GrandetAssets struct {
	zipped map[string][]byte
	raw    map[string][]byte
}

// Init initialize GrandetAssets
func (ga *GrandetAssets) Init() {
	ga.zipped = map[string][]byte{}
	ga.raw = map[string][]byte{}
}

func (ga *GrandetAssets) unzipped(zipped []byte) []byte {

	reader := bytes.NewBuffer(zipped)
	decoder, err := gzip.NewReader(reader)
	if err != nil {
		panic(err)
	}
	defer decoder.Close()

	unzipped, err := ioutil.ReadAll(decoder)
	if err != nil {
		panic(err)
	}

	return unzipped
}

// Grandet#Asset
func (ga *GrandetAssets) Asset(name string) []byte {

	if asset, ok := ga.raw[name]; ok {
		return asset
	}

	if zipped, ok := ga.zipped[name]; ok {
		raw := ga.unzipped(zipped)
		ga.raw[name] = raw
		delete(ga.zipped, name)
		return raw
	}

	return nil
}

// Grandet#Foldl
func (ga *GrandetAssets) Foldl(

	value interface{},
	process func(interface{}, string, []byte) interface{},

) interface{} {

	result := value

	// range raw
	for name, content := range ga.raw {
		result = process(result, name, content)
	}
	// range zipped
	for name, zipped := range ga.zipped {
		result = process(result, name, ga.unzipped(zipped))
	}

	return result
}

// RegistAsset register an asset into grandet
func (ga *GrandetAssets) RegistAsset(name string, content []byte) {
	ga.zipped[name] = content
}
