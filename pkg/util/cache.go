package util

import (
	"sync"
	"time"
)

type entry[V any] struct {
	expiry int64
	value  V
}

type TTLCache[V any] struct {
	ttl   time.Duration
	mutex sync.RWMutex
	cache map[string]entry[V]
}

func (c *TTLCache[V]) Put(key string, value V) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	now := time.Now()
	c.cache[key] = entry[V]{
		expiry: now.Add(c.ttl).Unix(),
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

func NewCache[V any](ttl time.Duration) *TTLCache[V] {
	return &TTLCache[V]{
		ttl:   ttl,
		cache: make(map[string]entry[V]),
	}
}
