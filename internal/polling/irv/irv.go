// Package irv implements an instant-runoff voting system (also known as
// ranked choice voting or the alternative vote).
package irv

import (
	"sort"

	"github.com/zikaeroh/strawrank/internal/polling"
	"github.com/zikaeroh/strawrank/internal/polling/fptp"
)

// Tally tallies ballots using an instant-runoff voting system. If there is
// a tie, then nobody is declared a winner.
func Tally(ballots []polling.Ballot) polling.Result {
	r, _ := TallyWithRounds(ballots)
	return r
}

// TallyWithRounds tallies ballots the same way as Tally, but also shows the
// intermediate rounds leading to the final result.
func TallyWithRounds(ballots []polling.Ballot) (polling.Result, []polling.Result) {
	if len(ballots) == 0 {
		return polling.Result{}, nil
	}

	ballots = copyBallots(ballots)

	var eliminated []polling.Candidate
	var rounds []polling.Result

	for {
		result := fptp.Tally(ballots)
		rounds = append(rounds, result)

		majority := result.Total/2 + 1

		if len(result.Ranking) == 0 || result.Winners[0].Count >= majority {
			return withEliminated(result, eliminated), rounds
		}

		end := len(result.Ranking) - 1
		bottomCount := result.Ranking[end].Count

		toRemove := make(map[int64]bool)

		for i := end; i >= 0; i-- {
			can := result.Ranking[i]
			if can.Count > bottomCount {
				break
			}

			toRemove[can.ID] = true
		}

		toRemoveCounts := make(map[int64]int)

		for i, b := range ballots {
			ok := true

			for _, v := range b.Votes {
				if toRemove[v] {
					ok = false
					break
				}
			}

			if ok {
				continue
			}

			votes := make([]int64, 0, len(b.Votes))

			first := true
			var firstRemoved int64

			for _, v := range b.Votes {
				if toRemove[v] {
					if first {
						first = false
						firstRemoved = v
					}
				} else {
					votes = append(votes, v)
				}
			}

			if len(votes) == 0 {
				toRemoveCounts[firstRemoved]++
			}

			ballots[i].Votes = votes
		}

		newElim := make([]polling.Candidate, 0, len(toRemoveCounts))

		for id, count := range toRemoveCounts {
			newElim = append(newElim, polling.Candidate{
				ID:    id,
				Count: count,
			})
		}

		sort.Slice(newElim, func(i, j int) bool {
			return newElim[i].ID > newElim[j].ID
		})

		eliminated = append(eliminated, newElim...)
	}
}

func copyBallots(ballots []polling.Ballot) []polling.Ballot {
	tmp := make([]polling.Ballot, len(ballots))
	copy(tmp, ballots)
	return tmp
}

func withEliminated(result polling.Result, eliminated []polling.Candidate) polling.Result {
	for i := len(eliminated)/2 - 1; i >= 0; i-- {
		opp := len(eliminated) - 1 - i
		eliminated[i], eliminated[opp] = eliminated[opp], eliminated[i]
	}

	total := result.Total

	for _, r := range eliminated {
		total += r.Count
	}

	return polling.Result{
		Winners: result.Winners,
		Ranking: append(result.Ranking, eliminated...),
		Total:   total,
	}
}
