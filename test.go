package main

import(

	"fmt"
	"bufio"
	"os"
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

func Game(input, word string) {
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

var alphabet string = Letters()

func CheckLetters(input string) string {
	var check string
	for _, v := range input {
		check += string(v - 'a' + 'A')

	}
	return check
}

//checkWord := rune(word1[i])

var word string = "hello"

func UpdateLetters(input, word string) string {
	new := CheckLetters(input)
	var newAnswer string
	for i := 0; i < len(alphabet); i++ {
		var checkIn bool = false
		for _, a := range word {
			if rune(alphabet[i]) == a {
				checkIn = true
			}
		}
		for _, v := range new {
			if rune(alphabet[i]) == v {
				if checkIn == true {
					newAnswer += string(rune(alphabet[i]))
				} else {
					i++
				}
			}
		}
		newAnswer += string(rune(alphabet[i]))
	}
	alphabet = newAnswer
}


func main() {
	input := Input()
	//fmt.Println(UpdateLetters(input, word1))
	Game(input, "hello")
}
