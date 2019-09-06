package fptp_test

import (
	"testing"

	"github.com/zikaeroh/strawrank/internal/polling"
	"github.com/zikaeroh/strawrank/internal/polling/fptp"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"
)

func TestTally(t *testing.T) {
	tests := []struct {
		name    string
		ballots []polling.Ballot
		result  polling.Result
	}{
		{
			name: "Empty",
		},
		{
			name: "Single empty",
			ballots: []polling.Ballot{
				polling.NewBallot(),
			},
		},
		{
			name: "Three ballots",
			ballots: []polling.Ballot{
				polling.NewBallot(1),
				polling.NewBallot(1),
				polling.NewBallot(2),
			},
			result: polling.Result{
				Winners: []polling.Candidate{
					{ID: 1, Count: 2},
				},
				Ranking: []polling.Candidate{
					{ID: 1, Count: 2},
					{ID: 2, Count: 1},
				},
				Total: 3,
			},
		},
		{
			name: "Four ballot tie",
			ballots: []polling.Ballot{
				polling.NewBallot(1),
				polling.NewBallot(1),
				polling.NewBallot(2),
				polling.NewBallot(2),
			},
			result: polling.Result{
				Winners: []polling.Candidate{
					{ID: 1, Count: 2},
					{ID: 2, Count: 2},
				},
				Ranking: []polling.Candidate{
					{ID: 1, Count: 2},
					{ID: 2, Count: 2},
				},
				Total: 4,
			},
		},
		{
			name: "Three ballots with spoiled",
			ballots: []polling.Ballot{
				polling.NewBallot(1),
				polling.NewBallot(1),
				polling.NewBallot(2),
				polling.NewBallot(),
				polling.NewBallot(),
			},
			result: polling.Result{
				Winners: []polling.Candidate{
					{ID: 1, Count: 2},
				},
				Ranking: []polling.Candidate{
					{ID: 1, Count: 2},
					{ID: 2, Count: 1},
				},
				Total: 3,
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			result := fptp.Tally(test.ballots)

			assert.Check(t, cmp.DeepEqual(test.result, result))
		})
	}
}
