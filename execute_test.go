package commando

import (
	"testing"

	"github.com/Meduzz/commando/flags"
)

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
