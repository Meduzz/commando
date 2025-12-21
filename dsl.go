package commando

import (
	"github.com/Meduzz/commando/builder"
	"github.com/Meduzz/commando/delegate"
	"github.com/Meduzz/commando/model"
)

// build a command (name) from a handlerFunc (handler)
func Command(name string, handler model.Handler) *model.Command {
	return model.CreateCommand(name, handler)
}

// build a command (name) from handlerRef (handlerRef)
func CommandRef(name string, handlerRef string) *model.Command {
	return model.CreateCommandRef(name, handlerRef)
}

// build a command (name) via builder (bldr)
func CommandBuilder(name string, bldr builder.CommandBuilderFunc) *model.Command {
	return builder.Command(name, bldr)
}

// build a handlerRef from a normal func (fun), described via builder (bldr)
func HandlerRefBuilder(fun any, bldr builder.HandlerRefBuilderFunc) *model.HandlerRef {
	return builder.HandlerRef(fun, bldr)
}

// build a command (name) around a normal func (fun) by describing its params (in) and returns (out).
func DelegateCommand(name string, fun any, in []model.Param, out []model.Param) *model.Command {
	return delegate.DelegateCommand(name, fun, in, out)
}
