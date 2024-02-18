package mux

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (m *Mux) Embed(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	embed := &discordgo.MessageEmbed{}
	embed.Title = "This is an embed"
	embed.Color = 1752220
	embed.Description = "This is a description"

	// embed.Video = &discordgo.MessageEmbedVideo{}
	// embed.Video.URL = "https://gifdb.com/images/high/singer-rick-astley-never-gonna-give-you-up-bfkvglwju2mte6ff.gif"

	embed.Thumbnail = &discordgo.MessageEmbedThumbnail{}
	embed.Thumbnail.URL = "https://hips.hearstapps.com/hmg-prod/images/Emma-Watson_GettyImages-619546914.jpg?crop=1xw:1.0xh;center,top&resize=640:*"

	_, err := ds.ChannelMessageSendEmbed(dm.ChannelID, embed)

	if err != nil {
		fmt.Println(err)
	}
}
