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

func Stats(name string, ){
	fmt.Sprintf("Do you want to see your stats? (yes/no):")
	answer := input.Input()
	for {
		if answer == "yes"{
			fmt.Sprintf("Stats for: %s", name)
			fmt.Sprintf("Games played: %s: ", all)
			fmt.Sprintf("Games won: %s", won)
			fmt.Sprintf("Average attempts per game: %f", avg)
		} else if answer == "no"{
			input.Exit()
		} else {
				fmt.Sprintf("Try again!")
		}
	}
}
