package builder

import "github.com/Meduzz/commando/model"

type (
	commandBuilder struct {
		cmd *model.Command
	}
)

var (
	_ CommandBuilder = &commandBuilder{}
)

func Builder(name string, handler BuilderFunc) *model.Command {
	cb := NewCommandBuilder(name)
	handler(cb)

	return cb.build()
}

func NewCommandBuilder(name string) CommandBuilder {
	return &commandBuilder{
		cmd: &model.Command{
			Name: name,
		},
	}
}

func (c *commandBuilder) Description(description string) CommandBuilder {
	c.cmd.Description = description
	return c
}

func (c *commandBuilder) Handler(handler model.Handler) CommandBuilder {
	c.cmd.Handler = handler
	return c
}

func (c *commandBuilder) HandlerRef(handlerRef string) CommandBuilder {
	c.cmd.HandlerRef = handlerRef
	return c
}

func (c *commandBuilder) SubCommand(name string, builder BuilderFunc) CommandBuilder {
	cmd := NewCommandBuilder(name)
	builder(cmd)
	c.cmd.Children = append(c.cmd.Children, cmd.build())

	return c
}

func (c *commandBuilder) Flag(flag *model.Flag) CommandBuilder {
	c.cmd.AddFlag(flag)
	return c
}

func (c *commandBuilder) build() *model.Command {
	return c.cmd
}
