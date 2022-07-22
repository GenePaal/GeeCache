package geecache

import (
	"prac/geecache/lru"
	"sync"
)

type cache struct {
	mu         sync.Mutex
	lru        *lru.Cache
	cacheBytes int64
}

func (c *cache) push(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes, nil)
	}
	c.lru.Push(key, value)
}

func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}

	if v, ok := c.lru.Get(key); ok {
		// v 是 lru包下的 Value 类型, 由于ByteView实现了
		// Value 接口，因此可以转换为ByteView
		// 将接口类型转换为具体类型可以使用
		// Interface.(Struct)
		// v中存的是真正的缓存值，不是地址
		// 严格意义上来说 v中存的是切片， 也不能说是真正的值(笑哭)
		// 真正的值是v中的切片指向的底层byte数组
		return v.(ByteView), ok
	}

	return
}
