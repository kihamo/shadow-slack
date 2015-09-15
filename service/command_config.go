package service

import (
	"fmt"
	"strings"

	"github.com/kihamo/shadow/service/system"
	sl "github.com/nlopes/slack"
)

type ConfigCommand struct {
	AbstractSlackCommand
}

func (c *ConfigCommand) GetName() string {
	return "config"
}

func (c *ConfigCommand) GetDescription() string {
	return "Информация о конфигурации системы"
}

func (c *ConfigCommand) Run(m *sl.MessageEvent, args ...string) {
	service, _ := c.Application.GetService("system")
	systemService := service.(*system.SystemService)

	if len(args) == 0 {
		values := []string{}

		for name := range systemService.Config.GetAll() {
			values = append(values, fmt.Sprintf("*%s* = %s", name, fmt.Sprint(systemService.Config.Get(name))))
		}

		c.SendMessage(m.Channel, strings.Join(values, "\n"))
		return
	}

	if !systemService.Config.Has(args[0]) {
		c.SendMessagef(m.Channel, "Настройки *%s* не существует", args[0])
		return
	}

	c.SendMessagef(m.Channel, "*%s* = %s", args[0], fmt.Sprint(systemService.Config.Get(args[0])))
}
