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

func UpdateUserStats(s *User, won bool, attempts int) {
	s.Stats.GamesPlayed++
	if won {
		s.Stats.GamesWon++
	}
	s.Stats.TotalAttempts += attempts
}

func Stats(s *User) {
	fmt.Println("Do you want to see your stats? (yes/no):")
	answer := input.Input()
	for {
		if answer == "yes" {
			fmt.Printf("\nStats for: %s\n", s.Name)
			fmt.Printf("Games played: %d\n", s.Stats.GamesPlayed)
			fmt.Printf("Games won: %d\n", s.Stats.GamesWon)
			if s.Stats.GamesWon > 0 {
				fmt.Printf("Average attempts per game: %d\n", s.Stats.TotalAttempts/s.Stats.GamesWon)
			} else {
				fmt.Printf("No wins yet\n")
			}
			break
		} else if answer == "no" {
			game.Exit()
		} else {
			fmt.Sprintf("Try again!")
			break
		}
	}
}
