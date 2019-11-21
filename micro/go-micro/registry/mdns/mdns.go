// Package mdns provides a multicast dns registry
package mdns

import (
	"micro_learn/micro/go-micro/registry"
)

// NewRegistry returns a new mdns registry
func NewRegistry(opts ...registry.Option) registry.Registry {
	return registry.NewRegistry(opts...)
}
