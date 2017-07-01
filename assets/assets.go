package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// AssetInfo presents an asset file information
type AssetInfo struct {
	AssetPackage  string
	PackageImport string

	AssetName     string
	AssetContent  BytesAsset
	AssetModTime  int64
	AssetRegister string
}

// ReadAssetInfo read an asset file to generate AssetInfo
func ReadAssetInfo(asset_file string) (*AssetInfo, error) {

	asset_name := filepath.Base(asset_file)
	asset_content := bytes.Buffer{}
	asset_modtime := int64(0)

	// zipping file content into asset_content
	err := func() error {

		zipper := gzip.NewWriter(&asset_content)
		defer zipper.Close()

		file, err := os.Open(asset_file)
		if err != nil {
			return err
		}
		defer file.Close()

		fstat, err := file.Stat()
		if err != nil {
			return err
		}
		asset_modtime = fstat.ModTime().Unix()

		_, err = io.Copy(zipper, file)
		if err != nil {
			return err
		}

		return zipper.Flush()
	}()
	if err != nil {
		return nil, err
	}

	return NewAssetInfo(asset_name, asset_content.Bytes(), asset_modtime), nil
}

// NewAssetInfo construct an AssetInfo, with empty informations
func NewAssetInfo(name string, content []byte, modtime int64) *AssetInfo {
	return &AssetInfo{
		AssetName:     name,
		AssetContent:  BytesAsset(content),
		AssetModTime:  modtime,
		AssetRegister: "registAsset" + BytesAsset(content).UniqueName(),
	}
}

// Codes return AssetInfo parsed golang code
func (info *AssetInfo) Codes() []byte { return asset_go_tmpl.parse(info) }

// WriteFile write Codes() down to file
func (info *AssetInfo) WriteFile(file_name string) error {
	return ioutil.WriteFile(file_name, info.Codes(), os.ModePerm)
}

// BytesAsset contains the content of package asset
type BytesAsset []byte

// UniqueName calculate a unique name for the content
func (ba BytesAsset) UniqueName() string {
	h := md5.New()
	_, err := bytes.NewBuffer(ba).WriteTo(h)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Stringer#String
func (ba BytesAsset) String() string {
	str := ""
	for _, b := range ba {
		str = str + strconv.Itoa(int(b)) + ", "
	}
	str = strings.TrimSpace(str)
	return str
}

// AssetPackage presents the package assets information
type AssetPackage struct {
	AssetPackage   string
	PackageImport  string
	AssetRegisters []string
}

// NewAssetPackage create an AssetPackage
func NewAssetPackage(pkg_name, pkg_import string) *AssetPackage {
	return &AssetPackage{AssetPackage: pkg_name, PackageImport: pkg_import}
}

// Collect collects all asset belong to it
func (pkg *AssetPackage) Collect(asset *AssetInfo) {
	asset.AssetPackage = pkg.AssetPackage
	asset.PackageImport = pkg.PackageImport
	pkg.AssetRegisters = append(pkg.AssetRegisters, asset.AssetRegister)
}

// Codes return AssetInfo parsed golang code
func (pkg *AssetPackage) Codes() []byte {
	sort.Strings(pkg.AssetRegisters)
	return grandet_go_tmpl.parse(pkg)
}

// WriteFile write Codes() down to file
func (pkg *AssetPackage) WriteFile(file_name string) error {
	return ioutil.WriteFile(file_name, pkg.Codes(), os.ModePerm)
}
