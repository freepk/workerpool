package workerpool

import (
	"testing"
)

func TestWorkerPool128(b *testing.T) {
	p := NewPool(128)
	go p.Start()
	x := func() {}
	y := func() {}
	for i := 0; i < 1000000; i++ {
		p.Run(x)
		p.Run(y)
	}
}

func BenchmarkPool(b *testing.B) {
	p := NewPool(128)
	go p.Start()
	t := func() {}
	for i := 0; i < b.N; i++ {
		p.Run(t)
	}
}
