package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

type InviteDelete struct {
	Data *discord.Invite `json:"d"`
}

func NewInviteDelete(rest *rest.Client, data []byte) (*InviteDelete, error) {
	pk := new(InviteDelete)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
