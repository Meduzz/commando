package flags

import (
	"github.com/Meduzz/commando/model"
	"github.com/spf13/cobra"
)

type (
	FlagVisitor interface {
		Kind() model.FlagKind
		Setup(*model.Flag, *cobra.Command)
		Runtime(string, *cobra.Command) (any, error)
	}
)
