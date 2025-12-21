package delegate

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/Meduzz/commando/model"
	"github.com/Meduzz/commando/registry"
	"github.com/Meduzz/helper/fp/result"
	"github.com/Meduzz/helper/fp/slice"
	"github.com/Meduzz/helper/utilz"
	"github.com/spf13/cobra"
)

type (
	flag struct {
		model.Flag
	}

	body struct {
		strategy model.Strategy
	}

	errror struct{}

	jsonStrategy struct {
		factory func() any
	}

	stringStrategy struct{}

	env struct {
		name  string
		value string
	}
)

func Flag(name string, kind model.FlagKind, value any, description string) model.Param {
	return &flag{
		Flag: model.Flag{
			Name:    name,
			Kind:    kind,
			Default: value,
		},
	}
}

func Env(name, value string) model.Param {
	return &env{
		name:  name,
		value: value,
	}
}

func Body(strategy model.Strategy) model.Param {
	return &body{
		strategy: strategy,
	}
}

func Error() model.Param {
	return &errror{}
}

func In(params ...model.Param) []model.Param {
	return params
}

func Out(params ...model.Param) []model.Param {
	return params
}

func DelegateCommand(name string, delegate any, in []model.Param, out []model.Param) *model.Command {
	cmd := &model.Command{}
	cmd.Name = name

	handlerValue := reflect.ValueOf(delegate)

	if handlerValue.Kind() != reflect.Func {
		panic("delegate is not a function")
	}

	HandlerRef(cmd, &model.HandlerRef{
		Delegate: delegate,
		In:       in,
		Out:      out,
	})

	return cmd
}

func HandlerRef(cmd *model.Command, ref *model.HandlerRef) {
	cmd.Handler = func(c *cobra.Command, args []string) error {
		maybePS := slice.Fold(ref.In, &result.Operation[[]reflect.Value]{}, func(p model.Param, agg *result.Operation[[]reflect.Value]) *result.Operation[[]reflect.Value] {
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

		handlerValue := reflect.ValueOf(ref.Delegate)
		rs := handlerValue.Call(ps)

		for i, p := range ref.Out {
			if p.Kind() == model.BODY {
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
			} else if p.Kind() == model.ERROR {
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

	slice.ForEach(slice.Filter(ref.In, func(p model.Param) bool {
		return p.Kind() == model.FLAG
	}), func(p model.Param) {
		// only look at in params
		flag, ok := p.(*flag)

		if ok {
			cmd.AddFlag(&flag.Flag)
		}
	})
}

func (f *flag) Fetch(cmd *cobra.Command) (any, error) {
	visitor := registry.VisitorByKind(f.Flag.Kind)

	if visitor != nil {
		return visitor.Runtime(f.Flag.Name, cmd)
	}

	return nil, fmt.Errorf("unknown FlagKind: %s", f.Flag.Kind)
}

func (f *flag) Kind() model.ParamKind {
	return model.FLAG
}

func (b *body) Fetch(cmd *cobra.Command) (any, error) {
	bs, err := io.ReadAll(os.Stdin)

	if err != nil {
		return nil, err
	}

	return b.strategy.Read(bs)
}

func (b *body) Kind() model.ParamKind {
	return model.BODY
}

func (e *errror) Fetch(cmd *cobra.Command) (any, error) {
	return nil, fmt.Errorf("fetch is not implemented for error")
}

func (e *errror) Kind() model.ParamKind {
	return model.ERROR
}

func Json[T any]() model.Strategy {
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

func String() model.Strategy {
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

func (e *env) Kind() model.ParamKind {
	return model.ENV
}

func (e *env) Fetch(cmd *cobra.Command) (any, error) {
	val := utilz.Env(e.name, e.value)

	return val, nil
}
