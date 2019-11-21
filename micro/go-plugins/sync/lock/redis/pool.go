package redis

import (
	"sync"

	"github.com/gomodule/redigo/redis"
)

// 可以实现redis的连接池
type pool struct {
	sync.Mutex
	i     int
	addrs []string
}

func (p *pool) Get() redis.Conn {
	for i := 0; i < 3; i++ {
		p.Lock()
		addr := p.addrs[p.i%len(p.addrs)]
		p.i++
		p.Unlock()

		c, err := redis.Dial("tcp", addr)
		if err != nil {
			continue
		}
		return c
	}
	return nil
}
