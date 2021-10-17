package main

import (
	"fmt"
	"os"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}

	client := twitch.NewClient("oura_bot", os.Getenv("OAUTH"))

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message.Message)

		if message.User.DisplayName == "AuroR6S" && message.Message == "pajaS 🚨 ALERT" {
			client.Say(message.Channel, "PepeA 🚨 ALERT?")
		}
	})

	client.OnConnect(func() {
		fmt.Println("Connected!")
	})

	client.Join("auror6s")

	err = client.Connect()
	if err != nil {
		panic(err)
	}
}
