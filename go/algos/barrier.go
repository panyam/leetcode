package algos

import (
	"sync"
)

type Barrier struct {
	m      sync.Mutex
	c      uint
	n      uint
	before chan bool
}

func NewBarrier(n uint) *Barrier {
	out := &Barrier{n: n,
		before: make(chan bool, n),
	}
	return out
}

func (b *Barrier) Wait() {
	if b.n <= 1 {
		return
	}
	b.m.Lock()
	b.c += 1
	if b.c == b.n {
		// All have stopped so reset c
		b.c = 0
		for range b.n {
			b.before <- true
		}
	}
	b.m.Unlock()
	// If N - 1 goroutines here - they will wait
	// remember recv will always block - only put will block if > capacity
	<-b.before
}
