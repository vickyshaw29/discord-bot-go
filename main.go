package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/vickyshaw29/discord-goBot/x/mux"
)

var Session, _ = discordgo.New()
var Router = mux.New()

func init() {
	test := godotenv.Load()
	if test != nil {
		log.Fatal("Error loading .env file")
	}

	Session.Token = os.Getenv("TOKEN")
	Session.AddHandler(Router.OnMessageCreate)
	Router.Route("/help", "Display this message.", Router.Help)
	Router.Route("/joke", "Send a joke", Router.Joke)
	Router.Route("/quote", "Send a quote", Router.Quote)
}

func main() {
	test := godotenv.Load()
	Session.Token = os.Getenv("TOKEN")
	if test != nil {
		log.Fatal("Error loading .env file")
	}

	err := Session.Open()
	if err != nil {
		log.Printf("error opening connection to Discord, %s\n", err)
		os.Exit(1)
	}

	// Wait for a CTRL-C
	log.Printf(`Now running. Press CTRL-C to exit.`)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-sc

	// Clean up
	Session.Close()
}
