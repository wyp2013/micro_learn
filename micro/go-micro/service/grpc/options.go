package grpc

import (
	"crypto/tls"

	"micro_learn/micro/go-micro"
	gc "micro_learn/micro/go-micro/client/grpc"
	gs "micro_learn/micro/go-micro/server/grpc"
)

// WithTLS sets the TLS config for the service
func WithTLS(t *tls.Config) micro.Option {
	return func(o *micro.Options) {
		o.Client.Init(
			gc.AuthTLS(t),
		)
		o.Server.Init(
			gs.AuthTLS(t),
		)
	}
}
