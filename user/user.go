package user

import (
	"fmt"
	"koodWordle/game"
	"koodWordle/input"
)

func UserName() string {
	fmt.Println("Enter your username:")
	name := input.Input()

	return name
}

func CreateUser(name string) *User {
	return &User{
		Name: name,
		Stats: GameStats{
			GamesPlayed:   0,
			GamesWon:      0,
			TotalAttempts: 0,
		},
	}
}

func GreetUser() {
	fmt.Println("Welcome to Wordle! Guess the 5-letter word.")
}

type User struct {
	Name  string
	Stats GameStats
}

type GameStats struct {
	GamesPlayed   int
	GamesWon      int
	TotalAttempts int
}

func Stats(s *User) {
	fmt.Println("Do you want to see your stats? (yes/no):")
	answer := input.Input()
	for {
		if answer == "yes" {
			fmt.Sprintf("Stats for: %s", s.Name)
			fmt.Sprintf("Games played: %d: ", s.Stats.GamesPlayed)
			fmt.Sprintf("Games won: %d", s.Stats.GamesWon)
			fmt.Sprintf("Average attempts per game: %f", float64(s.Stats.TotalAttempts)/float64(s.Stats.GamesWon))
		} else if answer == "no" {
			game.Exit()
		} else {
			fmt.Sprintf("Try again!")
		}
	}
}

func UpdateUserStats(s *User, won bool, attempts int) {
	s.Stats.GamesPlayed++
	if won {
		s.Stats.GamesWon++
	}
	s.Stats.TotalAttempts += attempts
}
