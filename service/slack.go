package service

func (s *SlackService) GetSlackCommands() []SlackCommand {
	return []SlackCommand{
		&ConfigCommand{},
		&HelloCommand{},
		&HelpCommand{},
		&LogCommand{},
		&PingCommand{},
		&ServicesCommand{},
		&UnknownCommand{},
		&VersionCommand{},
	}
}
