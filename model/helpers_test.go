package model_test

import (
	"testing"

	"github.com/Meduzz/commando/model"
	"github.com/spf13/cobra"
)

func TestCreateCommand(t *testing.T) {
	cmd := model.CreateCommand("test", func(c *cobra.Command, s []string) error { return nil })
	if cmd.Name != "test" {
		t.Errorf("Expected command name to be 'test', but got '%s'", cmd.Name)
	}
}
