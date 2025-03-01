package flags

import (
	"github.com/Meduzz/commando/model"
	"github.com/spf13/pflag"
)

type (
	Converter interface {
		Kind() model.FlagKind
		Convert(*model.Flag, *pflag.FlagSet) error
	}
)
