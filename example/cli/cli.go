package main

import (
	"fmt"

	"github.com/Meduzz/commando"
	"github.com/Meduzz/commando/cli"
	"github.com/Meduzz/commando/model"
)

type (
	Greeting struct {
		Message string `json:"message"`
	}
)

func main() {
	cmd := cli.Example("greet", greet, cli.Flag("name", model.FlagStringKind, ""), cli.InBody(cli.Json[Greeting]()), cli.OutBody(cli.Json[Greeting]()), cli.Error())
	commando.RegisterCommand(cmd.WithDescription("Greet a name with the provided message."))

	err := commando.Execute()

	if err != nil {
		panic(err)
	}
}

func greet(name string, data *Greeting) (*Greeting, error) {
	resp := fmt.Sprintf(data.Message, name)

	g := &Greeting{
		Message: resp,
	}

	return g, nil
}
