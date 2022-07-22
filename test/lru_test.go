package test

import (
	"prac/geecache/lru"
	"testing"
)

type String string

func (d String) Len() int {
	return len(d)
}

func TestLRUGet(t *testing.T) {

	instance := lru.New(int64(0), nil)
	instance.Push("key1", String("1234"))
	if v, ok := instance.Get("key1"); !ok || string(v.(String)) != "1234" {
		t.Fatalf("cache hit key1=1234 failed")
	}
	if _, ok := instance.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed")
	} else {
		t.Fatalf("没有找到key2")
	}
}
