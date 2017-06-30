# grandet

another golang assets generator

```shell
go get github.com/setekhid/grandet/grandet
grandet \
	-i github.com/setekhid/grandet/assets \
	-d $GOPATH/src/github.com/setekhid/grandet/assets \
	asset.go.tmpl grandet.go.tmpl
# gofmt as you like
```

```golang
package assets

func init() {
	// regist all local assets
	(&grandetAssets{}).registAssets()
}
```

```golang
import "github.com/setekhid/grandet/assets"

func TasteAssets() {
	content := assets.Grandet.Asset("asset.go.tmpl")
	_ = content
}
```
