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
var converters = make([]flags.Converter, 0)

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

func RegisterConverter(converter flags.Converter) {
	converters = append(converters, converter)
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
	c.RunE = handlerWrapper(cmd)

	errorz := slice.Map(cmd.Flags, func(f *model.Flag) error {
		matches := slice.Filter(converters, func(c flags.Converter) bool {
			return c.Kind() == f.Kind
		})

		match := slice.Head(matches)

		var err error

		if match != nil {
			err = match.Convert(f, c.Flags())
		} else {
			err = fmt.Errorf("no flag converter found for FlagKind %s (%s --%s) was found", f.Kind, cmd.Name, f.Name)
		}

		return err
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

func handlerWrapper(cmd *model.Command) func(*cobra.Command, []string) error {
	var handler model.Handler = cmd.Handler

	if cmd.HandlerRef != "" {
		matches := slice.Filter(handlers, func(h *HandlerSpec) bool {
			return h.Name == cmd.HandlerRef
		})

		match := slice.Head(matches)

		if match == nil {
			return func(c *cobra.Command, s []string) error {
				return fmt.Errorf("unkown handler named '%s'", cmd.HandlerRef)
			}
		} else {
			handler = match.Handler
		}
	}

	if handler == nil {
		return func(c *cobra.Command, s []string) error {
			return fmt.Errorf("no handler defined for command '%s'", cmd.Name)
		}
	}

	return func(c *cobra.Command, s []string) error {
		cmd := model.NewExecuteCommand(cmd, c, s)
		return handler(cmd)
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
