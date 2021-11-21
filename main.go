package main

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/joho/godotenv"
)

func main() {
	lastPajas := 0
	cooldowns := make(map[string]int)

	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	client := twitch.NewClient("oura_bot", os.Getenv("OAUTH"))

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		if message.User.DisplayName == "pajbot" && message.Message == "pajaS 🚨 ALERT" {
			client.Say(message.Channel, "/me PepeA 🚨 ALERT?")
			lastPajas = int(message.Time.Unix())
		}

		match, _ := regexp.MatchString("^!shuffle ((/me|pajaS|🚨|ALERT) *){3,}", message.Message)

		if match {
			client.Say(message.Channel, "pajaCMON shufflers")
		}

		if (message.Message == "pajaS ❓" || message.Message == "pajaS ?") && cooldowns[message.User.DisplayName] == 0 {
			cooldowns[message.User.DisplayName] = 1
			go func() {
				time.Sleep(5 * time.Second)
				cooldowns[message.User.DisplayName] = 0
			}()

			if lastPajas == 0 {
				client.Say(message.Channel, "I haven't been up for the pajaS 🚨 yet")
				return
			}

			nextPajas := lastPajas + (2 * 60 * 60)
			client.Say(message.Channel, fmt.Sprintf("%s, PAJAS 🚨 in %s", message.User.DisplayName, time.Until(time.Unix(int64(nextPajas), 0)).Round(time.Second)))

		}

		if message.Message == "PepeA ping?" && message.User.DisplayName == "AuroR6S" {
			client.Say(message.Channel, "PAJAS Pong!")
		}

		if message.User.DisplayName == "slchbot" && message.Message == "PepeA pajbot" {
			client.Say(message.Channel, "/me GachiPls [emote] 🚨")
		}
	})

	client.OnConnect(func() {
		fmt.Println("Connected!")
	})

	client.Join("pajlada")

	err = client.Connect()
	if err != nil {
		panic(err)
	}
}
