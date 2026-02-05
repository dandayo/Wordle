package game

import (
	"fmt"
	"strings"
	"koodWordle/input"
)

const (
	Green = "\033[32m"
	Yellow  = "\033[33m"
	Reset = "\033[0m"
)

func IsLower(s string) bool {
	for _, ch := range s {
		if ch < 'a' || ch > 'z' {
			return false
		}
	}
	return true
}

var word string = "hello"
func Game(input, word string){
	var result string
	if input != word {
		for i := 0; i < 5; i++ {
			checkIn := string(input[i])
			if input[i] == word[i] {
				result += Green + checkIn + Reset
			} else {
				found := false
				for j := 0 ; j < 5 ; j++ {
					if input[i] == word[j] {
						result += Yellow + string(word[j]) + Reset
						found = true
						break
					}
				}
				if found == false {
					result += checkIn
				}
			}
		}
	} else {
		result += Green + input + Reset
	}
	fmt.Println(result)
}

func StartGame(){
	var count int = 0
	var won bool = false

	for count < 5 {
		userInput := strings.TrimSpace(input.Input())

		if input.IsWord(userInput) == false{
			continue
		}

		fmt.Println(input.UpdateLetters(userInput, word))
		Game(userInput, word)

		if userInput == word {
			won = true
			fmt.Println("You won! This word is", word)
			return
		}
		 count++
	}
	if !won {
		fmt.Println("You lost!!!! This word is", word)
	}
	fmt.Println("Attempts:", count)
}
