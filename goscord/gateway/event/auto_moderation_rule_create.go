package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

type AutoModerationRuleCreate struct {
	Data *discord.AutoModerationRule `json:"d"`
}

func NewAutoModerationRuleCreate(rest *rest.Client, data []byte) (*AutoModerationRuleCreate, error) {
	pk := new(AutoModerationRuleCreate)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
