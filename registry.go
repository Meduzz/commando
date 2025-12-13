package commando

import (
	"errors"
	"fmt"

	"github.com/Meduzz/commando/flags"
	"github.com/Meduzz/commando/model"
	"github.com/Meduzz/helper/fp/slice"
	"github.com/spf13/cobra"
)

var commands = make([]*model.Command, 0)
var handlers = make([]*HandlerSpec, 0)
var visitors = make([]flags.FlagVisitor, 0)

type (
	HandlerSpec struct {
		Name    string
		Handler model.Handler
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

func VisitorByKind(kind model.FlagKind) flags.FlagVisitor {
	return slice.Head(slice.Filter(visitors, func(v flags.FlagVisitor) bool {
		return v.Kind() == kind
	}))
}

func Execute() error {
	root := &cobra.Command{}

	errorz := slice.Map(commands, func(cmd *model.Command) error {
		c, err := asCobra(cmd)

		if err != nil {
			return err
		}

		root.AddCommand(c)

		return nil
	})

	err := mergeErrors(errorz)

	if err != nil {
		return err
	}

	return root.Execute()
}

func asCobra(cmd *model.Command) (*cobra.Command, error) {
	c := &cobra.Command{}

	c.Use = cmd.Name
	c.Short = cmd.Description
	c.RunE = handlerForCommand(cmd)

	errorz := slice.Map(cmd.Flags, func(f *model.Flag) error {
		visitor := VisitorByKind(f.Kind)

		if visitor == nil {
			return fmt.Errorf("no visitor for flag kind: %s", f.Kind)
		}

		// TODO opportunity to return error
		visitor.Setup(f, c)

		return nil
	})

	err := mergeErrors(errorz)

	if err != nil {
		return nil, err
	}

	errorz = slice.Map(cmd.Children, func(child *model.Command) error {
		childC, err := asCobra(child)

		if err == nil {
			c.AddCommand(childC)
		}

		return err
	})

	err = mergeErrors(errorz)

	if err != nil {
		return nil, err
	}

	return c, nil
}

func handlerForCommand(cmd *model.Command) model.Handler {
	if cmd.Handler != nil {
		return cmd.Handler
	}

	if cmd.HandlerRef != "" {
		match := slice.Head(slice.Filter(handlers, func(h *HandlerSpec) bool {
			return h.Name == cmd.HandlerRef
		}))

		if match != nil {
			return match.Handler
		}
	}

	return nil
}

func mergeErrors(errorz []error) error {
	return slice.Fold(errorz, nil, func(e error, agg error) error {
		if e == nil {
			return agg
		}

		if agg == nil {
			return e
		}

		return errors.Join(agg, e)
	})

}
