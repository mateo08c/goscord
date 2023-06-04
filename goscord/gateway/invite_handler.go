package gateway

import (
	"github.com/Goscord/goscord/goscord/gateway/event"
)

type InviteCreateHandler struct{}

func (_ *InviteCreateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewInviteCreate(s.rest, data)

	if err != nil {
		return
	}

	s.Publish(event.EventInviteCreate, ev.Data)
}

type InviteDeleteHandler struct{}

func (_ *InviteDeleteHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewInviteDelete(s.rest, data)

	if err != nil {
		return
	}

	s.Publish(event.EventInviteDelete, ev.Data)
}
