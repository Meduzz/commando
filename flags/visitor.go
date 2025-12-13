package flags

import (
	"github.com/Meduzz/commando/model"
	"github.com/spf13/cobra"
)

type (
	flagVisitor struct {
		kind    model.FlagKind
		setup   func(*model.Flag, *cobra.Command)
		runtime func(string, *cobra.Command) (any, error)
	}
)

var (
	_ FlagVisitor = &flagVisitor{}
)

func NewVisitor(kind model.FlagKind, setup func(*model.Flag, *cobra.Command), runtime func(string, *cobra.Command) (any, error)) FlagVisitor {
	v := &flagVisitor{
		kind:    kind,
		setup:   setup,
		runtime: runtime,
	}

	return v
}

func (f *flagVisitor) Kind() model.FlagKind {
	return f.kind
}

func (f *flagVisitor) Setup(flag *model.Flag, cmd *cobra.Command) {
	f.setup(flag, cmd)
}

func (f *flagVisitor) Runtime(name string, cmd *cobra.Command) (any, error) {
	return f.runtime(name, cmd)
}
