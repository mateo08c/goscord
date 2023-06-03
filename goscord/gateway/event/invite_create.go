package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
	"log"
)

type InviteCreate struct {
	Data *discord.Invite `json:"d"`
}

func NewInviteCreate(rest *rest.Client, data []byte) (*InviteCreate, error) {
	pk := new(InviteCreate)

	log.Println(string(data))

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
