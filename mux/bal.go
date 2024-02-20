package mux

import (
	"GoDiscordBot/db"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
)

// Help function provides a build in "help" command that will display a list
// of all registered routes (commands). To use this function it must first be
// registered with the Mux.Route function.
func (m *Mux) Bal(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	dbCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user db.User

	if err := db.UsersCollection.FindOne(dbCtx, bson.M{"Id": dm.Author.ID}).Decode(&user); err != nil {
		ds.ChannelMessageSend(dm.ChannelID, "User not found, try running the profit command to make a user account")
		return
	}

	embed := &discordgo.MessageEmbed{}
	embed.Title = "Balance"
	embed.Color = 0

	embed.Thumbnail = &discordgo.MessageEmbedThumbnail{}
	embed.Thumbnail.URL = "https://www.pngitem.com/pimgs/m/101-1016890_icon-bank-logo-png-transparent-png.png"

	embed.Description = user.Name + " has " + strconv.Itoa(user.Balance) + " units!"

	_, err := ds.ChannelMessageSendEmbed(dm.ChannelID, embed)

	if err != nil {
		fmt.Println(err)
	}

}
