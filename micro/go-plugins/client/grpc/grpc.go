// Package grpc provides a gRPC client
// Deprecated: use `micro_learn/micro/go-micro/client/grpc` instead
package grpc

import (
	"micro_learn/micro/go-micro/client"
	"micro_learn/micro/go-micro/client/grpc"
)

// Deprecated: use `micro_learn/micro/go-micro/client/grpc` instead
func NewClient(opts ...client.Option) client.Client {
	return grpc.NewClient(opts...)
}
