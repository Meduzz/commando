package main

import (
	"fmt"

	"github.com/Meduzz/commando"
	"github.com/Meduzz/commando/delegate"
	"github.com/Meduzz/commando/model"
	"github.com/Meduzz/commando/registry"
)

type (
	Greeting struct {
		Message string `json:"message"`
	}
)

func main() {
	cmd := delegate.DelegateCommand("greet", greet, delegate.In(delegate.Flag("name", model.FlagStringKind, "", "The name to greet"), delegate.Body(delegate.Json[Greeting]())), delegate.Out(delegate.Body(delegate.Json[Greeting]()), delegate.Error()))
	registry.RegisterCommand(cmd.WithDescription("Greet a name with the provided message."))

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
