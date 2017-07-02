#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

GRANDET_ROOT=$(dirname "${BASH_SOURCE[@]}")

go test ./...

( \
	go get github.com/setekhid/grandet/grandet && \
	grandet \
		-i github.com/setekhid/grandet/assets \
		-d "${GRANDET_ROOT}/assets" \
		"${GRANDET_ROOT}/assets/asset.go.tmpl" \
		"${GRANDET_ROOT}/assets/grandet.go.tmpl" \
)

git status assets/
