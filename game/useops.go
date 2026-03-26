package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func LoadWord() ([]string, error) { // take a words from file
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

func GetRandomIndex() int { //get a random index of list of words
	words, _ := LoadWord()
	n := len(words)
	return rand.Intn(n)
}

func GetWord() string { // get a random word
	n := GetRandomIndex()
	word, _ := LoadWord()
	return word[n]
}

func Exit() {
	fmt.Println("Press Enter to exit...")
	os.Exit(0)
}
