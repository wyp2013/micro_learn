package runtimevar

import (
	"context"

	"micro_learn/micro/go-micro/config/source"
	"gocloud.dev/runtimevar"
)

type variableKey struct{}

// WithVariable sets the runtimevar.Variable.
func WithVariable(v *runtimevar.Variable) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, variableKey{}, v)
	}
}
