package user

import (
	"encoding/csv"
	"fmt"
	"koodWordle/game"
	"koodWordle/input"
	"os"
)

func UserName() string {
	fmt.Println("Enter your username:")
	name := input.Input()

	return name
}

func CreateUser(name string) *User {
	return &User{
		Name: name,
		Stats: GameStats{
			GamesPlayed:   0,
			GamesWon:      0,
			TotalAttempts: 0,
		},
	}
}

func GreetUser() {
	fmt.Println("Welcome to Wordle! Guess the 5-letter word.")
}

type User struct {
	Name  string
	Stats GameStats
}

type GameStats struct {
	GamesPlayed   int
	GamesWon      int
	TotalAttempts int
}

func UpdateUserStats(s *User, won bool, attempts int) {
	s.Stats.GamesPlayed++
	if won {
		s.Stats.GamesWon++
	}
	s.Stats.TotalAttempts += attempts
}

func LoadUser(name string) *User {
	file, err := os.Open("user/user_stats.csv")
	if err != nil {
		fmt.Println("Error opening stats file:", err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading stats file:", err)
		return nil
	}
	for _, record := range records[1:] { // Skip header
		if record[0] == name {

		}
	}
	return nil
}

func GetUser(s *User) { //get user stats and save to csv file
	file, err := os.OpenFile(
		"user/user_stats.csv",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	defer file.Close()
	writer := csv.NewWriter(file)

	info, _ := file.Stat()
	isEmpty := info.Size() == 0

	if isEmpty {
		writer.Write([]string{
			"Username",
			"Games Played",
			"Games Won",
			"Total Attempts",
		})
	}

	err = writer.Write([]string{s.Name, fmt.Sprintf("%d", s.Stats.GamesPlayed), fmt.Sprintf("%d", s.Stats.GamesWon), fmt.Sprintf("%d", s.Stats.TotalAttempts)}) // Save stats to CSV file

	if err != nil {
		fmt.Println("Error writing stats to file:", err)
		return
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Println("Error flushing CSV writer:", err)
	}
}

func Stats(s *User) {
	fmt.Println("Do you want to see your stats? (yes/no):")
	answer := input.Input()
	for {
		if answer == "yes" {
			fmt.Printf("\nStats for: %s\n", s.Name)
			fmt.Printf("Games played: %d\n", s.Stats.GamesPlayed)
			fmt.Printf("Games won: %d\n", s.Stats.GamesWon)
			if s.Stats.GamesWon > 0 {
				fmt.Printf("Average attempts per game: %d\n", s.Stats.TotalAttempts/s.Stats.GamesWon)
			} else {
				fmt.Printf("No wins yet\n")
			}
			break
		} else if answer == "no" {
			game.Exit()
		} else {
			fmt.Sprintf("Try again!")
			break
		}
	}
	GetUser(s)
}
