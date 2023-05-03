package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

type ChannelPinsUpdate struct {
	Data *discord.ChannelPinsUpdateEventFields `json:"d"`
}

func NewChannelPinsUpdate(rest *rest.Client, data []byte) (*ChannelPinsUpdate, error) {
	pk := new(ChannelPinsUpdate)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
