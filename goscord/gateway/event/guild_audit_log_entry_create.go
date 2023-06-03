package event

import (
	"encoding/json"
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
)

type GuildAuditLogEntryCreate struct {
	Data *discord.GuildAuditLog `json:"d"`
}

func NewGuildAuditLogEntryCreate(rest *rest.Client, data []byte) (*GuildAuditLogEntryCreate, error) {
	pk := new(GuildAuditLogEntryCreate)
	
	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
