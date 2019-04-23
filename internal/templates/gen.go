package templates

//go:generate gobin -m -run github.com/valyala/quicktemplate/qtc
//go:generate sh -c "set -x; sed -i '/.qtpl:/d' *.qtpl.go"
//go:generate sh -c "set -x; gofmt -s -w *.qtpl.go"
