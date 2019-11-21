// Package registry uses the go-micro registry for selection
package registry

import (
	"micro_learn/micro/go-micro/client/selector"
)

// NewSelector returns a new registry selector
func NewSelector(opts ...selector.Option) selector.Selector {
	return selector.NewSelector(opts...)
}
