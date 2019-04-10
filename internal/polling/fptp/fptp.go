// Package fptp implements a first-past-the-post electoral system.
package fptp

import (
	"sort"

	"github.com/zikaeroh/strawrank/internal/polling"
)

// Tally tallies ballots into a single first-past-the-post result. If there is
// a tie for first, both are returned as winners.
func Tally(ballots []polling.Ballot) polling.Result {
	if len(ballots) == 0 {
		return polling.Result{}
	}

	tally := make(map[int64]int)
	total := 0

	for _, b := range ballots {
		for _, v := range b.Votes {
			tally[v]++
			total++
			break
		}
	}

	if len(tally) == 0 {
		return polling.Result{}
	}

	candidates := make([]polling.Candidate, 0, len(tally))

	for id, count := range tally {
		candidates = append(candidates, polling.Candidate{ID: id, Count: count})
	}

	// Undo map randomness in ID ordering.
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].ID < candidates[j].ID
	})

	sort.SliceStable(candidates, func(i, j int) bool {
		return candidates[i].Count > candidates[j].Count
	})

	winners := candidates

	for i := 1; i < len(winners); i++ {
		count := winners[i].Count
		prev := winners[i-1].Count

		if prev != count {
			winners = winners[:i]
			break
		}
	}

	return polling.Result{
		Winners: winners,
		Ranking: candidates,
		Total:   total,
	}
}
