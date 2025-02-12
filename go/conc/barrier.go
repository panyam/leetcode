package conc

import (
	"sync"
)

type Barrier struct {
	m        sync.Mutex
	nWritten int
	c        uint
	n        uint
	before   chan bool
	after    chan bool
}

func NewBarrier(n uint) *Barrier {
	if n <= 1 {
		panic("n must be > 1")
	}
	out := &Barrier{n: n,
		before: make(chan bool, n),
		after:  make(chan bool, n),
	}
	return out
}

func (b *Barrier) Before() {
	b.m.Lock()
	b.c += 1
	if b.c == b.n {
		// open 2nd gate
		for i := uint(0); i < b.n; i++ {
			b.before <- true
		}
	}
	b.m.Unlock()
	<-b.before
}
func (b *Barrier) After() {
	b.m.Lock()
	b.c -= 1
	if b.c == 0 {
		// open 1st gate
		for i := uint(0); i < b.n; i++ {
			b.after <- true
		}
	}
	b.m.Unlock()
	<-b.after
}
