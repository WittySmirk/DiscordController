package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + "BOT_TOKEN")
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch m.Content {
	case "w":
		s.ChannelMessageSend(m.ChannelID, "walking...")
		walk()
	case "W":
		s.ChannelMessageSend(m.ChannelID, "running..")
		sprint()
	case "a":
		s.ChannelMessageSend(m.ChannelID, "looking right..")
		lR()
	case "d":
		s.ChannelMessageSend(m.ChannelID, "looking left..")
		lL()
	case "s":
		s.ChannelMessageSend(m.ChannelID, "walking back..")
		walkBack()
	case "j":
		s.ChannelMessageSend(m.ChannelID, "jumping..")
		jump()
	case "p":
		s.ChannelMessageSend(m.ChannelID, "planting/defusing..")
		defuse()
	case "q":
		s.ChannelMessageSend(m.ChannelID, "ability1..")
		a1()
	case "e":
		s.ChannelMessageSend(m.ChannelID, "ability2..")
		a2()
	case "c":
		s.ChannelMessageSend(m.ChannelID, "ability3..")
		a3()
	case "x":
		s.ChannelMessageSend(m.ChannelID, "ulting..")
		ult()
	case "t":
		s.ChannelMessageSend(m.ChannelID, "toggling aim in..")
		aim()
	case "y":
		s.ChannelMessageSend(m.ChannelID, "shooting..")
		shoot()
	case "f":
		s.ChannelMessageSend(m.ChannelID, "looking down..")
		lD()
	case "r":
		s.ChannelMessageSend(m.ChannelID, "reloading")
		reload()
	case "sw":
		s.ChannelMessageSend(m.ChannelID, "stopping walking..")
		stopWalking()
	case "SW":
		s.ChannelMessageSend(m.ChannelID, "stopping running..")
		stopSprinting()
	case "sd":
		s.ChannelMessageSend(m.ChannelID, "stopping defusing..")
		stopDefusing()
	case "ss":
		s.ChannelMessageSend(m.ChannelID, "stopping walking back//")
		stopWalkingB()
	case "u":
		s.ChannelMessageSend(m.ChannelID, "looking up..")
		lU()
	}
}