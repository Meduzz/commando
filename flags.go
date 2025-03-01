package commando

import (
	"github.com/Meduzz/commando/flags"
	"github.com/Meduzz/commando/model"
	"github.com/spf13/pflag"
)

type (
	flagConverter struct {
		kind    model.FlagKind
		handler func(*model.Flag, *pflag.FlagSet) error
	}
)

func NewConverter(kind model.FlagKind, handler func(*model.Flag, *pflag.FlagSet) error) flags.Converter {
	c := &flagConverter{
		kind:    kind,
		handler: handler,
	}

	RegisterConverter(c)

	return c
}

func (c *flagConverter) Kind() model.FlagKind {
	return c.kind
}

func (c *flagConverter) Convert(flag *model.Flag, flagSet *pflag.FlagSet) error {
	return c.handler(flag, flagSet)
}
