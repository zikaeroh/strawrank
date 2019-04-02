package templates

//go:generate sh -c "gobin -m -run github.com/valyala/quicktemplate/qtc && sed -i 's://line:// line:g' *.qtpl.go && gofmt -s -w *.qtpl.go"
