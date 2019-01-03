package hp

import (
	"github.com/Tskana/CentralCommand/core"
)

type HelpCommand struct {
	*core.Command

	HelpOptions map[string]core.HandlerFunction
}

func (h *HelpCommand) Init(command *core.Command) map[string]core.HandlerFunction {
	h.Command = command

	return h.HelpOptions
}
