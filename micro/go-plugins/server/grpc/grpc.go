// Package grpc provides a grpc server
// Deprecated: use `micro_learn/micro/go-micro/server/grpc` instead
package grpc

import (
	"micro_learn/micro/go-micro/server"
	"micro_learn/micro/go-micro/server/grpc"
)

// We use this to wrap any debug handlers so we preserve the signature Debug.{Method}
// Deprecated: use `micro_learn/micro/go-micro/server/grpc` instead
type Debug = grpc.Debug

var (
	// DefaultMaxMsgSize define maximum message size that server can send
	// or receive.  Default value is 4MB.
	// Deprecated: use `micro_learn/micro/go-micro/server/grpc` instead
	DefaultMaxMsgSize = grpc.DefaultMaxMsgSize
)

// Deprecated: use `micro_learn/micro/go-micro/server/grpc` instead
func NewServer(opts ...server.Option) server.Server {
	return grpc.NewServer(opts...)
}
