package gateway

import "github.com/Goscord/goscord/goscord/gateway/event"

type UserUpdateHandler struct{}

func (_ *UserUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewUserUpdate(s.rest, data)

	if err != nil {
		return
	}

	s.Publish(event.EventUserUpdate, ev.Data)
}
