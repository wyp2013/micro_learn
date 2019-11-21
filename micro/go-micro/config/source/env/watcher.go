package env

import (
	"errors"

	"micro_learn/micro/go-micro/config/source"
)

type watcher struct {
	exit chan struct{}
}

func (w *watcher) Next() (*source.ChangeSet, error) {
	<-w.exit

	return nil, errors.New("watcher stopped")
}

func (w *watcher) Stop() error {
	close(w.exit)
	return nil
}

func newWatcher() (source.Watcher, error) {
	return &watcher{exit: make(chan struct{})}, nil
}
