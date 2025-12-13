package commando

import (
	"testing"

	"github.com/Meduzz/commando/flags"
	"github.com/Meduzz/commando/model"
	"github.com/spf13/cobra"
)

func TestRegisterCommand(t *testing.T) {
	cmd := &model.Command{Name: "test", Handler: func(c *cobra.Command, s []string) error { return nil }}
	RegisterCommand(cmd)

	if len(commands) != 1 || commands[0].Name != "test" {
		t.Errorf("Expected command to be registered")
	}
}

func TestRegisterHandler(t *testing.T) {
	RegisterHandler("test", func(c *cobra.Command, s []string) error { return nil })

	if len(handlers) != 1 || handlers[0].Name != "test" {
		t.Errorf("Expected handler to be registered")
	}
}

func TestRegisterCommandWithFlags(t *testing.T) {
	cmd := CommandRef("test", "test")
	flag1 := flags.StringFlag("flag1", "default1", "flag1 description")
	cmd.AddFlag(flag1)

	if cmd.Flags[0].Name != "flag1" {
		t.Errorf("Expected flag to be registered")
	}
}

func TestAsCobra(t *testing.T) {
	cmd := CommandRef("test", "asdf")
	cmd.ChildCommandRef("child1", "asdf")
	intFlag := flags.IntFlag("int", 1, "int flag descriptioin")
	cmd.AddFlag(intFlag)
	cmd.WithDescription("Test command")

	c, err := asCobra(cmd)

	if err != nil {
		t.Error(err)
	}

	if c.Use != "test" || c.Short != "Test command" {
		t.Errorf("Expected Use to be 'test' and Short to be 'Test command', but got %s and %s", c.Use, c.Short)
	}

	flag := c.Flag("int")

	if flag == nil || flag.DefValue != "1" {
		t.Errorf("Expected flag to be registered with correct default value")
	}
}
