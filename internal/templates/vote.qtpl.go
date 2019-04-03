// This file is automatically generated by qtc from "vote.qtpl".
// See https://github.com/valyala/quicktemplate for details.

// line vote.qtpl:1
package templates

// line vote.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

// line vote.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

// line vote.qtpl:2
type VotePage struct {
	BasePage
	CSRF    string
	Name    string
	Choices []string
}

// line vote.qtpl:10
func (p *VotePage) StreamPageTitle(qw422016 *qt422016.Writer) {
	// line vote.qtpl:10
	qw422016.N().S(`
	StrawRank - `)
	// line vote.qtpl:11
	qw422016.E().S(p.Name)
	// line vote.qtpl:11
	qw422016.N().S(`
`)
	// line vote.qtpl:12
}

// line vote.qtpl:12
func (p *VotePage) WritePageTitle(qq422016 qtio422016.Writer) {
	// line vote.qtpl:12
	qw422016 := qt422016.AcquireWriter(qq422016)
	// line vote.qtpl:12
	p.StreamPageTitle(qw422016)
	// line vote.qtpl:12
	qt422016.ReleaseWriter(qw422016)
	// line vote.qtpl:12
}

// line vote.qtpl:12
func (p *VotePage) PageTitle() string {
	// line vote.qtpl:12
	qb422016 := qt422016.AcquireByteBuffer()
	// line vote.qtpl:12
	p.WritePageTitle(qb422016)
	// line vote.qtpl:12
	qs422016 := string(qb422016.B)
	// line vote.qtpl:12
	qt422016.ReleaseByteBuffer(qb422016)
	// line vote.qtpl:12
	return qs422016
	// line vote.qtpl:12
}

// line vote.qtpl:14
func (p *VotePage) StreamPageMeta(qw422016 *qt422016.Writer) {
	// line vote.qtpl:14
	qw422016.N().S(`
	<style>
        .vote-list:empty {
            padding: 12px 20px;
            border: 1px solid transparent;
            border-radius: 0.25rem;
            border-color: #444;
        }

        .vote-list:empty::after {
            content: "Drag here";
            text-align: center;
            font-style: italic;
        }

        .badge {
            margin-right: 1em;
        }

        .remove {
            margin-left: 1em;
        }

        #vote-unchosen .list-group-item .badge {
            visibility: hidden;
        }

        #vote-unchosen .list-group-item .remove {
            visibility: hidden;
        }
    </style>
`)
	// line vote.qtpl:45
}

// line vote.qtpl:45
func (p *VotePage) WritePageMeta(qq422016 qtio422016.Writer) {
	// line vote.qtpl:45
	qw422016 := qt422016.AcquireWriter(qq422016)
	// line vote.qtpl:45
	p.StreamPageMeta(qw422016)
	// line vote.qtpl:45
	qt422016.ReleaseWriter(qw422016)
	// line vote.qtpl:45
}

// line vote.qtpl:45
func (p *VotePage) PageMeta() string {
	// line vote.qtpl:45
	qb422016 := qt422016.AcquireByteBuffer()
	// line vote.qtpl:45
	p.WritePageMeta(qb422016)
	// line vote.qtpl:45
	qs422016 := string(qb422016.B)
	// line vote.qtpl:45
	qt422016.ReleaseByteBuffer(qb422016)
	// line vote.qtpl:45
	return qs422016
	// line vote.qtpl:45
}

// line vote.qtpl:47
func (p *VotePage) StreamPageBody(qw422016 *qt422016.Writer) {
	// line vote.qtpl:47
	qw422016.N().S(`
    <div>
        <h1 style="text-align: center; padding: 3rem 1.5rem;">`)
	// line vote.qtpl:49
	qw422016.E().S(p.Name)
	// line vote.qtpl:49
	qw422016.N().S(`</h1>

        <p>To vote, drag your choices from the left to the right. Order by preference.</p>
    </div>

    <br/>

    <div class="row">
        <div class="col">
            <h4 style="text-align: center">Available choices</h4>
            <br/>
            <div id="vote-unchosen" class="vote-list list-group">
                `)
	// line vote.qtpl:61
	for i, choice := range p.Choices {
		// line vote.qtpl:61
		qw422016.N().S(`
                <div class="list-group-item d-flex" data-index="`)
		// line vote.qtpl:62
		qw422016.N().D(i)
		// line vote.qtpl:62
		qw422016.N().S(`">
                    <div class="mr-1 align-self-center">
                        <span class="badge badge-secondary"></span>
                    </div>
                    <span class="text flex-fill">`)
		// line vote.qtpl:66
		qw422016.E().S(choice)
		// line vote.qtpl:66
		qw422016.N().S(`</span>
                    <div class="ml-1 align-self-center">
                        <i class="fa fa-times close remove" onclick="removeVote(this)"></i>
                    </div>
                </div>
                `)
		// line vote.qtpl:71
	}
	// line vote.qtpl:71
	qw422016.N().S(`
            </div>
        </div>

        <div class="col">
            <h4 style="text-align: center">Your choices</h4>
            <br/>
            <div id="vote-chosen" class="vote-list list-group"></div>

            <br/>

            <form id="vote-form" method="POST">
                <button type="submit" class="btn btn-primary btn-block" id="submit-button" disabled>Submit</button>
                <input type="hidden" id="votes" name="votes" value="" />
                `)
	// line vote.qtpl:85
	qw422016.N().S(p.CSRF)
	// line vote.qtpl:85
	qw422016.N().S(`
            </form>
        </div>
    </div>
`)
	// line vote.qtpl:89
}

