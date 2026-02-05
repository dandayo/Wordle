package input

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

//Function for user input
func Input() string {
	fmt.Println("Input text:")
    reader := bufio.NewReader(os.Stdin)
    line, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Mistake")
    }
    return strings.TrimSpace(line)
}

func Letters() string {
	var i int = 65
	var answer string
	for ; i <= 90; i++ {
		answer += string(rune(i))
	}
	return answer
}


func IsWord(input string) bool {
	if len(input) != 5 {
		fmt.Println("Word must be 5 letters and lowcase")
		return false
	}
	for _, ch := range input {
		if ch < 'a' || ch > 'z' {
			fmt.Println("Word must be 5 letters and lowcase")
			return false
		}
	}
	return true
}


func CheckLetters(input string) string {
	var check string
	for _, v := range input {
		check += string(v - 'a' + 'A')

	}
	return check
}
var alphabet string = Letters()

func UpdateLetters(input, word string) string {
	new := CheckLetters(input)
	var newAnswer string
	alphabet = strings.TrimSpace(alphabet)
	for i := 0; i < len(alphabet); i++ {
		var checkIn bool = false
		for _, a := range word {
			if rune(alphabet[i]) == a-'a'+'A' {
				checkIn = true
			}
		}
		for _, v := range new {
			if rune(alphabet[i]) == v {
				if checkIn != true {
					i++
				}
			}
		}
		newAnswer += string(rune(alphabet[i]))
	}
	alphabet = newAnswer
	return alphabet
}
