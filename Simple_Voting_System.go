package main

import (
	"fmt"
	"math/rand/v2"
)

type Candidate struct {
	Name  string
	Votes int
}

type VotingSystem struct {
	Candidates []Candidate
}

func (self *VotingSystem) addCandicate(name string) {
	if name == "" {
		fmt.Println("failed to add candicate")
	}
	candidate := Candidate{Name: name, Votes: 0}
	self.Candidates = append(self.Candidates, candidate)
}

func (self *VotingSystem) Vote(name string) {
	for i := range self.Candidates {
		if self.Candidates[i].Name == name {
			self.Candidates[i].Votes += 1
		}
	}
}

func (self *VotingSystem) GetResults() map[string]int {
	results := make(map[string]int)
	for i := range self.Candidates {
		results[self.Candidates[i].Name] = self.Candidates[i].Votes
	}
	return results
}

func main() {
	var votingSystem VotingSystem
	names := []string{"Lara", "Lisa", "David"}

	for _, item := range names {
		votingSystem.addCandicate(item)
	}

	totalVotes := rand.IntN(100)

	for i := 0; i < totalVotes; i++ {
		randName := rand.IntN(3)
		votingSystem.Vote(names[randName])
	}

	resultsMap := votingSystem.GetResults()
	fmt.Println("Total Votes : ", totalVotes)
	for key, value := range resultsMap {
		fmt.Printf("%v (%v votes)\n", key, value)
	}

}
