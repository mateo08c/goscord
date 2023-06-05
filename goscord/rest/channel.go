package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/discord/builder"
	"github.com/bytedance/sonic"
	"io"
	"mime/multipart"
	"net/http"
)

type ChannelHandler struct {
	rest *Client
}

func NewChannelHandler(rest *Client) *ChannelHandler {
	return &ChannelHandler{rest: rest}
}

// Reactions

func (ch *ChannelHandler) GetReactions(channelId, messageId, emoji string) ([]*discord.User, error) {
	data, err := ch.rest.Request(fmt.Sprintf(EndpointGetReactions, channelId, messageId, emoji), http.MethodGet, nil, "application/json")
	if err != nil {
		return nil, err
	}

	var users []*discord.User
	err = sonic.Unmarshal(data, &users)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ch *ChannelHandler) AddReaction(channelId, messageId, emoji string) error {
	_, err := ch.rest.Request(fmt.Sprintf(EndpointOwnReaction, channelId, messageId, emoji), http.MethodPut, nil, "application/json")

	if err != nil {
		return err
	}

	return nil
}

func (ch *ChannelHandler) DeleteOwnReaction(channelId, messageId, emoji string) error {
	_, err := ch.rest.Request(fmt.Sprintf(EndpointOwnReaction, channelId, messageId, emoji), http.MethodDelete, nil, "application/json")

	if err != nil {
		return err
	}

	return nil
}

func (ch *ChannelHandler) DeleteUserReaction(channelId, messageId, emoji, userId string) error {
	_, err := ch.rest.Request(fmt.Sprintf(EndpointDeleteUserReaction, channelId, messageId, emoji, userId), http.MethodDelete, nil, "application/json")

	if err != nil {
		return err
	}

	return nil
}

func (ch *ChannelHandler) DeleteAllReactions(channelId, messageId string) error {
	_, err := ch.rest.Request(fmt.Sprintf(EndpointDeleteAllReactions, channelId, messageId), http.MethodDelete, nil, "application/json")

	if err != nil {
		return err
	}

	return nil
}

func (ch *ChannelHandler) DeleteAllReactionsForEmoji(channelId, messageId, emoji string) error {
	_, err := ch.rest.Request(fmt.Sprintf(EndpointDeleteAllReactionsForEmoji, channelId, messageId, emoji), http.MethodDelete, nil, "application/json")

	if err != nil {
		return err
	}

	return nil
}

func (ch *ChannelHandler) Get(channelId string) (*discord.Channel, error) {
	data, err := ch.rest.Request(fmt.Sprintf(EndpointGetChannel, channelId), http.MethodGet, nil, "application/json")

	if err != nil {
		return nil, err
	}

	var channel *discord.Channel
	err = sonic.Unmarshal(data, &channel)

	if err != nil {
		return nil, err
	}

	return channel, nil
}

// GetMessage gets a message from a channel
func (ch *ChannelHandler) GetMessage(channelId, messageId string) (*discord.Message, error) {
	res, err := ch.rest.Request(fmt.Sprintf(EndpointGetChannelMessage, channelId, messageId), http.MethodGet, nil, "application/json")

	if err != nil {
		return nil, err
	}

	msg := new(discord.Message)

	err = sonic.Unmarshal(res, msg)

	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (ch *ChannelHandler) SendMessage(channelId string, content any) (*discord.Message, error) {
	b, contentType, err := formatMessage(content, "")

	if err != nil {
		return nil, err
	}

	res, err := ch.rest.Request(fmt.Sprintf(EndpointCreateMessage, channelId), http.MethodPost, b, contentType)

	if err != nil {
		return nil, err
	}

	msg := new(discord.Message)

	err = sonic.Unmarshal(res, msg)

	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (ch *ChannelHandler) ReplyMessage(channelId, messageId string, content any) (*discord.Message, error) {
	b, contentType, err := formatMessage(content, messageId)

	if err != nil {
		return nil, err
	}

	res, err := ch.rest.Request(fmt.Sprintf(EndpointCreateMessage, channelId), http.MethodPost, b, contentType)

	if err != nil {
		return nil, err
	}

	msg := new(discord.Message)

	err = sonic.Unmarshal(res, msg)

	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (ch *ChannelHandler) EditMessage(channelId, messageId string, content any) (*discord.Message, error) {
	b, contentType, err := formatMessage(content, "")

	if err != nil {
		return nil, err
	}

	res, err := ch.rest.Request(fmt.Sprintf(EndpointEditMessage, channelId, messageId), http.MethodPatch, b, contentType)

	if err != nil {
		return nil, err
	}

	msg := new(discord.Message)

	err = sonic.Unmarshal(res, msg)

	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (ch *ChannelHandler) CrosspostMessage(channelId, messageId string) (*discord.Message, error) {
	data, err := ch.rest.Request(fmt.Sprintf(EndpointCrosspostMessage, channelId, messageId), http.MethodPost, nil, "application/json")

	if err != nil {
		return nil, err
	}

	var msg discord.Message
	err = sonic.Unmarshal(data, &msg)

	if err != nil {
		return nil, err
	}

	return &msg, nil
}

func (ch *ChannelHandler) DeleteMessage(channelId string, messageId string) error {
	_, err := ch.rest.Request(fmt.Sprintf(EndpointDeleteMessage, channelId, messageId), http.MethodDelete, nil, "application/json")

	if err != nil {
		return err
	}

	return nil
}

// formatMessage formats the message to be sent to the API it avoids code duplication. // ToDo : Create a custom type for it, use generics when available
func formatMessage(content any, messageId string) (*bytes.Buffer, string, error) {
	b := new(bytes.Buffer)
	contentType := "application/json"

	switch ccontent := content.(type) {
	case string:
		if messageId != "" {
			content = &discord.Message{Content: ccontent, MessageReference: &discord.MessageReference{MessageId: messageId}}
		} else {
			content = &discord.Message{Content: ccontent}
		}

		jsonb, err := json.Marshal(content)

		if err != nil {
			return nil, "", err
		}

		b = bytes.NewBuffer(jsonb)

	case *discord.Embed:
		if messageId != "" {
			content = &discord.Message{Embeds: []*discord.Embed{ccontent}, MessageReference: &discord.MessageReference{MessageId: messageId}}
		} else {
			content = &discord.Message{Embeds: []*discord.Embed{ccontent}}
		}

		jsonb, err := json.Marshal(content)

		if err != nil {
			return nil, "", err
		}

		b = bytes.NewBuffer(jsonb)

	case *discord.Message:
		jsonb, err := json.Marshal(content)

		if err != nil {
			return nil, "", err
		}

		b = bytes.NewBuffer(jsonb)

	case *builder.MessageBuilder:
		w := multipart.NewWriter(b)

		for i, file := range ccontent.Files() {
			fw, err := w.CreateFormFile(fmt.Sprintf("attachment-%d", i), file.Name)
			if err != nil {
				return nil, "", err
			}

			_, err = io.Copy(fw, file.Reader)
			if err != nil {
				return nil, "", err
			}
		}

		jsonb, err := json.Marshal(ccontent.Build())
		if err != nil {
			return nil, "", err
		}

		fw, err := w.CreateFormField("payload_json")
		if err != nil {
			return nil, "", err
		}

		_, err = fw.Write(jsonb)
		if err != nil {
			return nil, "", err
		}

		w.Close()

		contentType = w.FormDataContentType()

	default:
		return nil, "", errors.New("invalid content type, must be string, *builder.Embed, *discord.Message, *builder.MessageBuilder")
	}

	return b, contentType, nil
}

// TODO
