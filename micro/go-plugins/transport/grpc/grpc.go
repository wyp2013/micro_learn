// Package grpc provides a grpc transport
// Deprecated: use `micro_learn/micro/go-micro/transport/grpc` instead
package grpc

import (
	"micro_learn/micro/go-micro/transport"
	"micro_learn/micro/go-micro/transport/grpc"
)

// Deprecated: use `micro_learn/micro/go-micro/transport/grpc` instead
func NewTransport(opts ...transport.Option) transport.Transport {
	return grpc.NewTransport(opts...)
}
