// Code generated by qtc from "base.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package templates

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

type Page interface {
	PageTitle() string
	StreamPageTitle(qw422016 *qt422016.Writer)
	WritePageTitle(qq422016 qtio422016.Writer)
	PageBody() string
	StreamPageBody(qw422016 *qt422016.Writer)
	WritePageBody(qq422016 qtio422016.Writer)
	PageMeta() string
	StreamPageMeta(qw422016 *qt422016.Writer)
	WritePageMeta(qq422016 qtio422016.Writer)
	PageScripts() string
	StreamPageScripts(qw422016 *qt422016.Writer)
	WritePageScripts(qq422016 qtio422016.Writer)
}

func StreamPageTemplate(qw422016 *qt422016.Writer, p Page) {
	qw422016.N().S(`
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
        <title>`)
	p.StreamPageTitle(qw422016)
	qw422016.N().S(`</title>
        <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bootswatch/4.3.1/darkly/bootstrap.min.css" integrity="sha256-6W1mxPaAt4a6pkJVW5x5Xmq/LvxuQpR9dlzgy77SeZs=" crossorigin="anonymous" />
        
        `)
	p.StreamPageMeta(qw422016)
	qw422016.N().S(`
    </head>
    <body>
        <nav class="navbar navbar-expand-lg navbar-dark bg-primary sticky-top box-shadow">
            <div class="container">
                <a class="navbar-brand" href="/">StrawRank</a>
                `)
	if _, ok := p.(*IndexPage); !ok {
		qw422016.N().S(`
                <div class="nav-item">
                    <a class="nav-link active text-white" href="/">Create a new poll</a>
                </div>
                `)
	}
	qw422016.N().S(`
                <div class="nav-item ml-auto">
                    <a class="nav-link text-white" href="/about">About</a>
                </div>
            </div>
        </nav>

        <div class="container">
            `)
	p.StreamPageBody(qw422016)
	qw422016.N().S(`
        </div>

        <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.3.1/jquery.slim.min.js" integrity="sha256-3edrmyuQ0w65f8gfBsqowzjJe2iM6n0nKciPUp8y+7E=" crossorigin="anonymous"></script>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js" integrity="sha256-ZvOgfh+ptkpoa2Y4HkRY28ir89u/+VRyDE7sB7hEEcI=" crossorigin="anonymous"></script>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.3.1/js/bootstrap.min.js" integrity="sha256-CjSoeELFOcH0/uxWu6mC/Vlrc1AARqbm/jiiImDGV3s=" crossorigin="anonymous"></script>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.8.1/js/all.min.js" integrity="sha256-HT9Zb3b1PVPvfLH/7/1veRtUvWObQuTyPn8tezb5HEg=" crossorigin="anonymous"></script>
        
        `)
	p.StreamPageScripts(qw422016)
	qw422016.N().S(`
    </body>
</html>
`)
}

func WritePageTemplate(qq422016 qtio422016.Writer, p Page) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamPageTemplate(qw422016, p)
	qt422016.ReleaseWriter(qw422016)
}

func PageTemplate(p Page) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WritePageTemplate(qb422016, p)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

type BasePage struct{}

func (p *BasePage) StreamPageTitle(qw422016 *qt422016.Writer) {
	qw422016.N().S(`StrawRank`)
}

func (p *BasePage) WritePageTitle(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageTitle(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *BasePage) PageTitle() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageTitle(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *BasePage) StreamPageBody(qw422016 *qt422016.Writer) {
}

func (p *BasePage) WritePageBody(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageBody(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *BasePage) PageBody() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageBody(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *BasePage) StreamPageMeta(qw422016 *qt422016.Writer) {
}

func (p *BasePage) WritePageMeta(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageMeta(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *BasePage) PageMeta() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageMeta(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *BasePage) StreamPageScripts(qw422016 *qt422016.Writer) {
}

func (p *BasePage) WritePageScripts(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageScripts(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *BasePage) PageScripts() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageScripts(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
