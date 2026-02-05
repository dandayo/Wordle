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
			checkIn := input[i]
			if input[i] == word[i] {
				result += Green + string(rune(input[i]) - 'a' + 'A') + Reset
			} else {
				found := false
				for j := 0 ; j < 5 ; j++ {
					if input[i] == word[j] {
						result += Yellow + string(rune(word[j]) - 'a' + 'A') + Reset
						found = true
						break
					}
				}
				if found == false {
					result += string(rune(checkIn) - 'a' + 'A')
				}
			}
		}
	} else {
		for l := 0; l < 5 ; l++{
			result += Green + string(rune(input[l]) - 'a' + 'A') + Reset
		}
	}
	fmt.Println("Feedback:",result)
}

func StartGame() {
	var count int = 0
	var attempts int = 5
	var won bool = false

	for count < 5 {
		fmt.Println("Enter your guess:")
		userInput := strings.TrimSpace(input.Input())

		if input.IsWord(userInput) == false{
			continue
		}

		fmt.Println("Remaining letters:",input.UpdateLetters(userInput, word))
		Game(userInput, word)

		if userInput == word {
			won = true
			fmt.Println("Congratulations! You've guessed the word correctly.")
			return
		}
		fmt.Println("Attempts remaining:", attempts)
		attempts--
		count++
	}
	if !won {
		fmt.Println("Game over.The correct word was:", word)
	}
	fmt.Println("Attempts:", count)
}
