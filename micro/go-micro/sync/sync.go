// Package sync is a distributed synchronization framework
package sync

import (
	"micro_learn/micro/go-micro/data/store"
	"micro_learn/micro/go-micro/sync/leader"
	"micro_learn/micro/go-micro/sync/lock"
	"micro_learn/micro/go-micro/sync/task"
	"micro_learn/micro/go-micro/sync/time"
)

// Map provides synchronized access to key-value storage.
// It uses the store interface and lock interface to
// provide a consistent storage mechanism.
type Map interface {
	// Read value with given key
	Read(key, val interface{}) error
	// Write value with given key
	Write(key, val interface{}) error
	// Delete value with given key
	Delete(key interface{}) error
	// Iterate over all key/vals. Value changes are saved
	Iterate(func(key, val interface{}) error) error
}

// Cron is a distributed scheduler using leader election
// and distributed task runners. It uses the leader and
// task interfaces.
type Cron interface {
	Schedule(task.Schedule, task.Command) error
}

type Options struct {
	Leader leader.Leader
	Lock   lock.Lock
	Store  store.Store
	Task   task.Task
	Time   time.Time
}

type Option func(o *Options)
