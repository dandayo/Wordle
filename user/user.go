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

func ReadStats() [][]string {
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

	return records
}

func UpdateUserStats(s *User, won bool, attempts int) {
	s.Stats.GamesPlayed++
	if won {
		s.Stats.GamesWon++
	}
	s.Stats.TotalAttempts += attempts

	records := ReadStats()

	found := false

	// Update existing user stats if user already exists in the file
	for i := 1; i < len(records); i++ {
		if records[i][0] == s.Name {
			records[i][1] = fmt.Sprintf("%d", s.Stats.GamesPlayed)
			records[i][2] = fmt.Sprintf("%d", s.Stats.GamesWon)
			records[i][3] = fmt.Sprintf("%d", s.Stats.TotalAttempts)
			found = true
			break
		}
	}

	if !found {
		records = append(records, []string{
			s.Name,
			fmt.Sprintf("%d", s.Stats.GamesPlayed),
			fmt.Sprintf("%d", s.Stats.GamesWon),
			fmt.Sprintf("%d", s.Stats.TotalAttempts),
		})
	}

	SaveAll(records)
}

func LoadUser(name string) *User { // function to check if user exists in csv file and load stats if exists
	records := ReadStats()
	for _, record := range records[1:] { // Skip header
		if record[0] == name {
			gamesPlayed := 0
			gamesWon := 0
			totalAttempts := 0
			fmt.Sscanf(record[1], "%d", &gamesPlayed)
			fmt.Sscanf(record[2], "%d", &gamesWon)
			fmt.Sscanf(record[3], "%d", &totalAttempts)

			return &User{
				Name: name,
				Stats: GameStats{
					GamesPlayed:   gamesPlayed,
					GamesWon:      gamesWon,
					TotalAttempts: totalAttempts,
				},
			}
		}
	}
	// or create a new user
	return &User{
		Name: name,
		Stats: GameStats{
			GamesPlayed:   0,
			GamesWon:      0,
			TotalAttempts: 0,
		},
	}
}

func SaveAll(records [][]string) {
	file, err := os.Create("user/user_stats.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	err = writer.WriteAll(records)
	if err != nil {
		fmt.Println("Error writing records:", err)
		return
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
			Stats(s)
		}
	}
}
