package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function for user input
func Input() string {
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
		answer += string(rune(i)) + " "
	}
	return answer
}

<<<<<<< HEAD
=======

>>>>>>> 423f628b3877e2a8053b1749bac2ff34b1d675c8
func IsWord(input string) bool {
	if len(input) != 5 {
		fmt.Println("Your guess must be exactly 5-letter word.")
		return false
	}
	for _, ch := range input {
		if ch < 'a' || ch > 'z' {
			fmt.Println("Your guess must only contain lowercase letters.")
			return false
		}
	}
	return true
}

<<<<<<< HEAD
=======

>>>>>>> 423f628b3877e2a8053b1749bac2ff34b1d675c8
func CheckLetters(input string) string {
	var check string
	for _, v := range input {
		check += string(v - 'a' + 'A')

	}
	return check
}
<<<<<<< HEAD

=======
>>>>>>> 423f628b3877e2a8053b1749bac2ff34b1d675c8
var alphabet string = Letters()

func UpdateLetters(input, word string) string {
	new := CheckLetters(input)
	var newAnswer string
	newAlphabet := strings.TrimSpace(alphabet)
	for i := 0; i < len(newAlphabet); i++ {
		var checkIn bool = false
		for _, a := range word {
			if rune(newAlphabet[i]) == a-'a'+'A' {
				checkIn = true
			}
		}
		for _, v := range new {
			if rune(newAlphabet[i]) == v {
				if checkIn != true {
					i++
				}
			}
		}
		newAnswer += string(rune(newAlphabet[i]))
	}
	alphabet = newAnswer
	return alphabet
}

func LoadWord() ([]string, error) {
<<<<<<< HEAD
	file, err := os.Open("words/wordle-words.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words, scanner.Err()
}

func GetIndex() int {
=======
    file, err := os.Open("wordle-words.txt")
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var words []string
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        words = append(words, scanner.Text())
    }

    return words, scanner.Err()
}

func GetIndex() int{
>>>>>>> 423f628b3877e2a8053b1749bac2ff34b1d675c8
	n, _ := strconv.Atoi(os.Args[1])
	return n
}

<<<<<<< HEAD
func GetWord() string {
=======
func GetWord() string{
>>>>>>> 423f628b3877e2a8053b1749bac2ff34b1d675c8
	n := GetIndex()
	word, _ := LoadWord()
	return word[n]
}

<<<<<<< HEAD
func Exit() {
=======
func Exit(){
>>>>>>> 423f628b3877e2a8053b1749bac2ff34b1d675c8
	fmt.Println("Press Enter to exit...")
	os.Exit(0)
}
