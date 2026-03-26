package main

import (
	"koodWordle/game"
	"koodWordle/user"
)

func main() {
	user.GreetUser()
	currentUser := user.UserName()
	attempts, won := game.StartGame()
	user.UpdateStats(currentUser, attempts, won)
	user.showStat()
}
