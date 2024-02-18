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

	err = bot.Open()

	if err != nil {
		fmt.Println(err)
	}

	Router = mux.New()
	Router.Prefix = config.BotPrefix

	bot.AddHandler(Router.OnMessageCreate)

	Router.Route("ping", "Ping that returns latency", Router.Ping)
	Router.Route("embed", "Returns an embed", Router.Embed)

}

func main() {

	// Wait for a CTRL-C until SIGINT or SIGTERM is received.
	fmt.Println("Raybot is now running! Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	bot.Close()
}

// Exit Normally.
