package mux

import (
	"math/rand"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var grid = [3][3]string{}

func (m *Mux) TicTacToe(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	gameEmbed := &discordgo.MessageEmbed{}
	gameEmbed.Title = "Tic Tac Toe"
	gameEmbed.Color = 11845097

	grid = [3][3]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"}}
	tilesLeft := 9

	for !gameWon("X") {
		printGame(ds, dm, gameEmbed)
		ds.ChannelMessageSend(dm.ChannelID, "You are \"x\"\nType your move as row column\nEx: 0 0")
		userMoveStr := GetUserMsg()
		userMove := strings.Split(userMoveStr, " ")
		if len(userMove) == 2 {
			row, rowErr := strconv.Atoi(userMove[0])
			column, columnErr := strconv.Atoi(userMove[1])
			if rowErr != nil || columnErr != nil {
				ds.ChannelMessageSend(dm.ChannelID, "Only numbers accepted, game cancelled")
				return
			}
			if row < 0 || row > 2 || column < 0 || column > 2 {
				ds.ChannelMessageSend(dm.ChannelID, "Out of bound values, game cancelled")
				return
			}
			grid[row][column] = "X"
			tilesLeft -= 1
			if tilesLeft > 0 {
				computerMove := []int{rand.Intn(3), rand.Intn(3)}
				for grid[computerMove[0]][computerMove[1]] == "X" {
					computerMove = []int{rand.Intn(2), rand.Intn(2)}
				}
				grid[computerMove[0]][computerMove[1]] = "O"
				tilesLeft -= 1
			}
			if tilesLeft <= 0 {
				printGame(ds, dm, gameEmbed)
				ds.ChannelMessageSend(dm.ChannelID, "Game Tied!")
				return
			}
			if gameWon("O") {
				printGame(ds, dm, gameEmbed)
				ds.ChannelMessageSend(dm.ChannelID, "You Lost!")
				return
			}
		} else {
			ds.ChannelMessageSend(dm.ChannelID, "Wrong number of arguments, game cancelled")
			return
		}

	}
	printGame(ds, dm, gameEmbed)
	ds.ChannelMessageSend(dm.ChannelID, "You Win!")
}

func printGame(ds *discordgo.Session, dm *discordgo.Message, game *discordgo.MessageEmbed) {
	game.Description = ""

	for _, row := range grid {
		game.Description += strings.Join(row[:], " ") + "\n"
	}
	ds.ChannelMessageSendEmbed(dm.ChannelID, game)
}

func gameWon(val string) bool {
	for i := 0; i < 3; i++ {
		if grid[i][0] == val && grid[i][1] == val && grid[i][2] == val {
			return true
		}
	}
	for j := 0; j < 3; j++ {
		if grid[0][j] == val && grid[1][j] == val && grid[2][j] == val {
			return true
		}
	}
	if grid[1][1] == val {
		if grid[0][0] == val && grid[2][2] == val {
			return true
		}
		if grid[0][2] == val && grid[2][0] == val {
			return true
		}
	}
	return false
}
