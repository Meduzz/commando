package commando_test

import (
	"testing"

	"github.com/Meduzz/commando"
	"github.com/Meduzz/commando/builder"
	"github.com/spf13/cobra"
)

func TestCommand(t *testing.T) {
	cmd := commando.Command("test", func(c *cobra.Command, s []string) error { return nil })
	if cmd.Name != "test" {
		t.Errorf("Expected command name to be 'test', but got '%s'", cmd.Name)
	}

	t.Run("DSL", func(t *testing.T) {
		subject := commando.CommandBuilder("test", func(cb builder.CommandBuilder) {
			cb.Description("One long text")
		})

		if subject.Description != "One long text" {
			t.Errorf("Description did not match expected, was: '%s'", cmd.Description)
		}
	})
}
