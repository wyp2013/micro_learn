package template

var (
	WrapperAPI = `package client

import (
	"context"

	"micro_learn/micro/go-micro"
	"micro_learn/micro/go-micro/server"
	{{.Alias}} "path/to/service/proto/{{.Alias}}"
)

type {{.Alias}}Key struct {}

// FromContext retrieves the client from the Context
func {{title .Alias}}FromContext(ctx context.Context) ({{.Alias}}.{{title .Alias}}Service, bool) {
	c, ok := ctx.Value({{.Alias}}Key{}).({{.Alias}}.{{title .Alias}}Service)
	return c, ok
}

// Client returns a wrapper for the {{title .Alias}}Client
func {{title .Alias}}Wrapper(service micro.Service) server.HandlerWrapper {
	client := {{.Alias}}.New{{title .Alias}}Service("go.micro.srv.template", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, {{.Alias}}Key{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
`
)
