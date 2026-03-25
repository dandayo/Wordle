package user

import (
	"fmt"
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
	fmt.Sprintf("Do you want to see your stats? (yes/no):")
	answer := input.Input()
	for {
		if answer == "yes" {
			fmt.Sprintf("Stats for: %s", *s.name)
			fmt.Sprintf("Games played: %d: ", *s.GamesPlayed)
			fmt.Sprintf("Games won: %d", *s.GamesWon)
			fmt.Sprintf("Average attempts per game: %d", *s.TotalAttempts)
		} else if answer == "no" {
			input.Exit()
		} else {
			fmt.Sprintf("Try again!")
		}
	}
}
