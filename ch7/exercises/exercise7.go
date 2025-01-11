package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"sort"
)

type Ranker interface {
	Ranking() []string
}

type Team struct {
	Name        string
	PlayerNames []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l *League) MatchResult(
	firstTeam Team,
	firstTeamScore int,
	secondTeam Team,
	secondTeamScore int) {
	if firstTeamScore > secondTeamScore {
		l.Wins[firstTeam.Name]++
	} else if secondTeamScore > firstTeamScore {
		l.Wins[secondTeam.Name]++
	}
}

func (l *League) Ranking() []string {
	// Convert the map to a slice of team names
	rankings := make([]string, 0, len(l.Wins))
	for team := range l.Wins {
		rankings = append(rankings, team)
	}

	// Sort the team names by their wins in descending order
	sort.Slice(rankings, func(i, j int) bool {
		return l.Wins[rankings[i]] > l.Wins[rankings[j]]
	})

	return rankings
}

func RankPrinter(r Ranker, w io.Writer) {
	rankings := r.Ranking()
	for i, team := range rankings {
		line := fmt.Sprintf("%d. %s\n", i+1, team)
		_, err := io.WriteString(w, line)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	// Teams
	team1 := Team{Name: "Eagles", PlayerNames: []string{"Alice", "Bob", "Charlie"}}
	team2 := Team{Name: "Tigers", PlayerNames: []string{"David", "Ella", "Frank"}}
	team3 := Team{Name: "Lions", PlayerNames: []string{"George", "Hannah", "Ian"}}
	team4 := Team{Name: "Panthers", PlayerNames: []string{"Jack", "Karen", "Liam"}}
	team5 := Team{Name: "Bulls", PlayerNames: []string{"Mia", "Noah", "Olivia"}}
	team6 := Team{Name: "Wolves", PlayerNames: []string{"Parker", "Quinn", "Riley"}}
	team7 := Team{Name: "Sharks", PlayerNames: []string{"Sophia", "Thomas", "Uma"}}
	team8 := Team{Name: "Falcons", PlayerNames: []string{"Victor", "Wendy", "Xander"}}

	// Create a league
	league := League{
		Teams: []Team{team1, team2, team3, team4, team5, team6, team7, team8},
		Wins:  make(map[string]int),
	}

	// Simulate matches with random scores from 0-100
	league.MatchResult(team1, rand.Intn(101), team2, rand.Intn(101))
	league.MatchResult(team3, rand.Intn(101), team4, rand.Intn(101))
	league.MatchResult(team5, rand.Intn(101), team6, rand.Intn(101))
	league.MatchResult(team7, rand.Intn(101), team8, rand.Intn(101))
	league.MatchResult(team1, rand.Intn(101), team3, rand.Intn(101))
	league.MatchResult(team2, rand.Intn(101), team4, rand.Intn(101))
	league.MatchResult(team5, rand.Intn(101), team7, rand.Intn(101))
	league.MatchResult(team6, rand.Intn(101), team8, rand.Intn(101))
	league.MatchResult(team4, rand.Intn(101), team7, rand.Intn(101))
	league.MatchResult(team8, rand.Intn(101), team1, rand.Intn(101))

	// Print the team rankings
	//rankings := league.Ranking()
	fmt.Println("Team Rankings:")
	//for i, team := range rankings {
	//	fmt.Printf("%d. %s (%d wins)\n", i+1, team, league.Wins[team])
	//}

	RankPrinter(&league, os.Stdout)
}
