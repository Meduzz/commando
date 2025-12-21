package commando

import (
	"errors"
	"fmt"

	"github.com/Meduzz/commando/delegate"
	"github.com/Meduzz/commando/model"
	"github.com/Meduzz/commando/registry"
	"github.com/Meduzz/helper/fp/slice"
	"github.com/spf13/cobra"
)

func Execute() error {
	root := &cobra.Command{}

	errorz := slice.Map(registry.Commands(), func(cmd *model.Command) error {
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
	handlerForCommand(cmd, c)

	errorz := slice.Map(cmd.Flags, func(f *model.Flag) error {
		visitor := registry.VisitorByKind(f.Kind)

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

func handlerForCommand(cmd *model.Command, real *cobra.Command) {
	if cmd.HandlerRef != "" {
		match := registry.HandlerByName(cmd.HandlerRef)

		if match != nil {
			if match.Handler != nil {
				cmd.Handler = match.Handler
			} else if match.Delegate != nil {
				delegate.HandlerRef(cmd, match.Delegate)
			}
		}
	}

	if cmd.Handler != nil {
		real.RunE = cmd.Handler
	}
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
