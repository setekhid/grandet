package grandet

import ()

// Barn collects all Grandet assets
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
	// TODO
	return nil
}

// Grandet#Foldl
func (b *barnImpl) Foldl(

	value interface{},
	process func(interface{}, string, []byte) interface{},

) interface{} {
	// TODO
	return nil
}

// Barn#Record
func (b *barnImpl) Record(pkg_import string, assets Grandet) {
	// TODO
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

func (ga *GrandetAssets) barnRegist() { barn.Record(ga.pkg_import, ga) }
