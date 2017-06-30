package main

import (
	"flag"
	"go/build"
	"path/filepath"
	"strings"

	"github.com/setekhid/grandet/assets"
)

func main() {

	var pkg_dir, pkg_import, pkg_name string

	// parsing parameters
	flag.StringVar(&pkg_dir, "d", "", "package directory")
	flag.StringVar(&pkg_import, "i", "",
		"package import path, if empty, will read the directory to find out")
	flag.StringVar(&pkg_name, "n", "",
		"package name, if empty, will read the directory to find out")
	flag.Parse()

	// asset files
	pkg_files := flag.Args()

	// Abs pkg_dir
	pkg_dir, err := filepath.Abs(pkg_dir)
	if err != nil {
		panic(err)
	}

	// check pkg_name & pkg_import
	if len(pkg_name) <= 0 || len(pkg_import) <= 0 {

		read_name, read_import, err := PkgNames(pkg_dir)
		if err != nil {
			panic(err)
		}

		if len(pkg_name) <= 0 {
			pkg_name = read_name
		}
		if len(pkg_import) <= 0 {
			pkg_import = read_import
		}
	}

	// generate AssetPackage
	asset_pkg := assets.NewAssetPackage(pkg_name, pkg_import)

	// range files to generate asset informations
	asset_infos := map[string]*assets.AssetInfo{}
	for _, file_name := range pkg_files {

		file_name, err := filepath.Abs(file_name)
		if err != nil {
			panic(err)
		}
		if !strings.HasPrefix(file_name, pkg_dir) {
			panic("shit, your asset file doesn't belong to the package")
		}

		asset_info, err := assets.ReadAssetInfo(file_name)
		if err != nil {
			panic(err)
		}

		// collect
		asset_pkg.Collect(asset_info)
		asset_infos[file_name] = asset_info
	}

	// write down asset golang files
	for file_name, asset_info := range asset_infos {
		err := asset_info.WriteFile(file_name + ".go")
		if err != nil {
			panic(err)
		}
	}

	// write down grandet.go
	err = asset_pkg.WriteFile(filepath.Join(pkg_dir, "grandet.go"))
	if err != nil {
		panic(err)
	}
}

// PkgNames parse package name and import name
func PkgNames(pkg_dir string) (string, string, error) {

	pkg, err := build.ImportDir(pkg_dir, 0)
	if err != nil {
		return "", "", err
	}

	return pkg.Name, pkg.ImportPath, err
}
