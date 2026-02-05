package user

import (
	"fmt"
	"koodWordle/input"
)



func UserName(){
	fmt.Println("Enter your username:")
	input.Input()
}

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
