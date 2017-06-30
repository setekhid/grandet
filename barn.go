package grandet

import (
	"path"
)

// Barn collects all imported Grandet assets
type Barn interface {
	Grandet
	Record(pkg_import string, assets Grandet)
}

type barnImpl struct {
	grandets map[string]Grandet
}

func newBarnImpl() *barnImpl {
	return &barnImpl{grandets: map[string]Grandet{}}
}

// Grandet#Asset
func (b *barnImpl) Asset(name string) []byte {

	name = path.Clean(name)
	pkg_import := path.Dir(name)
	asset_name := path.Base(name)

	if ga, ok := b.grandets[pkg_import]; ok {
		return ga.Asset(asset_name)
	}

	return nil
}

// Grandet#Foldl
func (b *barnImpl) Foldl(

	value interface{},
	process func(interface{}, string, []byte) interface{},

) interface{} {

	result := value

	for pkg_import, pkg_grandet := range b.grandets {

		result = pkg_grandet.Foldl(

			result,

			func(value interface{}, name string, content []byte) interface{} {
				return process(
					value,
					path.Join(pkg_import, name),
					content,
				)
			},
		)
	}

	return result
}

// Barn#Record
func (b *barnImpl) Record(pkg_import string, assets Grandet) {

	pkg_import = path.Clean(pkg_import)

	if _, ok := b.grandets[pkg_import]; ok {
		panic("duplicated package assets recording " + pkg_import)
	}

	b.grandets[pkg_import] = assets
}

var barn Barn = newBarnImpl()

// Asset return an asset from Barn
func Asset(name string) []byte { return barn.Asset(name) }

// Foldl loop all assets in Barn
func Foldl(

	value interface{},
	process func(interface{}, string, []byte) interface{},

) interface{} {
	return barn.Foldl(value, process)
}

func (ga *GrandetAssets) barnRegist(pkg_import string) {
	barn.Record(pkg_import, ga)
}
