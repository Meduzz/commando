package builder

import "github.com/Meduzz/commando/model"

type (
	CommandBuilder interface {
		Description(string) CommandBuilder
		Handler(model.Handler) CommandBuilder
		HandlerRef(string) CommandBuilder
		SubCommand(string, BuilderFunc) CommandBuilder
		Flag(*model.Flag) CommandBuilder
		build() *model.Command
	}

	BuilderFunc func(CommandBuilder)
)
