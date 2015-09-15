package service

import (
	"github.com/nlopes/slack"
)

type PingCommand struct {
	AbstractSlackCommand
}

func (c *PingCommand) GetName() string {
	return "ping"
}

func (c *PingCommand) GetDescription() string {
	return "Простая команда ping-pong"
}

func (c *PingCommand) Run(m *slack.MessageEvent, args ...string) {
	c.SendMessagef(m.Channel, "Pong. <@%s>: твой ход!", m.User)
}
