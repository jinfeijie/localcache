package localcache

import (
	"testing"
	"time"
)

func TestCache_Size(t *testing.T) {
	a := NewCache()
	b := NewCache()

	a.GetOrSet("哈哈哈哈哈", time.Now().UnixNano())
	b.GetOrSet("哈哈哈哈哈", time.Now().UnixNano())

	t.Log(a.Get("哈哈哈哈哈"))
	t.Log(a.Size())
	t.Log(b.Get("哈哈哈哈哈"))
	t.Log(b.Size())
}
