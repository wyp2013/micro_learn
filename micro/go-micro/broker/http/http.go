// Package http provides a http based message broker
package http

import (
	"micro_learn/micro/go-micro/broker"
)

// NewBroker returns a new http broker
func NewBroker(opts ...broker.Option) broker.Broker {
	return broker.NewBroker(opts...)
}
