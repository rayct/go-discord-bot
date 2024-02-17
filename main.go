package main

import (
	"GoDiscordBot/config"
	"GoDiscordBot/mux"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	bot    *discordgo.Session
	Router *mux.Mux
)

func init() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err)
	}

	bot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err)
	}

	Router = mux.New()

	bot.AddHandler(Router.OnMessageCreate)

}

func main() {
	err := bot.Open()

	if err != nil {
		fmt.Println(err)
	}

	// Wait for a CTRL-C
	fmt.Println("Bot is now running! Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Clean up
	bot.Close()
}

// Exit Normally.

// func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
// 	if m.Author.ID == BotID {
// 		return
// 	}

// 	if m.Content == config.BotPrefix+" ping" {
// 		_, err := s.ChannelMessageSend(m.ChannelID, "Hello! "+(time.Since(m.Timestamp)).String())

// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 	}

// 	if m.Content == config.BotPrefix+" embed" {
// 		embed := &discordgo.MessageEmbed{}
// 		embed.Title = "This is an embed"
// 		embed.Color = 1752220
// 		embed.Description = "This is a description"

// 		// embed.Video = &discordgo.MessageEmbedVideo{}
// 		// embed.Video.URL = "https://gifdb.com/images/high/singer-rick-astley-never-gonna-give-you-up-bfkvglwju2mte6ff.gif"

// 		embed.Thumbnail = &discordgo.MessageEmbedThumbnail{}
// 		embed.Thumbnail.URL = "https://hips.hearstapps.com/hmg-prod/images/Emma-Watson_GettyImages-619546914.jpg?crop=1xw:1.0xh;center,top&resize=640:*"

// 		_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)

// 		if err != nil {
// 			fmt.Println(err)
// 		}

// 	}
// }
