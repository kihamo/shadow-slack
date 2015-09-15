package service

import (
	"strings"

	"github.com/kihamo/shadow/service/system"
	sl "github.com/nlopes/slack"
)

type ServicesCommand struct {
	AbstractSlackCommand
}

func (c *ServicesCommand) GetName() string {
	return "services"
}

func (c *ServicesCommand) GetDescription() string {
	return "Список подключенных служб"
}

func (c *ServicesCommand) Run(m *sl.MessageEvent, args ...string) {
	service, _ := c.Application.GetService("system")
	systemService := service.(*system.SystemService)

	values := []string{}

	for _, s := range systemService.Application.GetServices() {
		values = append(values, s.GetName())
	}

	c.SendMessage(m.Channel, strings.Join(values, "\n"))
	return
}
