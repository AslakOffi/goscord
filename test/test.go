package main

import (
	"fmt"
	"github.com/Seyz123/yalis"
	"github.com/Seyz123/yalis/channel/message"
)

var client *yalis.Client

func main() {
	fmt.Println("Testing...")

	client = yalis.NewClient("NzM1NjQyNjE2NDc3MjUzNjg0.XxjOkw.DxpP72dLDdLbJ6IqE2OvV-zX7-k")

	_ = client.On("ready", OnReady)
	_ = client.On("message", OnMessage)

	if err := client.Login(); err != nil {
		panic(err)
	}

	select {}
}

func OnReady() {
	fmt.Println("Logged in as " + client.User().Tag())
}

func OnMessage(msg *message.Message) {
	if !msg.Author.Bot {
		_, _ = msg.Reply("Test")
	}
}
