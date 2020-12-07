package gateway

import (
	"github.com/Goscord/goscord/gateway/event"
)

type ReadyHandler struct{}

func (h *ReadyHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewReady(data)

	if err != nil {
		return
	}

	s.user = ev.Data.User
	s.sessionID = ev.Data.SessionID

	s.Bus().Publish("ready")
}
