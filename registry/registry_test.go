package registry

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
	cmd := model.CreateCommandRef("test", "test")
	flag1 := flags.StringFlag("flag1", "default1", "flag1 description")
	cmd.AddFlag(flag1)

	if cmd.Flags[0].Name != "flag1" {
		t.Errorf("Expected flag to be registered")
	}
}
