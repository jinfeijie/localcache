package localcache

import (
	"sync"
)

// go原生支持的sync.map并发读写安全的内存缓存
type Cache struct {
	m sync.Map
}

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) Set(key string, val interface{}) {
	c.m.Store(key, val)
}

func (c *Cache) Get(key string) interface{} {
	if val, ok := c.m.Load(key); ok {
		return val
	}
	return nil
}

// 如果不存在，就设置，并且返回设置的值
func (c *Cache) GetOrSet(key string, val interface{}) interface{} {
	if ac, ok := c.m.LoadOrStore(key, val); ok {
		return ac
	}
	return val
}

func (c *Cache) Delete(key string) {
	c.m.Delete(key)
}

func (c *Cache) IsExist(key string) bool {
	_, ok := c.m.Load(key)
	return ok
}

func (c *Cache) Size() (len int) {
	c.m.Range(func(key, value interface{}) bool {
		len++
		return true
	})
	return len
}
