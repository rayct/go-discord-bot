package mux

import (
	"math/rand"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func (m *Mux) HiLo(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	number := rand.Intn(101)
	ds.ChannelMessageSend(dm.ChannelID, "Guess a number between 0 and 100")
	guess := getGuess(ds, dm)

	for guess != number {
		if guess < number {
			ds.ChannelMessageSend(dm.ChannelID, "Too low, guess again")
		}
		if guess > number {
			ds.ChannelMessageSend(dm.ChannelID, "Too high, guess again")
		}
		guess = getGuess(ds, dm)
	}
	ds.ChannelMessageSend(dm.ChannelID, "Correct! The number was "+strconv.Itoa(number))

}

func getGuess(ds *discordgo.Session, dm *discordgo.Message) int {
	returnGuess, err := strconv.Atoi(GetUserMsg())

	for err != nil {
		ds.ChannelMessageSend(dm.ChannelID, "That is not a number, guess again.")
		returnGuess, err = strconv.Atoi(GetUserMsg())
	}
	return returnGuess
}

// func getGuess(ds *discordgo.Session, dm *discordgo.Message) int {
// 	// Send a message to the user to prompt for their guess
// 	ds.ChannelMessageSend(dm.ChannelID, "Enter your guess:")

// 	// Wait for a response from the user
// 	msg, err := ds.ChannelMessage(dm.ChannelID, dm.ID)
// 	if err != nil {
// 		// Handle error (e.g., unable to retrieve user's message)
// 		// For simplicity, you can send a message indicating an error
// 		ds.ChannelMessageSend(dm.ChannelID, "Error occurred. Please try again.")
// 		return getGuess(ds, dm) // Retry getting the guess
// 	}

// 	// Convert the user's message to an integer
// 	returnGuess, err := strconv.Atoi(msg.Content)
// 	if err != nil {
// 		// Handle error (e.g., user's input is not a valid number)
// 		ds.ChannelMessageSend(dm.ChannelID, "That is not a number, guess again.")
// 		return getGuess(ds, dm) // Retry getting the guess
// 	}

// 	return returnGuess
// }
