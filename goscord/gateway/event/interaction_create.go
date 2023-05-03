package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

type InteractionCreate struct {
	Data *discord.Interaction `json:"d"`
}

func NewInteractionCreate(rest *rest.Client, data []byte) (*InteractionCreate, error) {
	pk := new(InteractionCreate)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
