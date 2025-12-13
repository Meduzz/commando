package commando

import (
	"github.com/Meduzz/commando/builder"
	"github.com/Meduzz/commando/model"
)

func Command(name string, handler model.Handler) *model.Command {
	cmd := &model.Command{}

	cmd.Name = name
	cmd.Handler = handler

	return cmd
}

func CommandRef(name string, handlerRef string) *model.Command {
	cmd := &model.Command{}

	cmd.Name = name
	cmd.HandlerRef = handlerRef

	return cmd
}

func CommandBuilder(name string, handler builder.BuilderFunc) *model.Command {
	return builder.Builder(name, handler)
}
