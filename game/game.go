package game

import (
	"bufio"
	"fmt"
	"os"
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

func Game(input, word string) {
	if !IsLower(input) {
		fmt.Println("Use only lowercase")
		return
	}
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
