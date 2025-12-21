package builder

import (
	"github.com/Meduzz/commando/delegate"
	"github.com/Meduzz/commando/model"
)

type (
	paramBuilder struct {
		builder *handlerRefBuilder
		in      bool
	}
)

var _ ParamBuilder = (*paramBuilder)(nil)

func NewParamBuilder(builder *handlerRefBuilder, in bool) ParamBuilder {
	return &paramBuilder{
		builder: builder,
		in:      in,
	}
}

func (p *paramBuilder) Body(strategy model.Strategy) ParamBuilder {
	if p.in {
		p.builder.in = append(p.builder.in, delegate.Body(strategy))
	} else {
		p.builder.out = append(p.builder.out, delegate.Body(strategy))
	}

	return p
}

func (p *paramBuilder) Env(name, value string) ParamBuilder {
	if p.in {
		p.builder.in = append(p.builder.in, delegate.Env(name, value))
	}

	return p
}

func (p *paramBuilder) Flag(name string, kind model.FlagKind, defaultValue any, description string) ParamBuilder {
	if p.in {
		p.builder.in = append(p.builder.in, delegate.Flag(name, kind, defaultValue, description))
	}

	return p
}

func (p *paramBuilder) Error() ParamBuilder {
	if !p.in {
		p.builder.out = append(p.builder.out, delegate.Error())
	}

	return p
}
