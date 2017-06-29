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
func init() {
	// regist all local assets
	(&grandetAssets{}).registAssets()
}
```

```golang
import "github.com/setekhid/grandet/assets"

content := assets.Grandet.Asset(
	"github.com/setekhid/grandet/assets/asset.go.tmpl")
_ = content
```
