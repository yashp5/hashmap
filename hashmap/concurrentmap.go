package hashmap

import (
	"sync"
)

const (
	defaultShardCount = 16
)

type shard struct {
	mu sync.RWMutex
	m  *ChainingMap
}

type ConcurrentMap struct {
	shards     []*shard
	shardCount int
}

func NewConcurrentMap() *ConcurrentMap {
	shards := make([]*shard, defaultShardCount)
	for i := range shards {
		shards[i] = &shard{
			m: NewChainingMap(),
		}
	}
	return &ConcurrentMap{
		shards:     shards,
		shardCount: defaultShardCount,
	}
}

func (c *ConcurrentMap) getShard(key string) *shard {
	idx := Hash(key) % c.shardCount
	return c.shards[idx]
}

func (c *ConcurrentMap) Get(key string) (any, bool) {
	s := c.getShard(key)
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.m.Get(key)
}

func (c *ConcurrentMap) Put(key string, value any) {
	s := c.getShard(key)
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m.Put(key, value)
}

func (c *ConcurrentMap) Delete(key string) bool {
	s := c.getShard(key)
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.m.Delete(key)
}
