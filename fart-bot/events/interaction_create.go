package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/dickeyy/fart-bot/commands"
	"github.com/dickeyy/fart-bot/components"
	"github.com/rs/zerolog/log"
)

func init() {
	Events = append(Events, onInteractionCreate)
}

func onInteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		commands.OnInteraction(s, i)
	case discordgo.InteractionMessageComponent:
		components.OnInteraction(s, i)
	default:
		log.Warn().Msgf("Unknown interaction type %d", i.Type)
	}
}