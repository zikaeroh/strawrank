// +build tools

package tools

import (
	_ "github.com/mjibson/esc"
	_ "github.com/valyala/quicktemplate/qtc"
	// Not included here for now.
	// sqlboiler expects partner binaries for drivers which wouldn't be found with gobin.
	// _ "github.com/volatiletech/sqlboiler"
)
