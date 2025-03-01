package model

import (
	"fmt"

	"github.com/spf13/cobra"
)

type (
	ExecuteCommand interface {
		ExecuteFlags
		Name() string
		Arg(int) (string, error)
	}

	ExecuteFlags interface {
		Int(string) (int, error)
		Int64(string) (int64, error)
		String(string) (string, error)
		Bool(string) (bool, error)
	}

	command struct {
		design *Command
		impl   *cobra.Command
		args   []string
	}
)

func NewExecuteCommand(design *Command, impl *cobra.Command, args []string) ExecuteCommand {
	return &command{design: design, impl: impl, args: args}
}

func (c *command) Name() string {
	return c.design.Name
}

func (c *command) Arg(index int) (string, error) {
	if index >= len(c.args) {
		return "", fmt.Errorf("argument index out of range")
	}
	return c.args[index], nil
}

func (c *command) Int(name string) (int, error) {
	return c.impl.Flags().GetInt(name)
}

func (c *command) String(name string) (string, error) {
	return c.impl.Flags().GetString(name)
}

func (c *command) Bool(name string) (bool, error) {
	return c.impl.Flags().GetBool(name)
}

func (c *command) Int64(name string) (int64, error) {
	return c.impl.Flags().GetInt64(name)
}
