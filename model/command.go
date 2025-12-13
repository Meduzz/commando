package model

import "github.com/spf13/cobra"

type (
	Command struct {
		Name        string     `json:"name"`
		Description string     `json:"description,omitempty"`
		HandlerRef  string     `json:"handlerRef,omitempty"`
		Handler     Handler    `json:"-"`
		Flags       []*Flag    `json:"flags,omitempty"`
		Children    []*Command `json:"children,omitempty"`
	}

	Handler func(*cobra.Command, []string) error
)

func (c *Command) ChildCommand(name string, handler Handler) *Command {
	child := &Command{Name: name, Handler: handler}
	c.Children = append(c.Children, child)
	return child
}

func (c *Command) ChildCommandRef(name, handlerRef string) *Command {
	child := &Command{Name: name, HandlerRef: handlerRef}
	c.Children = append(c.Children, child)
	return child
}

func (c *Command) WithDescription(text string) *Command {
	c.Description = text
	return c
}

func (c *Command) AddFlag(flag *Flag) *Command {
	c.Flags = append(c.Flags, flag)
	return c
}

func (c *Command) Flag(name string, kind FlagKind, defaultValue any, description string) *Flag {
	flag := &Flag{Name: name, Kind: kind, Default: defaultValue, Description: description}
	c.Flags = append(c.Flags, flag)
	return flag
}
