package memory

import (
	"micro_learn/micro/go-micro/config/loader"
	"micro_learn/micro/go-micro/config/reader"
	"micro_learn/micro/go-micro/config/source"
)

// WithSource appends a source to list of sources
func WithSource(s source.Source) loader.Option {
	return func(o *loader.Options) {
		o.Source = append(o.Source, s)
	}
}

// WithReader sets the config reader
func WithReader(r reader.Reader) loader.Option {
	return func(o *loader.Options) {
		o.Reader = r
	}
}
