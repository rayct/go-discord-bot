package mux

import (
	"GoDiscordBot/config"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/bwmarrin/discordgo"
)

func (m *Mux) GPT(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	client := gpt3.NewClient(config.OpenAIApiKey)

	response, err := getResponse(client, ctx.Content)
	if err != nil {
		fmt.Println(err)
	}

	embed := &discordgo.MessageEmbed{}
	embed.Title = "Chat GPT"
	embed.Color = 42622

	embed.Thumbnail = &discordgo.MessageEmbedThumbnail{}
	embed.Thumbnail.URL = "https://preview.redd.it/what-does-the-chatgpt-symbol-mean-does-it-have-meaning-like-v0-ogw5okm1bz5a1.jpg?auto=webp&s=51f427fda04a94889e4abfcbc6e6848b77be81ba"

	embed.Description = response

	_, err = ds.ChannelMessageSendEmbed(dm.ChannelID, embed)

	if err != nil {
		fmt.Println(err)
	}

}

func getResponse(client gpt3.Client, question string) (response string, err error) {
	sb := strings.Builder{}

	goCtx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	err = client.CompletionStreamWithEngine(
		goCtx,
		gpt3.TextDavinci003Engine,
		gpt3.CompletionRequest{
			Prompt: []string{
				question,
			},
			MaxTokens:   gpt3.IntPtr(3000),
			Temperature: gpt3.Float32Ptr(0.5),
		},
		func(resp *gpt3.CompletionResponse) {
			text := resp.Choices[0].Text

			sb.WriteString(text)
		},
	)
	if err != nil {
		return "", err
	}

	response = sb.String()
	response = strings.TrimLeft(response, "\n")

	return response, nil
}
