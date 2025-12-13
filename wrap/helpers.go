package wrap

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/Meduzz/commando"
	"github.com/Meduzz/commando/model"
	"github.com/Meduzz/helper/fp/result"
	"github.com/Meduzz/helper/fp/slice"
	"github.com/spf13/cobra"
)

type (
	ParamKind string

	Param interface {
		Fetch(*cobra.Command) (any, error)
		Kind() ParamKind
	}

	Strategy interface {
		Read([]byte) (any, error)
		Write(any) ([]byte, error)
	}

	flag struct {
		model.Flag
	}

	body struct {
		strategy Strategy
	}

	errror struct{}

	jsonStrategy struct {
		factory func() any
	}

	stringStrategy struct{}
)

var (
	FLAG  = ParamKind("flag")
	BODY  = ParamKind("body")
	ERROR = ParamKind("error")
)

func Flag(name string, kind model.FlagKind, value any, description string) Param {
	return &flag{
		Flag: model.Flag{
			Name:    name,
			Kind:    kind,
			Default: value,
		},
	}
}

func Body(strategy Strategy) Param {
	return &body{
		strategy: strategy,
	}
}

func Error() Param {
	return &errror{}
}

func In(params ...Param) []Param {
	return params
}

func Out(params ...Param) []Param {
	return params
}

func Wrap(name string, delegate any, in []Param, out []Param) *model.Command {
	cmd := &model.Command{}
	cmd.Name = name

	handlerValue := reflect.ValueOf(delegate)

	if handlerValue.Kind() != reflect.Func {
		panic("delegate is not a function")
	}

	cmd.Name = name
	cmd.Handler = func(c *cobra.Command, args []string) error {
		maybePS := slice.Fold(in, &result.Operation[[]reflect.Value]{}, func(p Param, agg *result.Operation[[]reflect.Value]) *result.Operation[[]reflect.Value] {
			return result.Then(agg, func(it []reflect.Value) ([]reflect.Value, error) {
				v, err := p.Fetch(c)

				if err != nil {
					return it, err
				}

				return append(it, reflect.ValueOf(v)), nil
			})
		})

		ps, err := maybePS.Get()

		if err != nil {
			return err
		}

		rs := handlerValue.Call(ps)

		for i, p := range out {
			if p.Kind() == BODY {
				exe, ok := p.(*body)
				actual := rs[i].Interface()

				if !ok {
					return fmt.Errorf("could not cast param to body")
				}

				bs, err := exe.strategy.Write(actual)

				if err != nil {
					return err
				}

				_, err = os.Stdout.Write(bs)

				if err != nil {
					return err
				}
			} else if p.Kind() == ERROR {
				actual := rs[i]

				if !actual.IsNil() {
					err, ok := actual.Interface().(error)

					if ok {
						return err
					}
					// else?
				}
			}
		}

		return nil
	}

	slice.ForEach(slice.Filter(in, func(p Param) bool {
		return p.Kind() == FLAG
	}), func(p Param) {
		// only look at in params
		flag, ok := p.(*flag)

		if ok {
			cmd.AddFlag(&flag.Flag)
		}
	})

	return cmd
}

func (f *flag) Fetch(cmd *cobra.Command) (any, error) {
	visitor := commando.VisitorByKind(f.Flag.Kind)

	if visitor != nil {
		return visitor.Runtime(f.Flag.Name, cmd)
	}

	return nil, fmt.Errorf("unknown FlagKind: %s", f.Flag.Kind)
}

func (f *flag) Kind() ParamKind {
	return FLAG
}

func (b *body) Fetch(cmd *cobra.Command) (any, error) {
	bs, err := io.ReadAll(os.Stdin)

	if err != nil {
		return nil, err
	}

	return b.strategy.Read(bs)
}

func (b *body) Kind() ParamKind {
	return BODY
}

func (e *errror) Fetch(cmd *cobra.Command) (any, error) {
	return nil, fmt.Errorf("fetch is not implemented for error")
}

func (e *errror) Kind() ParamKind {
	return ERROR
}

func Json[T any]() Strategy {
	return &jsonStrategy{
		factory: func() any { return new(T) },
	}
}

func (j *jsonStrategy) Read(bs []byte) (any, error) {
	target := j.factory()
	err := json.Unmarshal(bs, target)

	return target, err
}

func (j *jsonStrategy) Write(target any) ([]byte, error) {
	return json.Marshal(target)
}

func String() Strategy {
	return &stringStrategy{}
}

func (s *stringStrategy) Read(bs []byte) (any, error) {
	return string(bs), nil
}

func (s *stringStrategy) Write(target any) ([]byte, error) {
	it, ok := target.(string)

	if !ok {
		return nil, fmt.Errorf("not a string")
	}

	return []byte(it), nil
}
