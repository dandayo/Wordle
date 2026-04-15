# Wordle CLI in Go

This project is a terminal version of Wordle written in Go.

You enter a username, guess a hidden 5-letter word, get color feedback after each guess, and the game saves your stats in a CSV file.

## What This Project Does

- Runs a Wordle-style game in the terminal
- Chooses a random 5-letter word from a word list
- Lets the player guess until they win or run out of tries
- Shows letter feedback after every guess
- Tracks user stats across games

## How The Game Works

When you start the program:

1. The game prints a welcome message.
2. You enter your username.
3. The program loads your old stats from `user/user_stats.csv` if your username already exists.
4. A random word is selected from `words/wordle-words.txt`.
5. You type a 5-letter lowercase guess.
6. The game checks the guess and prints feedback.
7. When the game ends, your stats are updated and saved.
8. You can choose whether to view your stats.

## Feedback Meaning

The game prints your guess in uppercase letters with colors:

- Green: the letter is correct and in the correct position
- Yellow: the letter exists in the word, but in a different position
- No color: the letter is not in the word

It also shows the remaining available letters after each guess.

## Rules For Valid Input

Your guess must:

- Be exactly 5 letters long
- Use lowercase letters only
- Exist in the word list

If the guess is invalid, the game asks again.

## Attempts

- The game is designed around 6 guess cycles
- After each valid wrong guess, feedback is shown and the attempts display updates
- If you do not guess the word in time, the correct word is shown at the end

## User Stats

Stats are stored in `user/user_stats.csv`.

Each user record includes:

- Username
- Games played
- Games won
- Total attempts

After a game, the program can show:

- Games played
- Games won
- Average attempts per win

## Project Structure

```text
Wordle-main/
├── main.go
├── go.mod
├── README.md
├── game/
│   ├── game.go
│   └── useops.go
├── input/
│   └── input.go
├── user/
│   ├── user.go
│   └── user_stats.csv
└── words/
    └── wordle-words.txt
```

## File Overview

- `main.go`: starts the full program flow
- `game/game.go`: main game loop, word checking, and feedback output
- `game/useops.go`: loads words and picks a random answer
- `input/input.go`: reads user input and validates guesses
- `user/user.go`: handles usernames, loading stats, saving stats, and printing stats
- `words/wordle-words.txt`: list of valid answer words
- `user/user_stats.csv`: saved player statistics

## How To Run

Make sure Go is installed first.

From the project folder, run:

```bash
GOCACHE=$(pwd)/.gocache go run .
```

Using a local `GOCACHE` keeps the build files inside the project folder, which can help avoid permission issues on some systems.

## Example Play Flow

```text
Welcome to Wordle! Guess the 5-letter word.
Enter your username:
daniil
Enter your guess:
apple
Remaining letters: A B C D F G H I J K L M N O P Q R S T U V W X Y Z
Feedback: A P P L E
Attempts remaining: 5
```

The exact feedback colors depend on the hidden word.

## Notes

- The word is chosen randomly each time the program starts
- Stats are shared through the CSV file in the `user` folder
- This project is a CLI application, so everything happens in the terminal

## Possible Improvements

- Add a replay option without restarting the program
- Improve the attempt counter display
- Add tests for input validation and game logic
- Separate valid guesses and possible answer words
- Improve duplicate-letter handling to match official Wordle rules more closely
