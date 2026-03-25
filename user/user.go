package user

import (
	"fmt"
	"koodWordle/input"
)

<<<<<<< HEAD
func UserName() {
=======


func UserName(){
>>>>>>> 423f628b3877e2a8053b1749bac2ff34b1d675c8
	fmt.Println("Enter your username:")
	input.Input()
}

<<<<<<< HEAD
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
=======
func GreetUser(){
	fmt.Println("Welcome to Wordle! Guess the 5-letter word.")
}

type UserStat struct {
	name string
	all	int
	won	int
	totalAtt int
}


/*func Stats(s *UserStat){
	fmt.Sprintf("Do you want to see your stats? (yes/no):")
	answer := input.Input()
	for {
		if answer == "yes"{
			fmt.Sprintf("Stats for: %s", *s.name)
			fmt.Sprintf("Games played: %s: ", *s.all)
			fmt.Sprintf("Games won: %s", *s.won)
			fmt.Sprintf("Average attempts per game: %f", *s.totalAtt)
		} else if answer == "no"{
			input.Exit()
		} else {
				fmt.Sprintf("Try again!")
		}
	}
	}*/
>>>>>>> 423f628b3877e2a8053b1749bac2ff34b1d675c8
