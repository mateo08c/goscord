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
	Guild  *Guild
}

func (m *GuildMember) AddRole(roleId string) error {
	err := m.Client.Rest.Guild.AddMemberRole(m.GuildId, m.User.Id, roleId)
	if err != nil {
		return err
	}

	return nil
}

func (m *GuildMember) RemoveRole(s string) error {
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

		return &GuildMember{GuildMember: get, Client: g.Client, Guild: g}, nil
	}

	return &GuildMember{GuildMember: member, Client: g.Client, Guild: g}, nil
}

func (g *Guild) GetChannel(channelId string) (*Channel, error) {
	channel, err := g.Client.Session.State().Channel(channelId)
	if err != nil {
		get, err := g.Client.Rest.Channel.Get(channelId)
		if err != nil {
			return nil, err
		}

		return &Channel{Channel: get, Client: g.Client, Guild: g}, nil
	}

	return &Channel{Channel: channel, Client: g.Client, Guild: g}, nil
}
