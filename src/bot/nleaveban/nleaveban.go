package nleaveban

import "../discord"

func Run(channel chan *discord.Discord) {
	bot := discord.New()
	bot.StartWatching()
	err := bot.Session.Open()
	if err != nil {
		panic(err)
	}
}
