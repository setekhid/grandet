package assets

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
	"text/template"
)

// AssetInfo presents an asset file information
type AssetInfo struct {
	tmpl *template.Template

	AssetPackage  string
	PackageImport string
	AssetName     string
	AssetContent  BytesAsset
	AssetRegister string
}

// NewAssetInfo construct an AssetInfo, with empty informations
func NewAssetInfo() *AssetInfo {

	tmpl_str := Grandet.Asset("asset.go.tmpl")
	tmpl := template.New("asset.go.tmpl")
	tmpl = template.Must(tmpl.Parse(string(tmpl_str)))

	return &AssetInfo{
		tmpl: tmpl,
	}
}

// BytesAsset contains the content of package asset
type BytesAsset []byte

// Stringer#String
func (ba BytesAsset) String() string {
	str := ""
	for _, b := range ba {
		str = str + strconv.Itoa(int(b)) + ", "
	}
	str = strings.TrimSpace(str)
	return str
}

// Stringer#String
func (as *AssetInfo) String() string {
	buff := &bytes.Buffer{}
	err := as.tmpl.Execute(buff, as)
	if err != nil {
		panic(err)
	}
	return buff.String()
}

// AssetsInfo presents the package assets information
type AssetsInfo struct {
	tmpl *template.Template

	AssetPackage   string
	PackageImport  string
	AssetRegisters []string
}

// NewAssetsInfo create an AssetsInfo from AssetInfo-s, if they don't contained
// in the same package, an error occured
func NewAssetsInfo(infos []*AssetInfo) (*AssetsInfo, error) {

	tmpl_str := Grandet.Asset("grandet.go.tmpl")
	tmpl := template.New("grandet.go.tmpl")
	tmpl = template.Must(tmpl.Parse(string(tmpl_str)))

	pkg_name := ""
	pkg_import := ""
	registers := []string{}
	for _, info := range infos {

		if len(pkg_name) <= 0 {
			pkg_name = info.AssetPackage
		} else if pkg_name != info.AssetPackage {
			return nil, errors.New("assets are not in the same package")
		}

		if len(pkg_import) <= 0 {
			pkg_import = info.PackageImport
		} else if pkg_import != info.PackageImport {
			return nil, errors.New("assets are not in the same import name")
		}

		registers = append(registers, info.AssetRegister)
	}

	return &AssetsInfo{
		tmpl:           tmpl,
		AssetPackage:   pkg_name,
		PackageImport:  pkg_import,
		AssetRegisters: registers,
	}, nil
}

// Stringer#String
func (ass *AssetsInfo) String() string {
	buff := &bytes.Buffer{}
	err := ass.tmpl.Execute(buff, ass)
	if err != nil {
		panic(err)
	}
	return buff.String()
}
