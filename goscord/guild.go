package goscord

import (
	"github.com/Goscord/goscord/goscord/discord"
)

type Guild struct {
	*discord.Guild
	Client *Client
}

type GuildMember struct {
	*discord.GuildMember
	Client *Client
}

func (m GuildMember) AddRole(roleId string) error {
	err := m.Client.Rest.Guild.AddMemberRole(m.GuildId, m.User.Id, roleId)
	if err != nil {
		return err
	}

	return nil
}

func (m GuildMember) RemoveRole(s string) error {
	err := m.Client.Rest.Guild.RemoveMemberRole(m.GuildId, m.User.Id, s)
	if err != nil {
		return err
	}

	return nil
}

func (g *Guild) GetMember(memberId string) (*GuildMember, error) {
	member, err := g.Client.Session.State().Member(g.Id, memberId)
	if err != nil {
		get, err := g.Client.Rest.Guild.GetMember(g.Id, memberId)
		if err != nil {
			return nil, err
		}

		return &GuildMember{GuildMember: get, Client: g.Client}, nil
	}

	return &GuildMember{GuildMember: member, Client: g.Client}, nil
}
