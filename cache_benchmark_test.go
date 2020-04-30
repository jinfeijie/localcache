package localcache

import (
	"runtime"
	"testing"
)

func init() {
	runtime.GOMAXPROCS(8)
}

func BenchmarkNewCache(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			NewCache()
		}
	})
	b.ReportAllocs()
}

func BenchmarkCache_Set(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		set := NewCache()
		for pb.Next() {
			set.Set("111111", "1")
		}
	})
	b.ReportAllocs()
}

func BenchmarkCache_Get(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		set := NewCache()
		set.Set("get", "1")
		for pb.Next() {
			set.Get("get")
		}
	})
	b.ReportAllocs()
}

func BenchmarkCache_GetOrSet(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		set := NewCache()
		for pb.Next() {
			set.GetOrSet("GetOrSet", "1")
		}
	})
	b.ReportAllocs()
}

func BenchmarkCache_IsExist(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		set := NewCache()
		set.Set("IsExist", "1")
		for pb.Next() {
			set.IsExist("IsExist")
		}
	})
	b.ReportAllocs()
}

func BenchmarkCache_Size(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		set := NewCache()
		for i := 0; i < 50; i++ {
			set.Set("Size"+string(i), i)
		}
		for pb.Next() {
			set.Size()
		}
	})

	b.RunParallel(func(pb *testing.PB) {
		set := NewCache()
		for i := 0; i < 10; i++ {
			set.Set("Size"+string(i), i)
		}
		for pb.Next() {
			set.Size()
		}
	})
	b.ReportAllocs()
}

func BenchmarkCache_All(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		all := NewCache()
		for i := 0; i < 50; i++ {
			all.Set("all"+string(i), i)
		}
		for pb.Next() {
			all.All()
		}
	})

	b.RunParallel(func(pb *testing.PB) {
		all := NewCache()
		for i := 0; i < 10; i++ {
			all.Set("all"+string(i), i)
		}
		for pb.Next() {
			all.All()
		}
	})
	b.ReportAllocs()
}
