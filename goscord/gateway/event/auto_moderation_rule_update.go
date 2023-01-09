package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/goccy/go-json"
)

type AutoModerationRuleUpdate struct {
	Data *discord.AutoModerationRule `json:"d"`
}

func NewAutoModerationRuleUpdate(rest *rest.Client, data []byte) (*AutoModerationRuleUpdate, error) {
	pk := new(AutoModerationRuleUpdate)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
