package model

func CreateCommand(name string, handler Handler) *Command {
	cmd := &Command{}

	cmd.Name = name
	cmd.Handler = handler

	return cmd
}

func CreateCommandRef(name string, handlerRef string) *Command {
	cmd := &Command{}

	cmd.Name = name
	cmd.HandlerRef = handlerRef

	return cmd
}
