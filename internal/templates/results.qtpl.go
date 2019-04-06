// Code generated by qtc from "results.qtpl"; DO NOT EDIT.
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

type ResultsPage struct {
	Name string
	BasePage
}

func (p *ResultsPage) StreamPageTitle(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
	StrawRank - Results - `)
	qw422016.E().S(p.Name)
	qw422016.N().S(`
`)
}

func (p *ResultsPage) WritePageTitle(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageTitle(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *ResultsPage) PageTitle() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageTitle(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *ResultsPage) StreamPageBody(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
    <div class="row">
        <div class="col">
            <h1>Wassap</h1>
        </div>
    </div>
`)
}

func (p *ResultsPage) WritePageBody(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageBody(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *ResultsPage) PageBody() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageBody(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
