package cache

import "sync"

type Cache struct {
	sync.RWMutex
	data map[uint64]interface{}
}

func NewCache() *Cache {
	cache := &Cache{}
	cache.data = make(map[uint64]interface{})
	return cache
}

func (tool *Cache) SetCache(c map[uint64]interface{}) {
	tool.Lock()
	defer tool.Unlock()
	tool.data = c
}

func (tool *Cache) GetCache() map[uint64]interface{} {
	tool.RLock()
	defer tool.RUnlock()
	return tool.data
}

func (tool *Cache) AddCache(c map[uint64]interface{}) {
	tool.Lock()
	defer tool.Unlock()
	for k, v := range c {
		tool.data[k] = v
	}
}

func (tool *Cache) DeleteCache(c []uint64) {
	tool.Lock()
	defer tool.Unlock()
	for _, v := range c {
		delete(tool.data, v)
	}
}
