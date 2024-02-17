package main

import (
	"GoDiscordBot/config"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	BotID string
	bot   *discordgo.Session
)

func Start() {
	bot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err)
	}

	bot.AddHandler(messageHandler)

	err = bot.Open()
	if err != nil {
		fmt.Println(err)
	}

	BotID = bot.State.User.ID

	fmt.Println("Bot is now running!")

}

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err)
	}

	Start()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	bot.Close()
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Content == config.BotPrefix+" ping" {
		_, err := s.ChannelMessageSend(m.ChannelID, "Pong! "+(time.Since(m.Timestamp)).String())

		if err != nil {
			fmt.Println(err)
		}
	}
}
