package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Goscord/goscord"
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/discord/embed"
	"github.com/Goscord/goscord/gateway"
)

var client *gateway.Session

func main() {
	fmt.Println("Testing...")

	client = goscord.New(&gateway.Options{
		Token:   "",
		Intents: gateway.IntentGuilds | gateway.IntentGuildMembers | gateway.IntentGuildMessages,
	})

	client.On("ready", OnReady)
	client.On("messageCreate", OnMessageCreate)
	client.On("guildMemberAdd", OnGuildMemberAdd)
	client.On("interactionCreate", OnInteractionCreate)

	if err := client.Login(); err != nil {
		panic(err)
	}

	select {}
}

func OnReady() {
	fmt.Println("Logged in as " + client.Me().Tag())

	appCmd := &discord.ApplicationCommand{
		Name:        "ping",
		Type:        discord.ApplicationCommandChat,
		Description: "Pong pong pong!",
	}

	_ = client.Interaction.RegisterCommand(client.Me().Id, "", appCmd)

	_ = client.SetActivity(&discord.Activity{Name: "Luther - ALAKAZAM", Type: discord.ActivityListening})
	_ = client.SetStatus("idle")
}

func OnInteractionCreate(i *discord.Interaction) {
	if i.Type == discord.InteractionTypeApplicationCommand {

		_, err := client.Interaction.CreateResponse(i.Id, i.Token, "Pong! lol")
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		}
	}
}

func OnGuildMemberAdd(a *discord.GuildMember) {
	if c, ok := client.State().Channel("1001943782016688292"); ok == nil {
		client.Channel.SendMessage(c.Id, "Welcome to the server <@"+a.User.Id+"> !")
	} else {
		fmt.Println("Could not find channel")
	}
}

func OnMessageCreate(msg *discord.Message) {
	if strings.ToLower(msg.Content) == "ping" {
		_, _ = client.Channel.SendMessage(msg.ChannelId, "Pong!")
	}

	if strings.ToLower(msg.Content) == "embed" {
		embed := embed.NewEmbedBuilder()
		embed.SetAuthor("Testing", "")
		embed.SetDescription("Just testing some things lol")
		_, _ = client.Channel.SendMessage(msg.ChannelId, embed)
	}

	if strings.ToLower(msg.Content) == "reply" {
		_, _ = client.Channel.ReplyMessage(msg.ChannelId, msg.Id, "Replied to you!")
	}

	if strings.ToLower(msg.Content) == "serverinfo" {
		guild, err := client.State().Guild(msg.GuildId)

		if err != nil {
			_, _ = client.Channel.SendMessage(msg.ChannelId, err.Error())
			return
		}

		_, _ = client.Channel.SendMessage(msg.ChannelId, "Server name: "+guild.Name+"\nID: "+guild.Id+"\nMembers: "+fmt.Sprintf("%d", guild.MemberCount))
	}

	if strings.ToLower(msg.Content) == "memberinfo" {
		member, err := client.State().Member(msg.GuildId, msg.Author.Id)

		if err != nil {
			_, _ = client.Channel.SendMessage(msg.ChannelId, err.Error())
			return
		}

		_, _ = client.Channel.SendMessage(msg.ChannelId, "Username: "+member.User.Username+"\nID: "+member.User.Id+"\nNickname: "+member.Nick)
	}

	if strings.ToLower(msg.Content) == "dogeimage" {
		dogeImg, err := os.Open("doge.png")
		if err != nil {
			panic(err)
		}

		defer dogeImg.Close()

		_, _ = client.Channel.SendMessage(msg.ChannelId, dogeImg)
	}
}
