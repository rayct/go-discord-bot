package mux

import (
	"GoDiscordBot/db"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
)

// Create Function

func (m *Mux) Profit(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	var user db.User
	dbCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for {
		if err := db.UsersCollection.FindOne(dbCtx, bson.M{"Id": dm.Author.ID}).Decode(&user); err == nil {
			break
		}

		userResult, err := db.UsersCollection.InsertOne(dbCtx, bson.D{
			{Key: "Id", Value: dm.Author.ID},
			{Key: "Name", Value: dm.Author.Username},
			{Key: "Balance", Value: 0},
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(userResult.InsertedID)
	}

	err := db.UsersCollection.FindOneAndUpdate(
		dbCtx,
		bson.M{"Id": dm.Author.ID},
		bson.D{
			{Key: "$set", Value: bson.D{{Key: "Balance", Value: user.Balance + 100}}},
		},
	).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	embed := &discordgo.MessageEmbed{}
	embed.Title = "Profit"
	embed.Color = 42622

	embed.Thumbnail = &discordgo.MessageEmbedThumbnail{}
	embed.Thumbnail.URL = "https://e7.pngegg.com/pngimages/450/717/png-clipart-dollar-dollar.png"

	embed.Description = user.Name + " earned 100 units!"

	_, err = ds.ChannelMessageSendEmbed(dm.ChannelID, embed)

	if err != nil {
		fmt.Println(err)
	}

}
