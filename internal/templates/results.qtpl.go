// Code generated by qtc from "results.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package templates

import "strconv"

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

type ResultsPage struct {
	BasePage
	Question string
	Choices  []string

	IRV       ResultView
	IRVRounds []ResultView
	FPTP      ResultView

	id int
}

type ResultView struct {
	Rows      []ResultRow
	ChartData []byte
}

type ResultRow struct {
	Name  string
	Count int
}

func (p *ResultsPage) StreamPageTitle(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
	StrawRank - Results - `)
	qw422016.E().S(p.Question)
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

func (p *ResultsPage) StreamPageMeta(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.8.0/Chart.min.css" integrity="sha256-aa0xaJgmK/X74WM224KMQeNQC2xYKwlAt08oZqjeF0E=" crossorigin="anonymous" />
    <script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.8.0/Chart.min.js" integrity="sha256-Uv9BNBucvCPipKQ2NS9wYpJmi8DTOEfTA/nH2aoJALw=" crossorigin="anonymous"></script>
`)
}

func (p *ResultsPage) WritePageMeta(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageMeta(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *ResultsPage) PageMeta() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageMeta(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *ResultsPage) StreamPageBody(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
    <div class="px-3 py-3 pt-md-5 pb-md-4 mx-auto text-center">
        <h1>`)
	qw422016.E().S(p.Question)
	qw422016.N().S(`</h1>
    </div>

    <div class="row">
        <div class="col">
            <h3 class="mb-4">"Instant-runoff" results <small class="float-right"><a href="https://en.wikipedia.org/wiki/Instant-runoff_voting" target="_blank">(About)</a></small></h3>
            `)
	p.StreamResult(qw422016, p.IRV)
	qw422016.N().S(`

            <div class="card" style="border: 0px; border-radius: 0px">
                <p class="card-header" style="background-color: #222">
                    <a class="card-link" data-toggle="collapse" href="#irv-rounds">
                        Show / hide IRV rounds
                    </a>
                </p>
                <div id="irv-rounds" class="collapse">
                    <div class="card-body">
                        `)
	for i, result := range p.IRVRounds {
		qw422016.N().S(`
                            <h4>Round `)
		qw422016.N().D(i + 1)
		qw422016.N().S(`</h4>
                            `)
		p.StreamResult(qw422016, result)
		qw422016.N().S(`
                        `)
	}
	qw422016.N().S(`
                    </div>
                </div>
            </div>
        </div>
    </div>

    <div class="row pt-md-5">
        <div class="col">
            <h3 class="mb-4">"First-past-the-post" results <small class="float-right"><a href="https://en.wikipedia.org/wiki/First-past-the-post_voting" target="_blank">(About)</a></small></h3>
            `)
	p.StreamResult(qw422016, p.FPTP)
	qw422016.N().S(`
        </div>
    </div>

    <p>
        Note: this page will not yet live update.
    </p>
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

func (p *ResultsPage) StreamResult(qw422016 *qt422016.Writer, result ResultView) {
	qw422016.N().S(`
    <div class="row">
        <div class="col">
            <table class="table table-striped table-bordered text-center">
                <thead>
                    <tr>
                        <th scope="col">#</th>
                        <th scope="col">Choice</th>
                        <th scope="col">Votes</th>
                    </tr>
                </thead>
                <tbody>
                    `)
	for i, c := range result.Rows {
		qw422016.N().S(`
                        <tr>
                            <th scope="row">`)
		qw422016.N().D(i + 1)
		qw422016.N().S(`</th>
                            <td>`)
		qw422016.E().S(c.Name)
		qw422016.N().S(`</td>
                            <td>`)
		qw422016.N().D(c.Count)
		qw422016.N().S(`</td>
                        </tr>
                    `)
	}
	qw422016.N().S(`
                </tbody>
            </table>
        </div>

        <div class="col">
            `)
	p.StreamChart(qw422016, result.ChartData)
	qw422016.N().S(`
        </div>
    </div>
`)
}

func (p *ResultsPage) WriteResult(qq422016 qtio422016.Writer, result ResultView) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamResult(qw422016, result)
	qt422016.ReleaseWriter(qw422016)
}

func (p *ResultsPage) Result(result ResultView) string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WriteResult(qb422016, result)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *ResultsPage) StreamChart(qw422016 *qt422016.Writer, data []byte) {
	qw422016.N().S(`
`)
	id := p.nextID()

	qw422016.N().S(`    

    <canvas id="`)
	qw422016.E().S(id)
	qw422016.N().S(`"></canvas>

    <script>
        new Chart(document.getElementById("`)
	qw422016.E().S(id)
	qw422016.N().S(`"), {
            type: "horizontalBar",
            data: JSON.parse("`)
	qw422016.N().JZ(data)
	qw422016.N().S(`"),
            options: {
                legend: { display: false },
                scales: {
                    xAxes: [{
                        ticks: {
                            beginAtZero: true,
                            display: false
                        }
                    }]
                }
            }
        })
    </script>

`)
}

func (p *ResultsPage) WriteChart(qq422016 qtio422016.Writer, data []byte) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamChart(qw422016, data)
	qt422016.ReleaseWriter(qw422016)
}

func (p *ResultsPage) Chart(data []byte) string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WriteChart(qb422016, data)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *ResultsPage) nextID() string {
	p.id++
	return "element-" + strconv.Itoa(p.id)
}
