package main

import (
	"koodWordle/game"
	"koodWordle/user"
)

func main() {
	user.GreetUser()
	name := user.UserName()
	currentUser := user.LoadUser(name)
	attempts, won := game.StartGame()
	user.UpdateUserStats(currentUser, won, attempts)
	user.Stats(currentUser)
}
