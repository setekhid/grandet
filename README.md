# grandet

[![GoDoc](https://godoc.org/github.com/setekhid/grandet?status.svg)](https://godoc.org/github.com/setekhid/grandet) [![Go Report Card](https://goreportcard.com/badge/github.com/setekhid/grandet)](https://goreportcard.com/report/github.com/setekhid/grandet)

another golang assets handler, I made this just because the other solutions are
already dead.

# usage

```shell
go get github.com/setekhid/grandet/grandet
grandet \
	-i github.com/setekhid/grandet/assets \
	-d $GOPATH/src/github.com/setekhid/grandet/assets \
	$GOPATH/src/github.com/setekhid/grandet/assets/asset.go.tmpl \
	$GOPATH/src/github.com/setekhid/grandet/assets/grandet.go.tmpl
# gofmt as you like
```

```golang
package assets

func init() {
	// regist all local assets
	assets := &grandetAssets{}
	assets.registAssets()
	assets.BelongTo("github.com/setekhid/grandet")
}
```

```golang
package main

import "github.com/setekhid/grandet/assets"

func main() {
	content := assets.Grandet.Asset("asset.go.tmpl")
	_ = content
}
```

# license

this project is under bsd 3-clause license, see `LICENSE` file
