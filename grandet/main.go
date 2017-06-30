package main

import (
	"bytes"
	"compress/gzip"
	"flag"
	"go/build"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

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

	// check pkg_name & pkg_import
	if len(pkg_name) <= 0 || len(pkg_import) <= 0 {

		var err error
		build_name, build_import, err := PkgName(pkg_dir)
		if err != nil {
			panic(err)
		}

		if len(pkg_name) <= 0 {
			pkg_name = build_name
		}
		if len(pkg_import) <= 0 {
			pkg_import = build_import
		}
	}

	// range files
	infos := []*assets.AssetInfo{}
	for _, file := range pkg_files {

		asset, err := ReadAsset(filepath.Join(pkg_dir, file))
		if err != nil {
			panic(err)
		}

		info := assets.NewAssetInfo()
		info.AssetName = file
		info.AssetContent = asset
		info.AssetPackage = pkg_name
		info.PackageImport = pkg_import
		info.AssetRegister = "registAsset" + asset.UniqueName()

		infos = append(infos, info)

		// write down
		err = ioutil.WriteFile(
			filepath.Join(pkg_dir, file+".go"),
			[]byte(info.String()),
			os.ModePerm,
		)
		if err != nil {
			panic(err)
		}
	}

	all, err := assets.NewAssetsInfo(infos)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(
		filepath.Join(pkg_dir, "grandet.go"),
		[]byte(all.String()),
		os.ModePerm,
	)
	if err != nil {
		panic(err)
	}
}

func PkgName(pkg_dir string) (string, string, error) {

	pkg, err := build.ImportDir(pkg_dir, 0)
	if err != nil {
		return "", "", err
	}

	return pkg.Name, pkg.ImportPath, err
}

func ReadAsset(file string) (assets.BytesAsset, error) {

	buff := &bytes.Buffer{}

	err := func() error {

		reader, err := os.Open(file)
		if err != nil {
			return err
		}
		defer reader.Close()

		writer := gzip.NewWriter(buff)
		defer writer.Close()

		_, err = io.Copy(writer, reader)
		if err != nil {
			return err
		}
		err = writer.Flush()
		if err != nil {
			return err
		}

		return nil
	}()

	return buff.Bytes(), err
}
