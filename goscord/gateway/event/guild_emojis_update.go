package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/goccy/go-json"
)

type GuildEmojisUpdate struct {
	Data *discord.GuildEmojisUpdateEventFields `json:"d"`
}

func NewGuildEmojisUpdate(rest *rest.Client, data []byte) (*GuildEmojisUpdate, error) {
	pk := new(GuildEmojisUpdate)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
