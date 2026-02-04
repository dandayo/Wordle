package input

import (
	"fmt"
	"bufio"
	"os"
)

//Function for user input
func Input() string {
	fmt.Println("Input text:")
    reader := bufio.NewReader(os.Stdin)
    line, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Mistake")
    }
    return line
}
func Letters() string {
	var i int = 65
	var answer string
	for ; i <= 90; i++ {
		answer += string(rune(i))
	}
	return answer
}

var alphabet string = Letters()
var word1 string = "hello"

func CheckLetters(input string) string {
	var check string
	for _, v := range input {
		check += string(v - 'a' + 'A')

	}
	return check
}

func UpdateLetters(input, word1 string) string {
	new := CheckLetters(input)
	var newAnswer string
	for i := 0; i < len(alphabet); i++ {
		checkWord := rune(word1[i])
		for _, v := range new {
			if rune(alphabet[i]) == v{
				i++
			}
		}
		newAnswer += string(rune(alphabet[i]))
	}
	alphabet = newAnswer
	return alphabet
}
