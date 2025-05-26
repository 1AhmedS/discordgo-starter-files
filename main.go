package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"discordgo-starter-files/commands"

	"github.com/bwmarrin/discordgo"
)

var TOKEN = ""

func main() {
	Token := "Bot " + TOKEN

	session, err := discordgo.New(Token)
	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}

	session.AddHandler(Ready)
	session.Identify.Intents = discordgo.IntentsGuilds

	err = session.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")

	// Register Slash Commands
	commands.RegisterPingCommand(session)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	session.Close()
}

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Println("Bot is ready. Username:", s.State.User.Username)

	err := s.UpdateStatusComplex(discordgo.UpdateStatusData{
		Status: "online",
		Activities: []*discordgo.Activity{
			{
				Name: "With You",
				Type: discordgo.ActivityTypeGame,
			},
		},
	})

	if err != nil {
		fmt.Println("Error setting status:", err)
	}
}
