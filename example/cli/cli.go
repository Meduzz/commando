package main

import (
	"fmt"

	"github.com/Meduzz/commando"
	"github.com/Meduzz/commando/model"
	"github.com/Meduzz/commando/wrap"
)

type (
	Greeting struct {
		Message string `json:"message"`
	}
)

func main() {
	cmd := wrap.Wrap("greet", greet, wrap.In(wrap.Flag("name", model.FlagStringKind, "", "The name to greet"), wrap.Body(wrap.Json[Greeting]())), wrap.Out(wrap.Body(wrap.Json[Greeting]()), wrap.Error()))
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
