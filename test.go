package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
)

const (
	Green = "\033[32m"
	Yellow  = "\033[33m"
	Reset = "\033[0m"
)

func Input() string {
	fmt.Println("Input text:")
    reader := bufio.NewReader(os.Stdin)
    line, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Mistake")
    }
    return line
}

func IsLower(s string) bool {
	for _, ch := range s {
		if ch < 'a' || ch > 'z' {
			return false
		}
	}
	return true
}

func Game(input, word string){
	if IsLower(input) {
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

func Letters() string {
	var i int = 65
	var answer string
	for ; i <= 90; i++ {
		answer += string(rune(i))
	}
	return answer
}



func CheckLetters(input string) string {
	var check string
	for _, v := range input {
		check += string(v - 'a' + 'A')

	}
	return check
}

var word string = "hello"
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
			} else {
				fmt.Print("Try again")
			}

		}
		for _, v := range new {
			if rune(alphabet[i]) == v {
				if checkIn != true {
					i++
				}
			}
		}
		newAnswer += string(rune(alphabet[i]))+ " "
	}
	alphabet = newAnswer
	return alphabet
}

func StartGame(){
	var count int = 0
	var won bool = false
	for ;count <= 5; count++{
		input := Input()
		fmt.Println(UpdateLetters(input, word))
		Game(input, word)
		if input == word {
			won = true
			fmt.Println("You won! This word is ", word)
		}
		if count == 5 && won == false{
			fmt.Println("You lost!!!! This word is ", word)
		}
	}
}


func main() {
	StartGame()
}
