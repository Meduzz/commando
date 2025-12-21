package builder

import (
	"reflect"

	"github.com/Meduzz/commando/model"
)

type (
	handlerRefBuilder struct {
		delegate any
		in       []model.Param
		out      []model.Param
	}
)

var _ HandlerRefBuilder = (*handlerRefBuilder)(nil)

func NewHandlerRefBuilder(delegate any) HandlerRefBuilder {
	handlerValue := reflect.ValueOf(delegate)

	if handlerValue.Kind() != reflect.Func {
		panic("delegate is not a function")
	}

	return &handlerRefBuilder{
		delegate: delegate,
	}
}

func (h *handlerRefBuilder) In(builder ParamBuilderFunc) HandlerRefBuilder {
	b := NewParamBuilder(h, true)
	builder(b)

	return h
}

func (h *handlerRefBuilder) Out(builder ParamBuilderFunc) HandlerRefBuilder {
	b := NewParamBuilder(h, false)
	builder(b)

	return h
}

func (h *handlerRefBuilder) build() *model.HandlerRef {
	return &model.HandlerRef{
		Delegate: h.delegate,
		In:       h.in,
		Out:      h.out,
	}
}

func HandlerRef(handler any, builder HandlerRefBuilderFunc) *model.HandlerRef {
	b := NewHandlerRefBuilder(handler)
	builder(b)

	return b.build()
}
