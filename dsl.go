package commando

import (
	"github.com/Meduzz/commando/model"
)

func Command(name string, handler model.Handler) *model.Command {
	cmd := &model.Command{}

	cmd.Name = name
	cmd.Handler = handler

	RegisterCommand(cmd)

	return cmd
}

func CommandRef(name string, handlerRef string) *model.Command {
	cmd := &model.Command{}

	cmd.Name = name
	cmd.HandlerRef = handlerRef

	RegisterCommand(cmd)

	return cmd
}
