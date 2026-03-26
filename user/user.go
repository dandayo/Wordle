package user

import (
	"fmt"
	"koodWordle/game"
	"koodWordle/input"
)

func UserName() {
	fmt.Println("Enter your username:")
	input.Input()
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
			fmt.Sprintf("Average attempts per game: %d", s.Stats.TotalAttempts)
		} else if answer == "no" {
			game.Exit()
		} else {
			fmt.Sprintf("Try again!")
		}
	}
}
