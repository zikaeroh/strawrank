package irv_test

import (
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/zikaeroh/strawrank/internal/polling"
	"github.com/zikaeroh/strawrank/internal/polling/irv"
	"gotest.tools/assert"
	"gotest.tools/assert/cmp"
)

var tests = []struct {
	name    string
	ballots []polling.Ballot
	result  polling.Result
	rounds  []polling.Result
}{
	{
		name: "Empty",
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
		rounds: []polling.Result{
			{
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
			Ranking: []polling.Candidate{
				{ID: 1, Count: 2},
				{ID: 2, Count: 2},
			},
			Total: 4,
		},
		rounds: []polling.Result{
			{
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
			{},
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
		rounds: []polling.Result{
			{
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
	},
	{
		name: "Two dominant",
		ballots: []polling.Ballot{
			polling.NewBallot(1),
			polling.NewBallot(1),
			polling.NewBallot(1),
			polling.NewBallot(1),
			polling.NewBallot(1),
			polling.NewBallot(1),
			polling.NewBallot(2, 3),
			polling.NewBallot(2, 3),
			polling.NewBallot(2, 3),
			polling.NewBallot(2, 3),
			polling.NewBallot(2, 3),
			polling.NewBallot(3, 2),
			polling.NewBallot(3, 2),
			polling.NewBallot(3),
			polling.NewBallot(4),
			polling.NewBallot(5),
		},
		result: polling.Result{
			Winners: []polling.Candidate{
				{ID: 2, Count: 7},
			},
			Ranking: []polling.Candidate{
				{ID: 2, Count: 7},
				{ID: 1, Count: 6},
				{ID: 3, Count: 1},
				{ID: 4, Count: 1},
				{ID: 5, Count: 1},
			},
			Total: 16,
		},
		rounds: []polling.Result{
			{
				Winners: []polling.Candidate{
					{ID: 1, Count: 6},
				},
				Ranking: []polling.Candidate{
					{ID: 1, Count: 6},
					{ID: 2, Count: 5},
					{ID: 3, Count: 3},
					{ID: 4, Count: 1},
					{ID: 5, Count: 1},
				},
				Total: 16,
			},
			{
				Winners: []polling.Candidate{
					{ID: 1, Count: 6},
				},
				Ranking: []polling.Candidate{
					{ID: 1, Count: 6},
					{ID: 2, Count: 5},
					{ID: 3, Count: 3},
				},
				Total: 14,
			},
			{
				Winners: []polling.Candidate{
					{ID: 2, Count: 7},
				},
				Ranking: []polling.Candidate{
					{ID: 2, Count: 7},
					{ID: 1, Count: 6},
				},
				Total: 13,
			},
		},
	},
}

func TestTally(t *testing.T) {
	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			result := irv.Tally(test.ballots)

			assert.Check(t, cmp.DeepEqual(test.result, result, cmpopts.EquateEmpty()))
		})
	}
}

func TestTallyWithRounds(t *testing.T) {
	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			result, rounds := irv.TallyWithRounds(test.ballots)

			assert.Check(t, cmp.DeepEqual(test.result, result, cmpopts.EquateEmpty()))
			assert.Check(t, cmp.DeepEqual(test.rounds, rounds, cmpopts.EquateEmpty()))
		})
	}
}
