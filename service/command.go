package service

import (
	"fmt"

	"github.com/kihamo/shadow"
	"github.com/nlopes/slack"
)

type SlackCommand interface {
	Init(shadow.Service, *shadow.Application)
	GetName() string
	GetDescription() string
	Run(*slack.MessageEvent, ...string)
	AllowDirectMessage() bool
	AllowChannel() bool
	IsActive() bool
}

// TODO: flags command, see service/api/command_api.go

type AbstractSlackCommand struct {
	SlackCommand
	Application  *shadow.Application
	Service      shadow.Service
	SlackService *SlackService
}

func (c *AbstractSlackCommand) AllowDirectMessage() bool {
	return true
}

func (c *AbstractSlackCommand) AllowChannel() bool {
	return true
}

func (c *AbstractSlackCommand) IsActive() bool {
	return true
}

func (c *AbstractSlackCommand) Init(s shadow.Service, a *shadow.Application) {
	c.Application = a
	c.Service = s

	slackService, err := a.GetService("slack")
	if err == nil {
		if castService, ok := slackService.(*SlackService); ok {
			c.SlackService = castService
			return
		}
	}

	panic("Slack service not found")
}

func (c *AbstractSlackCommand) SendMessage(channelId string, message string) {
	c.SlackService.Rtm.SendMessage(c.SlackService.Rtm.NewOutgoingMessage(message, channelId))
}

func (c *AbstractSlackCommand) SendMessagef(channelId string, message string, args ...interface{}) {
	message = fmt.Sprintf(message, args...)
	c.SlackService.Rtm.SendMessage(c.SlackService.Rtm.NewOutgoingMessage(message, channelId))
}

func (c *AbstractSlackCommand) SendPostMessage(channelId string, message string, params slack.PostMessageParameters) error {
	params.AsUser = true

	_, _, err := c.SlackService.Rtm.Client.PostMessage(channelId, message, params)
	return err
}
