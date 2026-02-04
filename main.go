package main

import (
	"fmt"
	"koodWordle/input"
	"koodWordle/game"
)

func main(){
	fmt.Println("Welcome to Wordle! Guess the 5-letter word.")
	input.Input()
	fmt.Println(game.Letters())
}
