package workerpool

import (
	"sync/atomic"
	"testing"
)

var (
	testCounter = int64(0)
)

type task0 struct {
	name string
}

func (t task0) Run() {
	atomic.AddInt64(&testCounter, 1)
}

type task1 struct {
	name string
}

func (t task1) Run() {
	atomic.AddInt64(&testCounter, 1)
}

type task2 int

func (t task2) Run() {
	atomic.AddInt64(&testCounter, 1)
}

func TestWorkerPool128(b *testing.T) {
	p := NewPool(128)
	go p.Start()
	x := &task0{name: "task0"}
	y := &task1{name: "task1"}
	for i := 0; i < 1000000; i++ {
		p.Run(x)
		p.Run(y)
	}
}

func BenchmarkPool(b *testing.B) {
	p := NewPool(128)
	go p.Start()
	var t task2
	for i := 0; i < b.N; i++ {
		p.Run(t)
	}
}
