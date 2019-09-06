package polling_test

import (
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/zikaeroh/strawrank/internal/polling"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"
)

func TestWithout(t *testing.T) {
	tests := []struct {
		name  string
		input polling.Ballot
		id    int64
		want  polling.Ballot
	}{
		{
			name: "Empty",
		},
		{
			name:  "No change",
			input: polling.NewBallot(1, 2, 3),
			id:    0,
			want:  polling.NewBallot(1, 2, 3),
		},
		{
			name:  "Remove first",
			input: polling.NewBallot(1, 2, 3),
			id:    1,
			want:  polling.NewBallot(2, 3),
		},
		{
			name:  "Remove middle",
			input: polling.NewBallot(1, 2, 3),
			id:    2,
			want:  polling.NewBallot(1, 3),
		},
		{
			name:  "Remove end",
			input: polling.NewBallot(1, 2, 3),
			id:    3,
			want:  polling.NewBallot(1, 2),
		},
		{
			name:  "Remove many",
			input: polling.NewBallot(1, 2, 3, 2, 3, 1, 2),
			id:    3,
			want:  polling.NewBallot(1, 2, 2, 1, 2),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			orig := copyVotes(test.input.Votes)

			got := test.input.Without(test.id)

			assert.Check(t, cmp.DeepEqual(test.want, got, cmpopts.EquateEmpty()))
			assert.Check(t, cmp.DeepEqual(test.input.Votes, orig, cmpopts.EquateEmpty()))
		})
	}
}

func copyVotes(in []int64) []int64 {
	out := make([]int64, len(in))
	copy(out, in)
	return out
}
