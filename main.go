package main

import (
	"koodWordle/game"
	"koodWordle/user"
)

func main(){
	user.UserName()
	user.GreetUser()
	game.StartGame()
	user.Stats()
}
