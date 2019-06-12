// +build tools

package tools

import (
	_ "github.com/golang-migrate/migrate/v4/cmd/migrate"
	_ "github.com/mjibson/esc"
	_ "github.com/valyala/quicktemplate/qtc"
	_ "github.com/volatiletech/sqlboiler"
	_ "github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql"
)
