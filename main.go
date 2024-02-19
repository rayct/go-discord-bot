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

	var errDiscord error // Declare errDiscord to store discordgo.New error

	// Assign directly to the global bot variable
	bot, errDiscord = discordgo.New("Bot " + config.Token)

	if errDiscord != nil {
		fmt.Println(errDiscord)
	}

	err = bot.Open()

	if err != nil {
		fmt.Println(err)
	}

	Router = mux.New()
	Router.Prefix = config.BotPrefix

	bot.AddHandler(Router.OnMessageCreate)

	// Register Routes
	Router.Route("ping", "Ping that returns latency", Router.Ping)
	Router.Route("embed", "Returns an embed", Router.Embed)
	// Router.Route("gpt", "GPT Command", Router.GPT)
	Router.Route("hilo", "High Low Guessing Command", Router.HiLo)
	Router.Route("tictactoe", "Tic Tac Toe Command", Router.TicTacToe)

}

func main() {

	// Wait for a CTRL-C until SIGINT or SIGTERM is received.
	fmt.Println("Raybot is now running! Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	if err := bot.Close(); err != nil {
		fmt.Printf("Error closing Discord session: %v\n", err)
		// Handle error gracefully, perhaps log it or perform other actions.
	}
}

// Exit Normally.
