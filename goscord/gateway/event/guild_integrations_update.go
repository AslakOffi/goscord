package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

type GuildIntegrationsUpdate struct {
	Data *discord.GuildIntegrationsUpdateEventFields `json:"d"`
}

func NewGuildIntegrationsUpdate(rest *rest.Client, data []byte) (*GuildIntegrationsUpdate, error) {
	pk := new(GuildIntegrationsUpdate)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
