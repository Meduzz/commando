package registry

import (
	"github.com/Meduzz/commando/flags"
	"github.com/Meduzz/commando/model"
	"github.com/Meduzz/helper/fp/slice"
	"github.com/spf13/cobra"
)

var commands = make([]*model.Command, 0)
var handlers = make([]*HandlerSpec, 0)
var visitors = make([]flags.FlagVisitor, 0)
var cobras = make([]*cobra.Command, 0)

type (
	HandlerSpec struct {
		Name     string
		Handler  model.Handler
		Delegate *model.HandlerRef
	}
)

func RegisterCommand(cmd *model.Command) {
	commands = append(commands, cmd)
}

func RegisterHandler(name string, handler model.Handler) {
	handlers = append(handlers, &HandlerSpec{Name: name, Handler: handler})
}

func RegisterVisitor(visitor flags.FlagVisitor) {
	visitors = append(visitors, visitor)
}

func RegisterDelegateHandler(name string, delegate *model.HandlerRef) {
	handlers = append(handlers, &HandlerSpec{Name: name, Delegate: delegate})
}

func VisitorByKind(kind model.FlagKind) flags.FlagVisitor {
	return slice.Head(slice.Filter(visitors, func(v flags.FlagVisitor) bool {
		return v.Kind() == kind
	}))
}

func HandlerByName(name string) *HandlerSpec {
	return slice.Head(slice.Filter(handlers, func(it *HandlerSpec) bool {
		return it.Name == name
	}))
}

func Commands() []*model.Command {
	return commands
}
