// Code generated by Kitex v0.3.2. DO NOT EDIT.

package serviceproxy

import (
	"github.com/cloudwego/kitex-examples/grpcproxy/kitex_gen/grpcproxy"
	"github.com/cloudwego/kitex/server"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler grpcproxy.ServiceProxy, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}