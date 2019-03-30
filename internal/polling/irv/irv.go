// Package irv implements an instant-runoff voting system (also known as
// ranked choice voting or the alternative vote).
package irv

import (
	"github.com/zikaeroh/strawrank/internal/polling"
	"github.com/zikaeroh/strawrank/internal/polling/fptp"
)

// Tally tallies ballots using an instant-runoff voting system. If there is
// a tie, then nobody is declared a winner.
func Tally(ballots []polling.Ballot) polling.Result {
	r, _ := TallyWithRounds(ballots)
	return r
}

// Tally tallies ballots the same way as Tally, but also shows the intermediate
// rounds leading to the final result.
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

		for i := end; i >= 0; i-- {
			can := result.Ranking[i]
			if can.Count > bottomCount {
				break
			}

			id := can.ID
			can.Count = 0

			for j, b := range ballots {
				if len(b.Votes) == 0 {
					continue
				}

				newB := b.Without(id)
				ballots[j] = newB

				if len(newB.Votes) == 0 {
					can.Count++
				}
			}

			eliminated = append(eliminated, can)
		}
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
