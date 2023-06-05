package rest

import (
	"fmt"
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/bytedance/sonic"
	"net/http"
)

type GuildHandler struct {
	rest *Client
}

func NewGuildHandler(rest *Client) *GuildHandler {
	return &GuildHandler{rest: rest}
}

// Members

func (gh *GuildHandler) GetMember(guildId, userId string) (*discord.GuildMember, error) {
	data, err := gh.rest.Request(fmt.Sprintf(EndpointGetGuildMember, guildId, userId), http.MethodGet, nil, "application/json")

	if err != nil {
		return nil, err
	}

	var member *discord.GuildMember
	err = sonic.Unmarshal(data, &member)

	if err != nil {
		return nil, err
	}

	return member, nil
}

// Roles

func (gh *GuildHandler) AddMemberRole(guildId, userId, roleId string) error {
	_, err := gh.rest.Request(fmt.Sprintf(EndpointAddGuildMemberRole, guildId, userId, roleId), http.MethodPut, nil, "application/json")

	if err != nil {
		return err
	}

	return nil
}

func (gh *GuildHandler) RemoveMemberRole(guildId, userId, roleId string) error {
	_, err := gh.rest.Request(fmt.Sprintf(EndpointRemoveGuildMemberRole, guildId, userId, roleId), http.MethodDelete, nil, "application/json")

	if err != nil {
		return err
	}

	return nil
}

//Guilds

func (gh *GuildHandler) Get(s discord.Snowflake) (*discord.Guild, error) {
	data, err := gh.rest.Request(fmt.Sprintf(EndpointGetGuild, s), http.MethodGet, nil, "application/json")

	if err != nil {
		return nil, err
	}

	var guild *discord.Guild
	err = sonic.Unmarshal(data, &guild)

	if err != nil {
		return nil, err
	}

	return guild, nil
}
