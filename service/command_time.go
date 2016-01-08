package service

import (
	"time"

	sl "github.com/nlopes/slack"
)

type TimeCommand struct {
	AbstractSlackCommand
}

func (c *TimeCommand) GetName() string {
	return "time"
}

func (c *TimeCommand) GetDescription() string {
	return "Show server time"
}

func (c *TimeCommand) Run(m *sl.MessageEvent, args ...string) {
	c.SendMessagef(m.Channel, time.Now().Format(time.RFC3339))
}
