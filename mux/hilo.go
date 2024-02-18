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
		ds.ChannelMessageSend(dm.ChannelID, "That is not a number, guess again")
		returnGuess, err = strconv.Atoi(GetUserMsg())
	}
	return returnGuess
}
