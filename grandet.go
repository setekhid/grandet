package grandet

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"time"
)

// Assets is the interface for directly asset accessing, the return value may
// not be your original content.
type Assets interface {

	// Asset get an asset by name, if doesn't exists, return nil
	Asset(name string) (content []byte)

	// Foldl fold all assets in this Assets with value
	Foldl(
		value interface{},
		process func(
			value interface{}, name string, content []byte,
		) (result interface{}),
	) (result interface{})
}

// AssetsImpl provide the implementation of Assets
type AssetsImpl struct {
	zipped  map[string][]byte
	raw     map[string][]byte
	modtime map[string]time.Time
}

// Init initialize AssetsImpl
func (ga *AssetsImpl) Init(args ...string) {
	ga.zipped = map[string][]byte{}
	ga.raw = map[string][]byte{}
	ga.modtime = map[string]time.Time{}

	if len(args) > 0 {
		ga.barnRegist(args[0])
	}
}

func (ga *AssetsImpl) unzipped(zipped []byte) []byte {

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

// Asset implement Assets#Asset
func (ga *AssetsImpl) Asset(name string) []byte {

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

// Foldl implement Assets#Foldl
func (ga *AssetsImpl) Foldl(

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
func (ga *AssetsImpl) RegistAsset(
	name string, content []byte, modtime time.Time,
) {
	ga.zipped[name] = content
	ga.modtime[name] = modtime
}
