package model_test

import (
	"testing"

	"github.com/Meduzz/commando/model"
	"github.com/spf13/cobra"
)

func TestCommand(t *testing.T) {
	cmd := model.Command{Name: "test", Handler: func(c *cobra.Command, s []string) error { return nil }}
	if cmd.Name != "test" {
		t.Errorf("Expected command name to be 'test', but got '%s'", cmd.Name)
	}
}

func TestChildCommand(t *testing.T) {
	parent := model.Command{Name: "parent", Handler: func(c *cobra.Command, s []string) error { return nil }}
	child := parent.ChildCommand("child", func(c *cobra.Command, s []string) error { return nil })

	if len(parent.Children) != 1 || parent.Children[0].Name != "child" {
		t.Errorf("Expected child command to be added")
	}

	if child.Name != "child" {
		t.Errorf("Expected child name to be 'child', but got '%s'", child.Name)
	}

	if len(child.Children) > 0 {
		t.Errorf("Expected no children for the child command")
	}
}
