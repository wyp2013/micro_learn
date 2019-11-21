// Package kubernetes lets you initialise a web service using the k8s registry plugin
package kubernetes

import (
	"micro_learn/micro/go-micro"
	"micro_learn/micro/go-micro/service/grpc"
	"micro_learn/micro/go-micro/web"
	"micro_learn/micro/go-plugins/registry/kubernetes"

	// static selector offloads load balancing to k8s services
	// enable with MICRO_SELECTOR=static or --selector=static
	// requires user to create k8s services
	"micro_learn/micro/go-plugins/client/selector/static"
)

// NewService returns a web service for kubernetes
func NewService(opts ...web.Option) web.Service {
	// setup
	k := kubernetes.NewRegistry()
	st := static.NewSelector()

	// create new service
	service := grpc.NewService(
		micro.Registry(k),
		micro.Selector(st),
	)

	// prepend option
	options := []web.Option{
		web.MicroService(service),
	}

	options = append(options, opts...)

	// return new service
	return web.NewService(options...)
}
