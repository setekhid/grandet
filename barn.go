package grandet

import (
	"path"
	"sort"
	"strings"
	"time"
)

// Barn collects all imported Grandet assets
type Barn interface {
	Assets
	Grandet(pkg_import string) Assets
	Branches(pkg_import string) []string
}

type barnImpl struct {
	grandets map[string]Assets
	branches map[string][]string
}

func newBarnImpl() *barnImpl {
	return &barnImpl{
		grandets: map[string]Assets{},
		branches: map[string][]string{},
	}
}

// Assets#Asset
func (b *barnImpl) Asset(name string) []byte {

	name = pathFormatAndCheck(name)
	pkg_import := path.Dir(name)
	asset_name := path.Base(name)

	if ga, ok := b.grandets[pkg_import]; ok {
		return ga.Asset(asset_name)
	}

	return nil
}

// Assets#ModTime
func (b *barnImpl) ModTime(name string) time.Time {

	name = pathFormatAndCheck(name)
	pkg_import := path.Dir(name)
	asset_name := path.Base(name)

	if ga, ok := b.grandets[pkg_import]; ok {
		return ga.ModTime(asset_name)
	}

	return ModBeginningTime
}

// Assets#Foldl
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

// Assets#FoldlNames
func (b *barnImpl) FoldlNames(

	value interface{},
	process func(interface{}, string) interface{},

) interface{} {

	result := value

	for pkg_import, pkg_grandet := range b.grandets {

		result = pkg_grandet.FoldlNames(
			result,
			func(value interface{}, name string) interface{} {
				return process(value, path.Join(pkg_import, name))
			},
		)
	}

	return result
}

// Barn#Grandet
func (b *barnImpl) Grandet(pkg_import string) Assets {

	pkg_import = pathFormatAndCheck(pkg_import)

	return b.grandets[pkg_import]
}

// Barn#Branches
func (b *barnImpl) Branches(pkg_import string) []string {

	pkg_import = pathFormatAndCheck(pkg_import)

	branches := b.branches[pkg_import]
	if false {
		copied := make([]string, len(branches))
		copy(copied, branches)
		return copied
	}

	return branches
}

var (
	barn *barnImpl = newBarnImpl() // singleton instance

	// Asset return an asset from the Barn instance
	Asset = barn.Asset
	// ModTime return the modified time of an asset
	ModTime = barn.ModTime
	// Foldl loop all assets in Barn
	Foldl = barn.Foldl
	// FoldlNames loop all asset names in Barn
	FoldlNames = barn.FoldlNames
	// Grandet return the Grandet object of the asset package
	Grandet = barn.Grandet
	// Branches return the branch names of specific path
	Branches = barn.Branches
)

func (ga *AssetsImpl) barnRegist() { barn.record(ga.pkg_import, ga) }

func (b *barnImpl) record(pkg_import string, assets Assets) {

	pkg_import = path.Clean(pkg_import)

	if _, ok := b.grandets[pkg_import]; ok {
		panic("duplicated package assets recording " + pkg_import)
	}

	b.grandets[pkg_import] = assets
}

func (ga *AssetsImpl) linkParent() {

	pkg_import := ga.pkg_import
	parent_dir := path.Dir(ga.pkg_import)
	for pkg_import != "/" {

		barn.link(pkg_import, parent_dir)
		pkg_import = parent_dir
		parent_dir = path.Dir(pkg_import)
	}
}

func (b *barnImpl) link(pkg_import string, parent_import string) {

	pkg_import = path.Clean(pkg_import)
	parent_import = path.Clean(parent_import)

	parent_dir := parent_import + "/"
	if parent_import == "/" {
		parent_dir = parent_import
	}

	if !strings.HasPrefix(pkg_import, parent_dir) {
		panic(pkg_import + " is not belong to the " + parent_import)
	}

	relative_import := pkg_import[len(parent_dir)-1:]

	branches := b.branches[parent_import]

	if !sort.StringsAreSorted(branches) {
		panic("internal error, barn branches doesn't sort")
	}

	ind := sort.SearchStrings(branches, relative_import)
	if ind < len(branches) && branches[ind] == relative_import {
		return
	}

	branches = append(branches, "")
	copy(branches[ind+1:], branches[ind:])
	branches[ind] = relative_import

	b.branches[parent_import] = branches
}
