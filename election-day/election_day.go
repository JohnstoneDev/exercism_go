package electionday

import "fmt"

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	initialVotesPointer := &initialVotes

	return initialVotesPointer
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	}

	return *counter
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	*counter += increment
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	return &ElectionResult{Name: candidateName, Votes: votes}
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	name, votes := (*result).Name, (*result).Votes
	return fmt.Sprintf("%v (%v)", name, votes)
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	candidateVotes := results[candidate]
	candidateVotes -= 1

	results[candidate] = candidateVotes
}
