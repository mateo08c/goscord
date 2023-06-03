package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

type TypingStart struct {
	Data *discord.TypingStartEvent `json:"d"`
}

func NewTypingStart(rest *rest.Client, data []byte) (*TypingStart, error) {
	pk := new(TypingStart)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
