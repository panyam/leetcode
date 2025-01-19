package conc

import (
	"log"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

type Barrier struct {
	m        sync.Mutex
	nWritten int
	c        uint
	n        uint
	gen      uint
	before   chan bool
	seats    []bool
}

func NewBarrier(n uint) *Barrier {
	if n <= 1 {
		panic("n must be > 1")
	}
	out := &Barrier{n: n,
		before: make(chan bool, n),
	}
	for range n {
		out.seats = append(out.seats, false)
	}
	return out
}

func (b *Barrier) Wait(tid int) {
	b.seats[tid] = true
	b.m.Lock()
	oldGen := b.gen
	log.Printf("%d - %03d - Starting to Wait, C: %d, nWritten: %d, Seats: %v", tid, oldGen, b.c, b.nWritten, b.seats)
	b.c += 1
	if b.c == b.n {
		// All have stopped so reset c
		b.c = 0
		for range b.n {
			b.before <- true
			b.nWritten += 1
		}
		log.Printf("%d - %03d - Done adding signals, nWritten: %d...", tid, oldGen, b.nWritten)
		log.Printf("%d - %03d - ===============", tid, oldGen)
		b.gen += 1
	}
	b.m.Unlock()
	// If N - 1 goroutines here - they will wait
	// remember recv will always block - only put will block if > capacity
	//
	// What if:
	// N - 1 go routines are waiting here and they havent been resumed yet and since before is a buffered channel
	// the Nth goroutine that put 5 values in this chan, could be here, read its value.  Now if Wait (or Wait2) was called
	// then N would hit the c == n condition and put in 5 more values in the chan - blocking on the 6th value (as only 4 goroutines are waiting to read)
	// So this *does* need a before and after so that only 1 is called
	//
	// Way to test this is to have a small "delay" after the wait to ensure the Nth go routine has the least delay
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	b.m.Lock()
	log.Printf("%d - %03d - About to Read from chan, Seats: %v", tid, oldGen, b.seats)
	b.m.Unlock()

	<-b.before

	b.seats[tid] = false
	b.m.Lock()
	b.nWritten -= 1
	b.m.Unlock()
	log.Printf("%d - %03d - Done Waiting, Seats: %v ###############", tid, oldGen, b.seats)
}
