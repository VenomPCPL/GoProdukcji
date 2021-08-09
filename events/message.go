package events

import (
	"fmt"
	"github.com/BOOMfinity-Developers/bfcord/client/state"
	"github.com/BOOMfinity-Developers/bfcord/discord"
	"github.com/BOOMfinity-Developers/bfcord/discord/colors"
	"github.com/BOOMfinity-Developers/bfcord/gateway"
	"log"
	"os"
)

var prefix = "np!"

func HandleMessageCreate(c state.IState) func(message gateway.MessageCreateEvent) {
	return func(message gateway.MessageCreateEvent) {
		channel, _ := message.Channel().Get()
		//Repost all announcements/tweets
		if channel.Type == discord.GuildNewsChannel {
			err := message.Crosspost()
			if err != nil {
				log.Fatal(err)
			}
		}

		//Stats command
		if message.Content == prefix+"stats" {
			_, sendErr := message.Channel().SendMessage(&discord.MessageCreateOptions{
				Embed: &discord.MessageEmbed{
					Title: "GoProdukcji Stats",
					Description: fmt.Sprintf("Gateway ping: %vms\n"+
						"Version: [%v](https://github.com/MrBoombastic/GoProdukcji/commit/%v)",
						c.Manager().AveragePing(), os.Args[3], os.Args[3]),
					Color: colors.Orange,
				}})

			if sendErr != nil {
				log.Fatal(sendErr)
			}
		}
	}
}
