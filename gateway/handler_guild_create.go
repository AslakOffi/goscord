package gateway

import (
	"fmt"
	"github.com/Goscord/goscord/gateway/event"
)

type GuildCreateHandler struct{}

func (h *GuildCreateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildCreate(s.rest, data)

	if err != nil {
		return
	}

	fmt.Println("Got guild create")
	fmt.Println(ev.Data)
}
