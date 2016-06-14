package workerpool

import (
	"log"
	"testing"
	"time"
)

type task0 struct {
	name string
}

func (t task0) Run() {
	log.Println(t.name)
	time.Sleep(1 * time.Second)
}

type task1 struct {
	name string
}

func (t task1) Run() {
	log.Println(t.name)
	time.Sleep(2 * time.Second)
}

func TestWorkerPool128(b *testing.T) {
	p := NewPool(128)
	go p.Start()
	x := &task0{name: "task0"}
	y := &task1{name: "task1"}
	for i := 0; i < 256; i++ {
		p.Run(x)
		p.Run(y)
	}
}
