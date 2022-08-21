package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/discord/embed"
)

type ChannelHandler struct {
	rest *Client
}

func NewChannelHandler(rest *Client) *ChannelHandler {
	return &ChannelHandler{rest: rest}
}

func (ch *ChannelHandler) GetChannel(channelId string) (*discord.Channel, error) {
	data, err := ch.rest.Request(fmt.Sprintf(EndpointGetChannel, channelId), "GET", nil, "application/json")

	if err != nil {
		return nil, err
	}

	var channel discord.Channel
	err = json.Unmarshal(data, &channel)

	if err != nil {
		return nil, err
	}

	return &channel, nil
}

// Send / reply to a message
func (ch *ChannelHandler) GetMessage(channelId, messageId string) (*discord.Message, error) {
	data, err := ch.rest.Request(fmt.Sprintf(EndpointGetChannelMessage, channelId, messageId), "GET", nil, "application/json")

	if err != nil {
		return nil, err
	}

	var msg discord.Message
	err = json.Unmarshal(data, &msg)

	if err != nil {
		return nil, err
	}

	return &msg, nil
}

func (ch *ChannelHandler) SendMessage(channelId string, content interface{}) (*discord.Message, error) {
	b := new(bytes.Buffer)
	contentType := "application/json"

	switch ccontent := content.(type) {
	case string:
		content = &discord.Message{Content: ccontent}
		jsonb, err := json.Marshal(content)

		if err != nil {
			return nil, err
		}

		b = bytes.NewBuffer(jsonb)

	case *embed.Embed:
		content = &discord.Message{Embeds: []*embed.Embed{ccontent}}
		jsonb, err := json.Marshal(content)

		if err != nil {
			return nil, err
		}

		b = bytes.NewBuffer(jsonb)

	case []*os.File:
		w := multipart.NewWriter(b)

		for i, file := range ccontent {
			fw, err := w.CreateFormFile("attachment"+string(rune(i)), file.Name())
			if err != nil {
				return nil, err
			}

			_, err = io.Copy(fw, file)
			if err != nil {
				return nil, err
			}
		}

		w.Close()
		contentType = w.FormDataContentType()

	default:
		return nil, errors.New("bad content type")
	}

	res, err := ch.rest.Request(fmt.Sprintf(EndpointCreateMessage, channelId), "POST", b, contentType)

	if err != nil {
		return nil, err
	}

	msg := new(discord.Message)

	err = json.Unmarshal(res, msg)

	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (ch *ChannelHandler) ReplyMessage(channelId, messageId string, content interface{}) (*discord.Message, error) {
	var b *bytes.Buffer
	contentType := "application/json"

	switch ccontent := content.(type) {
	case string:
		content = &discord.Message{Content: ccontent, MessageReference: &discord.MessageReference{ChannelId: channelId, MessageId: messageId}}
		jsonb, err := json.Marshal(content)

		if err != nil {
			return nil, err
		}

		b = bytes.NewBuffer(jsonb)

	case *embed.Builder:
		content = &discord.Message{Content: ccontent.Content(), Embeds: []*embed.Embed{ccontent.Embed()}, MessageReference: &discord.MessageReference{ChannelId: channelId, MessageId: messageId}}
		jsonb, err := json.Marshal(content)

		if err != nil {
			return nil, err
		}

		b = bytes.NewBuffer(jsonb)

	case *embed.Embed:
		content = &discord.Message{Embeds: []*embed.Embed{ccontent}, MessageReference: &discord.MessageReference{ChannelId: channelId, MessageId: messageId}}
		jsonb, err := json.Marshal(content)

		if err != nil {
			return nil, err
		}

		b = bytes.NewBuffer(jsonb)

	/* TODO: implement file upload
	case *os.File:
		w := multipart.NewWriter(b)

		fw, err := w.CreateFormFile("attachment", ccontent.Name())
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(fw, ccontent)
		if err != nil {
			return nil, err
		}

		w.Close()
		contentType = w.FormDataContentType()
	*/

	default:
		return nil, errors.New("bad content type")
	}

	res, err := ch.rest.Request(fmt.Sprintf(EndpointCreateMessage, channelId), "POST", b, contentType)

	if err != nil {
		return nil, err
	}

	msg := new(discord.Message)

	err = json.Unmarshal(res, msg)

	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (ch *ChannelHandler) Edit(channelId, messageId string, content interface{}) (*discord.Message, error) {
	switch ccontent := content.(type) {
	case string:
		content = &discord.Message{Content: ccontent}

	case *embed.Builder:
		content = &discord.Message{Content: ccontent.Content(), Embeds: []*embed.Embed{ccontent.Embed()}}

	case *embed.Embed:
		content = &discord.Message{Embeds: []*embed.Embed{ccontent}}
	}

	b, err := json.Marshal(content)

	if err != nil {
		return nil, err
	}

	res, err := ch.rest.Request(fmt.Sprintf(EndpointEditMessage, channelId, messageId), "PATCH", bytes.NewBuffer(b), "application/json")

	if err != nil {
		return nil, err
	}

	msg := new(discord.Message)

	err = json.Unmarshal(res, msg)

	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (ch *ChannelHandler) CrosspostMessage(channelId, messageId string) (*discord.Message, error) {
	data, err := ch.rest.Request(fmt.Sprintf(EndpointCrosspostMessage, channelId, messageId), "POST", nil, "application/json")

	if err != nil {
		return nil, err
	}

	var msg discord.Message
	err = json.Unmarshal(data, &msg)

	if err != nil {
		return nil, err
	}

	return &msg, nil
}

// TODO
