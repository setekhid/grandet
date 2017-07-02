package grandet

import (
	"path"
)

func pathFormatAndCheck(pkg_import string) string {

	pkg_import = path.Clean(pkg_import)

	if !path.IsAbs(pkg_import) {
		pkg_import = "/" + pkg_import
	}

	return pkg_import
}
