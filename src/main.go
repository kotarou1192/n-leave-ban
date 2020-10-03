package main

import (
	"fmt"

	"./bot/discord"
	"./bot/nleaveban"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	channel := make(chan *discord.Discord)
	nleaveban.Run(channel)
	a := <-channel
	fmt.Println(a)
}
