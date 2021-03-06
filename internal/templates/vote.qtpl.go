// Code generated by qtc from "vote.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package templates

import "math/rand"

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

type VotePage struct {
	BasePage
	CSRF     string
	Path     string
	Question string
	Choices  []string
}

func (p *VotePage) StreamPageTitle(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
	StrawRank - Vote - `)
	qw422016.E().S(p.Question)
	qw422016.N().S(`
`)
}

func (p *VotePage) WritePageTitle(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageTitle(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *VotePage) PageTitle() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageTitle(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *VotePage) StreamPageMeta(qw422016 *qt422016.Writer) {
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
}

func (p *VotePage) WritePageMeta(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageMeta(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *VotePage) PageMeta() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageMeta(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *VotePage) StreamPageBody(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
    <div class="px-3 py-3 pt-md-5 pb-md-4 mx-auto text-center">
        <h1 class="mb-3">`)
	qw422016.E().S(p.Question)
	qw422016.N().S(`</h1>

        <h3 class="mb-3 font-italic">
            <a href="`)
	qw422016.E().S(p.Path)
	qw422016.N().S(`/r">Results</a>
        </h3>

        <p>To vote, drag your choices from the left to the right.</p>
        <p>Order by preference. Not all options need to be selected.</p>
    </div>

    <div class="row">
        <div class="col">
            <h4 class="text-center mb-4">Available choices</h4>

            <div id="vote-unchosen" class="vote-list list-group">
                `)
	type pair struct {
		i int
		v string
	}

	choices := make([]pair, len(p.Choices))

	for i, v := range p.Choices {
		choices[i] = pair{i: i, v: v}
	}

	rand.Shuffle(len(choices), func(i, j int) {
		choices[i], choices[j] = choices[j], choices[i]
	})

	qw422016.N().S(`
                `)
	for _, choice := range choices {
		qw422016.N().S(`
                <div class="list-group-item d-flex" data-index="`)
		qw422016.N().D(choice.i)
		qw422016.N().S(`">
                    <div class="mr-1 align-self-center">
                        <span class="badge badge-secondary"></span>
                    </div>
                    <span class="text flex-fill">`)
		qw422016.E().S(choice.v)
		qw422016.N().S(`</span>
                    <div class="ml-1 align-self-center">
                        <i class="fa fa-times close remove" onclick="removeVote(this)" data-toggle="tooltip" data-placement="right" title="Remove"></i>
                    </div>
                </div>
                `)
	}
	qw422016.N().S(`
            </div>
        </div>

        <div class="col">
            <h4 class="text-center mb-4">Your choices</h4>

            <div id="vote-chosen" class="vote-list list-group"></div>

            <form id="vote-form" class="mt-4" method="POST">
                <button type="submit" class="btn btn-primary btn-block" id="submit-button" disabled>Submit</button>
                <input type="hidden" id="votes" name="votes" value="" />
                `)
	qw422016.N().S(p.CSRF)
	qw422016.N().S(`
            </form>
        </div>
    </div>
`)
}

func (p *VotePage) WritePageBody(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageBody(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *VotePage) PageBody() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageBody(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *VotePage) StreamPageScripts(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
    <script src="https://cdnjs.cloudflare.com/ajax/libs/Sortable/1.8.4/Sortable.min.js" integrity="sha256-yEySJXdfoPg1V6xPh7TjRM0MRZnJCnIxsoBEp50u0as=" crossorigin="anonymous"></script>

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
            var tooltip = $(e).attr("aria-describedby");
            $(e).parent().parent().appendTo("#vote-unchosen");
            $("#"+tooltip).remove();
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

            $(document).ready(function() {
                $('[data-toggle="tooltip"]').tooltip();
            });
        });

        new Sortable($("#vote-unchosen")[0], {
            group: 'votes',
            animation: 150,
            chosenClass: 'list-group-item-light',
            onChange: onChange,
        });

        new Sortable($("#vote-chosen")[0], {
            group: 'votes',
            animation: 150,
            chosenClass: 'list-group-item-light',
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
}

func (p *VotePage) WritePageScripts(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageScripts(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *VotePage) PageScripts() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageScripts(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
