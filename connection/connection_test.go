package connection

import "testing"

func benchmarkFunc(s string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		go func() {
			Conn(s)
		}()
	}
}

func BenchmarkFib1(b *testing.B) { benchmarkFunc("CgObKRQxNTQ3MzQzMTc4MDU1NzM4MjU5MAAA", b) }
