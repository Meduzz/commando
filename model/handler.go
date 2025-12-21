package model

import "github.com/spf13/cobra"

type (
	HandlerRef struct {
		Delegate any     // delegate func
		In       []Param // describe in params
		Out      []Param // describe out params
	}

	ParamKind string

	Param interface {
		Fetch(*cobra.Command) (any, error)
		Kind() ParamKind
	}

	Strategy interface {
		Read([]byte) (any, error)
		Write(any) ([]byte, error)
	}
)

var (
	FLAG  = ParamKind("flag")
	BODY  = ParamKind("body")
	ERROR = ParamKind("error")
	ENV   = ParamKind("env")
)
