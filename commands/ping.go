package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func RegisterPingCommand(s *discordgo.Session) {
	_, err := s.ApplicationCommandCreate(s.State.User.ID, "", &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Check if the bot is alive",
	})
	if err != nil {
		fmt.Println("Error creating ping command:", err)
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == discordgo.InteractionApplicationCommand && i.ApplicationCommandData().Name == "ping" {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "🏓 Pong!",
				},
			})
		}
	})
}
