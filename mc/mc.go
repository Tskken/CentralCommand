package mc

import "github.com/Tskana/CentralCommand/core"

type MessageCommand struct {
	*core.Command

	MessageOptions map[string]core.HandlerFunction
}

func (m *MessageCommand) Init(command *core.Command) map[string]core.HandlerFunction {
	m.Command = command

	m.MessageOptions = make(map[string]core.HandlerFunction)

	m.MessageOptions["-delete"] = m.DeleteMessageHandler
	m.MessageOptions["-del"] = m.DeleteMessageHandler

	return m.MessageOptions
}