// line vote.qtpl:89
func (p *VotePage) WritePageBody(qq422016 qtio422016.Writer) {
	// line vote.qtpl:89
	qw422016 := qt422016.AcquireWriter(qq422016)
	// line vote.qtpl:89
	p.StreamPageBody(qw422016)
	// line vote.qtpl:89
	qt422016.ReleaseWriter(qw422016)
	// line vote.qtpl:89
}

// line vote.qtpl:89
func (p *VotePage) PageBody() string {
	// line vote.qtpl:89
	qb422016 := qt422016.AcquireByteBuffer()
	// line vote.qtpl:89
	p.WritePageBody(qb422016)
	// line vote.qtpl:89
	qs422016 := string(qb422016.B)
	// line vote.qtpl:89
	qt422016.ReleaseByteBuffer(qb422016)
	// line vote.qtpl:89
	return qs422016
	// line vote.qtpl:89
}

// line vote.qtpl:91
func (p *VotePage) StreamPageScripts(qw422016 *qt422016.Writer) {
	// line vote.qtpl:91
	qw422016.N().S(`
    <script src="https://cdn.jsdelivr.net/npm/sortablejs@latest/Sortable.min.js"></script>

    <script>
        var count = 0;
        var updateSubmit = function() {
            $("#submit-button").attr("disabled", count == 0);
        };
        
        var onChange = function() {
            $("#vote-chosen div.list-group-item span.badge").each(function(i, e) {
                $(e).text(i+1);
            })
            $("#vote-unchosen div.list-group-item span.badge").each(function(i, e) {
                $(e).empty();
            })
        };

        function removeVote(e) {
            $(e).parent().parent().appendTo("#vote-unchosen");
            count--;
            updateSubmit();
            onChange();
        }

        $(document).ready(function() {
            $("#vote-form").submit(function() {
                var votes = [];

                $("#vote-chosen div.list-group-item").each(function(i, e) {
                    var index = Number(e.dataset.index);
                    if (index !== NaN) {
                        votes.push(index);
                    }
                });

                $("#votes").val(JSON.stringify(votes));
                return true;
            })
        });

        new Sortable($("#vote-unchosen")[0], {
            group: 'votes',
            animation: 150,
            onChange: onChange,
        });

        new Sortable($("#vote-chosen")[0], {
            group: 'votes',
            animation: 150,
            onAdd: function(evt) {
                count++;
                updateSubmit();
            },
            onRemove: function(evt) {
                count--;
                updateSubmit();
            },
            onChange: onChange,
        });
    </script>
`)
	// line vote.qtpl:152
}

// line vote.qtpl:152
func (p *VotePage) WritePageScripts(qq422016 qtio422016.Writer) {
	// line vote.qtpl:152
	qw422016 := qt422016.AcquireWriter(qq422016)
	// line vote.qtpl:152
	p.StreamPageScripts(qw422016)
	// line vote.qtpl:152
	qt422016.ReleaseWriter(qw422016)
	// line vote.qtpl:152
}

// line vote.qtpl:152
func (p *VotePage) PageScripts() string {
	// line vote.qtpl:152
	qb422016 := qt422016.AcquireByteBuffer()
	// line vote.qtpl:152
	p.WritePageScripts(qb422016)
	// line vote.qtpl:152
	qs422016 := string(qb422016.B)
	// line vote.qtpl:152
	qt422016.ReleaseByteBuffer(qb422016)
	// line vote.qtpl:152
	return qs422016
	// line vote.qtpl:152
}
