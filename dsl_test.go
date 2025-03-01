package commando_test

import (
	"testing"

	"github.com/Meduzz/commando"
	"github.com/Meduzz/commando/model"
)

func TestCommand(t *testing.T) {
	cmd := commando.Command("test", func(model.ExecuteCommand) error { return nil })
	if cmd.Name != "test" {
		t.Errorf("Expected command name to be 'test', but got '%s'", cmd.Name)
	}
}
