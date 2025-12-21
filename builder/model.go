package builder

import "github.com/Meduzz/commando/model"

type (
	CommandBuilder interface {
		Description(string) CommandBuilder
		Handler(model.Handler) CommandBuilder
		HandlerRef(string) CommandBuilder
		SubCommand(string, CommandBuilderFunc) CommandBuilder
		Flag(*model.Flag) CommandBuilder
		build() *model.Command
	}

	CommandBuilderFunc func(CommandBuilder)

	HandlerRefBuilder interface {
		In(ParamBuilderFunc) HandlerRefBuilder  // body, flag & env
		Out(ParamBuilderFunc) HandlerRefBuilder // body & error
		build() *model.HandlerRef
	}

	HandlerRefBuilderFunc func(HandlerRefBuilder)

	ParamBuilder interface {
		Flag(name string, kind model.FlagKind, defaultValue any, description string) ParamBuilder
		Body(strategy model.Strategy) ParamBuilder
		Error() ParamBuilder
		Env(name, value string) ParamBuilder
	}

	ParamBuilderFunc func(ParamBuilder)
)
