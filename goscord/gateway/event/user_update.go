package event

import (
	"encoding/json"
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
)

type UserUpdate struct {
	Data *discord.User `json:"d"`
}

func NewUserUpdate(rest *rest.Client, data []byte) (*UserUpdate, error) {
	pk := new(UserUpdate)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
