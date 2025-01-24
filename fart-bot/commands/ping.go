package commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/dickeyy/perpbot/components"
	"github.com/dickeyy/perpbot/services"
	"github.com/dickeyy/perpbot/utils"
	"github.com/rs/zerolog/log"
	"github.com/sgaunet/perplexity-go/v2"
)

func init() {
	services.Commands[pingCmd.Name] = &services.Command{
		ApplicationCommand: pingCmd,
		Handler:            handlePing,
	}
}

var pingCmd = &discordgo.ApplicationCommand{
	Type:        discordgo.ChatApplicationCommand,
	Name:        "ping",
	Description: "Ping the bot",
}

func handlePing(s *discordgo.Session, i *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	return ContentResponse("Pong!", false
}