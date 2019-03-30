// Package polling defines types for different ballot measuring systems.
package polling

// Ballot represents a single poll ballot.
type Ballot struct {
	// Votes is an ordered list of candidate IDs, from most preferable to
	// least preferable.
	Votes []int
}

// NewBallot creates a new Ballot from a list of candidate IDs. It is shorthand
// for filling in Votes with a slice of IDs.
func NewBallot(ids ...int) Ballot {
	return Ballot{
		Votes: ids,
	}
}

// Without returns a copy of the ballot with all votes for a specific candidate
// removed.
func (b Ballot) Without(id int) Ballot {
	votes := make([]int, 0, len(b.Votes))

	for _, v := range b.Votes {
		if v != id {
			votes = append(votes, v)
		}
	}

	return Ballot{Votes: votes}
}

// Result holds the details of a tally result.
type Result struct {
	// Winners is a list of winning candidates. More than one implies a tie.
	Winners []Candidate

	// Ranking ranks all candidates from most popular to least popular.
	Ranking []Candidate

	// Total counts the total number of votes used in this result.
	Total int
}

// Candidate is a polling candidate for the tally result.
type Candidate struct {
	// ID is a unique number identifying a candidate.
	ID int

	// Count is the number of votes a candidate has.
	Count int
}
