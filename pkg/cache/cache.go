package cache

import (
	"sync"
	"time"
)

type entry[V any] struct {
	expiry int64
	value  V
}

type TTLCache[V any] struct {
	mutex sync.RWMutex
	cache map[string]entry[V]
}

func (c *TTLCache[V]) Put(key string, value V, expiration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	now := time.Now()
	c.cache[key] = entry[V]{
		expiry: now.Add(expiration).Unix(),
		value:  value,
	}
}

func (c *TTLCache[V]) Get(key string) (value V, ok bool) {
	entry, ok := c.cache[key]
	if !ok {
		return
	}
	if now := time.Now().Unix(); now <= entry.expiry {
		value = entry.value
	} else {
		ok = false
		c.Remove(key)
	}
	return
}

func (c *TTLCache[V]) Remove(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.cache, key)
}

func New[V any]() *TTLCache[V] {
	return &TTLCache[V]{
		cache: make(map[string]entry[V]),
	}
}
