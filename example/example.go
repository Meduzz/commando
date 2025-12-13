package main

import (
	"github.com/Meduzz/commando"
	"github.com/Meduzz/commando/flags"
	"github.com/Meduzz/commando/model"
	"github.com/spf13/cobra"
)

func init() {
	cmd := commando.Command("hello", handler)
	cmd.ChildCommand("world", handler)
	phrase := flags.StringFlag("phrase", "world", "phrase to use in hello greeting")
	cmd.AddFlag(phrase)
	cmd.Flag("counter", model.FlagIntKind, 0, "a counter of greetings")
	cmd.WithDescription("Im a static command with a child")

	dynamic := &model.Command{
		Name:        "dynamic",
		Description: "Im a dynamic command",
		HandlerRef:  "super-advanced",
		Flags: []*model.Flag{
			{Name: "phrase", Description: "A phrase to print", Kind: model.FlagStringKind, Default: "world"},
		},
	}

	commando.RegisterCommand(dynamic)
	commando.RegisterHandler("super-advanced", handler)
}

func main() {
	err := commando.Execute()

	if err != nil {
		panic(err)
	}
}

func handler(cmd *cobra.Command, args []string) error {
	switch cmd.Use {
	case "hello":
		phrase, err := cmd.Flags().GetString("phrase")

		if err != nil {
			return err
		}

		count, err := cmd.Flags().GetInt("counter")

		if err != nil {
			return err
		}

		println("hello", phrase)
		println("highly manual counter:", count)
	case "world":
		println("hello", cmd.Use)
	case "dynamic":
		phrase, err := cmd.Flags().GetString("phrase")

		if err != nil {
			return err
		}

		println("hello", phrase)
	}

	return nil
}
