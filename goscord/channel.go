package goscord

import "github.com/Goscord/goscord/goscord/discord"

type Channel struct {
	*discord.Channel
	Client *Client
	Guild  *Guild
}

type Message struct {
	*discord.Message
	Client  *Client
	Channel *Channel
}

type Reaction struct {
	*discord.Reaction
	Client  *Client
	Message *Message
}

func (c *Channel) SendMessage(content any) (*Message, error) {
	msg, err := c.Client.Rest.Channel.SendMessage(c.Id, content)
	if err != nil {
		return nil, err
	}

	return &Message{Message: msg, Client: c.Client, Channel: c}, nil
}

// Message

func (c *Channel) GetMessage(s string) (*Message, error) {
	msg, err := c.Client.Rest.Channel.GetMessage(c.Id, s)
	if err != nil {
		return nil, err
	}

	return &Message{Message: msg, Client: c.Client, Channel: c}, nil
}

func (m *Message) Delete() error {
	err := m.Client.Rest.Channel.DeleteMessage(m.ChannelId, m.Id)
	if err != nil {
		return err
	}

	return nil
}

func (m *Message) Edit(content any) (*Message, error) {
	msg, err := m.Client.Rest.Channel.EditMessage(m.ChannelId, m.Id, content)
	if err != nil {
		return nil, err
	}

	return &Message{Message: msg, Client: m.Client, Channel: m.Channel}, nil
}

// Reactions

func (m *Message) AddReaction(emojiId string) error {
	err := m.Client.Rest.Channel.AddReaction(m.ChannelId, m.Id, emojiId)
	if err != nil {
		return err
	}

	return nil
}

func (m *Message) RemoveReaction(emojiId string) error {
	err := m.Client.Rest.Channel.DeleteAllReactionsForEmoji(m.ChannelId, m.Id, emojiId)
	if err != nil {
		return err
	}

	return nil
}

func (m *Message) RemoveUserReaction(userId string, emojiId string) error {
	err := m.Client.Rest.Channel.DeleteUserReaction(m.ChannelId, m.Id, userId, emojiId)
	if err != nil {
		return err
	}

	return nil
}

func (m *Message) RemoveAllReactions() error {
	err := m.Client.Rest.Channel.DeleteAllReactions(m.ChannelId, m.Id)
	if err != nil {
		return err
	}

	return nil
}

func (m *Message) GetReactions(s string) ([]*discord.User, error) {
	users, err := m.Client.Rest.Channel.GetReactions(m.ChannelId, m.Id, s)
	if err != nil {
		return nil, err
	}

	return users, nil
}
