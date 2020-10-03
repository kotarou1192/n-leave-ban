package discord

import (
	"fmt"
	"log"

	"../../models/user"
	"../../utils/envloader"
	"github.com/bwmarrin/discordgo"
)

type Discord struct {
	UserID  string
	IsLeft  bool
	IsJoin  bool
	Session *discordgo.Session
}

func New() *Discord {
	config := envloader.LoadConfig()
	session, err := discordgo.New("Bot " + config.TOKEN)
	if err != nil {
		panic(err)
	}

	discord := Discord{Session: session}
	return &discord
}

func (bot *Discord) StartWatching() {
	bot.Session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.Bot {
			return
		}
		if m.Message.Type == discordgo.MessageTypeGuildMemberJoin {
			fmt.Println(m.Author.ID)
			nLeaveBanCore(s, m)
		}
	})
}

func nLeaveBanCore(s *discordgo.Session, m *discordgo.MessageCreate) {
	config := envloader.LoadConfig()
	userID := m.Author.ID
	log.Printf("JOIN SERVER: NAME: %s ID: %s \n", m.Author.Username, m.Author.ID)

	selectedUser, err := user.Find(userID)
	if err != nil {
		selectedUser, err = user.New(userID).Save()
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			return
		}
	}
	selectedUser.AddLeaveTime()
	leaveCount, err := selectedUser.GetLeaveCount()
	sendInfo(s, m, config.LEAVE_LIMIT-leaveCount+1)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(leaveCount, config.LEAVE_LIMIT)
	if leaveCount > config.LEAVE_LIMIT {
		banUser(s, m)
	}
}

func banUser(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Printf("BAN USER ID: %s NAME: %s", m.Author.ID, m.Author.Username)
	err := s.GuildBanCreateWithReason(m.GuildID, m.Author.ID, "サーバーへの出入りが規定回数を超えたためBAN", 7)
	log.Println(err)
}

func sendInfo(s *discordgo.Session, m *discordgo.MessageCreate, count int) {
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("ようこそ %s さん。こちら、ユーザー管理システムです。あと %d 回サーバーを抜けると自動的にBANされます。ご注意ください。", m.Author.Username, count))
}
