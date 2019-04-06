// Code generated by qtc from "index.qtpl"; DO NOT EDIT.
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

type IndexPage struct {
	BasePage
	CSRF string
}

func (p *IndexPage) StreamPageTitle(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
	StrawRank
`)
}

func (p *IndexPage) WritePageTitle(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageTitle(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *IndexPage) PageTitle() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageTitle(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *IndexPage) StreamPageBody(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
    <div class="px-3 py-3 pt-md-5 pb-md-4 mx-auto text-center">
        <h1 class="mb-3">Create your own poll!</h1>

        <p>
            Type your question and choices below.
            Voters will be able to choose and rank them as they prefer.
        </p>

        <p>
            Order does not matter; choices will be randomized on the voting page.
        </p>
    </div>

    <div class="row justify-content-center">
        <div class="col-6">
            <form method="POST" autocomplete="off">
                <div class="form-group input-group-lg">
                    <input name="question" class="form-control" type="text" placeholder="Type your question here" required>
                </div>
                
                <br/>
                <!-- Replace with padding -->
                
                <div id="choice-inputs">
                    `)
	p.StreamChoice(qw422016)
	qw422016.N().S(`
                    `)
	p.StreamChoice(qw422016)
	qw422016.N().S(`
                    `)
	p.StreamChoice(qw422016)
	qw422016.N().S(`
                </div>

                <div class="form-group d-flex">
                    <button class="btn btn-primary mr-auto" type="submit">Submit</button>
                    <button class="btn btn-secondary ml-auto" onclick="addAnother(); return false">Add another</button>
                </div>

                `)
	qw422016.N().S(p.CSRF)
	qw422016.N().S(`
            </form>
        </div>
    </div>
`)
}

func (p *IndexPage) WritePageBody(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageBody(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *IndexPage) PageBody() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageBody(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *IndexPage) StreamPageScripts(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
    <script>
        var choiceInput = `)
	{
		qb422016 := qt422016.AcquireByteBuffer()
		p.WriteChoice(qb422016)
		qw422016.N().QZ(qb422016.B)
		qt422016.ReleaseByteBuffer(qb422016)
	}
	qw422016.N().S(`;

        function addAnother() {
            $("#choice-inputs").append(choiceInput);
        }
    </script>
`)
}

func (p *IndexPage) WritePageScripts(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageScripts(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *IndexPage) PageScripts() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageScripts(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *IndexPage) StreamChoice(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
`)
	qw422016.N().S(`<div class="form-group"><input name="choice" class="form-control" type="text" placeholder="Choice" required></div>`)
	qw422016.N().S(`
`)
}

func (p *IndexPage) WriteChoice(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamChoice(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *IndexPage) Choice() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WriteChoice(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
