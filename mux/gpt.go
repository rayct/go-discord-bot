package mux

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func (m *Mux) GPT(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	ds.ChannelMessageSend(dm.ChannelID, "Gpt! "+(time.Since(dm.Timestamp)).String())
}
