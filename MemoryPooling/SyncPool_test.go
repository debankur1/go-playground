package main_test

import (
	"runtime"
	"sync"
	"testing"
)
type Small struct {
a int
}
var pool = sync.Pool{
New: func() interface{} {
	return new(Small)
    },
}
func inc(s *Small) {
	s.a++
}
func BenchmarkWithoutPool(b *testing.B) {
	var s *Small
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			s = &Small{ a: 1, }
			b.StopTimer(); inc(s); b.StartTimer()
		}
	}
}
func BenchmarkWithPool(b *testing.B) {
	var s *Small
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			s = pool.Get().(*Small)
			s.a = 1
			b.StopTimer(); inc(s); b.StartTimer()
			pool.Put(s)
		}
	}
	runtime.GC()
}