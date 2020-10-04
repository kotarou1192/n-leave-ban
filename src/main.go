package main

import (
	"fmt"

	"./bot/discord"
	"./bot/nleaveban"
	"./models/user"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	user.AddIndex()
	channel := make(chan *discord.Discord)
	nleaveban.Run(channel)
	a := <-channel
	fmt.Println(a)
}
