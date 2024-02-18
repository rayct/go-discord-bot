package mux

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func (m *Mux) Ping(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	ds.ChannelMessageSend(dm.ChannelID, "Pong! "+(time.Since(dm.Timestamp)).String())
}
