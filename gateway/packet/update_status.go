package packet

import (
	"github.com/Goscord/goscord/discord"
)

type UpdateStatus struct {
	Packet
	Data struct {
		Since  int                `json:"since"`
		Game   *discord.Activity  `json:"game"`
		Status discord.StatusType `json:"status"`
		AFK    bool               `json:"afk"`
	} `json:"d"`
}

func NewUpdateStatus(game *discord.Activity, status discord.StatusType) *UpdateStatus {
	update := &UpdateStatus{}

	update.Opcode = OpUpdateStatus
	update.Data.Game = game
	update.Data.Status = status

	return update
}
