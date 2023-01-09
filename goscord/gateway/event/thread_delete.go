package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/goccy/go-json"
)

type ThreadDelete struct {
	Data *discord.Channel `json:"d"`
}

func NewThreadDelete(rest *rest.Client, data []byte) (*ThreadDelete, error) {
	pk := new(ThreadDelete)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
