package discord

import (
	"encoding/json"
	"github.com/Goscord/goscord/rest"
)

type Guild struct {
	Rest              *rest.Client `json:"-"`
	Id                string       `json:"id"`
	Name              string       `json:"name"`
	Icon              string       `json:"icon,omitempty"`
	Splash            string       `json:"splash,omitempty"`
	DiscoverySplash   string       `json:"discovery_splash,omitempty"`
	OwnerId           string       `json:"owner_id"`
	Region            string       `json:"region"`
	AfkChannelId      string       `json:"afk_channel_id,omitempty"`
	AfkTimeout        int          `json:"afk_timeout"`
	WidgetEnabled     bool         `json:"widget_enabled"`
	WidgetChannelId   string       `json:"widget_channel_id,omitempty"`
	VerificationLevel int          `json:"verification_level"`
	//Roles []*Role `json:"roles"`
	//Emojis []*Emoji `json:"emojis"`
	Features    []string `json:"features"`
	MfaLevel    int      `json:"mfa_level"`
	Unavailable bool     `json:"unavailable"`
	MemberCount int      `json:"member_count"`
	//VoiceStates []*voice.State `json:"voice_states"`
	Members  []*Member  `json:"members"`
	Channels []*Channel `json:"channels"`
}

func NewGuild(rest *rest.Client, data []byte) (*Guild, error) {
	guild := new(Guild)

	err := json.Unmarshal(data, guild)

	if err != nil {
		return nil, err
	}

	guild.Rest = rest

	return guild, nil
}
