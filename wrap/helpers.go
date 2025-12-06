package wrap

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/Meduzz/commando/model"
	"github.com/Meduzz/helper/fp/result"
	"github.com/Meduzz/helper/fp/slice"
)

type (
	Param interface {
		Fetch(model.ExecuteCommand) (reflect.Value, error)
		Kind() string
	}

	Strategy interface {
		Read([]byte) (any, error)
		Write(any) ([]byte, error)
	}

	flag struct {
		name         string
		kind         model.FlagKind
		defaultValue any
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
	blank = reflect.ValueOf(nil)
)

func Flag(name string, kind model.FlagKind, value any) Param {
	return &flag{
		name:         name,
		kind:         kind,
		defaultValue: value,
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

	// TODO assert it's a func
	handlerValue := reflect.ValueOf(delegate)

	cmd.Name = name
	cmd.Handler = func(ec model.ExecuteCommand) error {
		maybePS := slice.Fold(in, &result.Operation[[]reflect.Value]{}, func(p Param, agg *result.Operation[[]reflect.Value]) *result.Operation[[]reflect.Value] {
			return result.Then(agg, func(it []reflect.Value) ([]reflect.Value, error) {
				v, err := p.Fetch(ec)

				if err != nil {
					return it, err
				}

				return append(it, v), nil
			})
		})

		ps, err := maybePS.Get()

		if err != nil {
			return err
		}

		rs := handlerValue.Call(ps)

		for i, p := range out {
			if p.Kind() == "body" {
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
			} else if p.Kind() == "error" {
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

	slice.ForEach(in, func(p Param) {
		// only look at in params
		flag, ok := p.(*flag)

		if ok {
			cmd.Flag(flag.name, flag.kind, flag.defaultValue, "")
		}
	})

	return cmd
}

func (f *flag) Fetch(cmd model.ExecuteCommand) (reflect.Value, error) {
	switch f.kind {
	case model.FlagInt64Kind:
		value, err := cmd.Int64(f.name)

		if err != nil {
			return blank, err
		}

		return reflect.ValueOf(value), nil
	case model.FlagIntKind:
		value, err := cmd.Int(f.name)

		if err != nil {
			return blank, err
		}

		return reflect.ValueOf(value), nil
	case model.FlagBoolKind:
		value, err := cmd.Bool(f.name)

		if err != nil {
			return blank, err
		}

		return reflect.ValueOf(value), nil
	case model.FlagStringKind:
		value, err := cmd.String(f.name)

		if err != nil {
			return blank, err
		}

		return reflect.ValueOf(value), nil
	}

	return blank, fmt.Errorf("unknown flagkind: %s", f.kind)
}

func (f *flag) Kind() string {
	return "flag"
}

func (b *body) Fetch(cmd model.ExecuteCommand) (reflect.Value, error) {
	bs, err := io.ReadAll(os.Stdin)

	if err != nil {
		return blank, err
	}

	it, err := b.strategy.Read(bs)

	if err != nil {
		return blank, err
	}

	return reflect.ValueOf(it), nil
}

func (b *body) Kind() string {
	return "body"
}

func (e *errror) Fetch(cmd model.ExecuteCommand) (reflect.Value, error) {
	return blank, fmt.Errorf("fetch is not implemented for error")
}

func (e *errror) Kind() string {
	return "error"
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
