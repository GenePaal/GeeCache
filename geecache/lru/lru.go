package lru

import (
	"container/list"
)

type Cache struct {
	// 最大的内存
	maxBytes int64
	// 已使用的内存
	nbytes int64
	// 双向链表的首地址
	ll *list.List
	// 键是字符串，值是list中元素的指针
	cache map[string]*list.Element
	// 某条记录被移除时的回调函数
	OnEvicted func(key string, value Value)
}

// 接口只能指向实现了自己的方法的实例
type Value interface {
	Len() int
}

type entry struct {
	key   string
	value Value
}

// 实例化Cache, 可以理解为Cache的构造函数
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// 查找
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		// any类型 转 entry指针类型
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// 删除
func (c *Cache) RemoveOldest() {
	// element.value

	// 取队首
	ele := c.ll.Back()
	if ele != nil {
		// 删除队首
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}

}

// 新增/修改
// 之前未存在, 则新增
// 之前存在过, 则更新

func (c *Cache) Push(key string, value Value) {
	if ele, ok := c.cache[key]; ok { // 更新
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else { // 新增
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nbytes += int64(value.Len()) + int64(len(key))
		if c.maxBytes <= 0 {
			return
		}
		for c.nbytes > c.maxBytes {
			c.RemoveOldest()
		}
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
