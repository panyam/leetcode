package conc

import (
	"fmt"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/exp/rand"
)

func producer(id int, nvals int, wg *sync.WaitGroup, outch chan uint64, counter *atomic.Uint64) {
	for range nvals {
		counter.Add(1)
		outch <- counter.Load()
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
	wg.Done()
}

func consumer(id int, nvals int, wg *sync.WaitGroup, inch chan uint64) {
	var vals []string
	for range nvals {
		val := <-inch
		vals = append(vals, fmt.Sprintf("Consumer: %03d, Value: %03d", id, val))
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
	sort.Strings(vals)
	for _, s := range vals {
		fmt.Println(s)
	}
	wg.Done()
}

func ExampleOneProducerOneConsumer() {
	var wg sync.WaitGroup
	wg.Add(2)
	var counter atomic.Uint64
	ch := make(chan uint64)
	nvals := 5

	go producer(1, nvals, &wg, ch, &counter)
	go consumer(1, nvals, &wg, ch)
	wg.Wait()
	// Output:
	// Consumer: 001, Value: 001
	// Consumer: 001, Value: 002
	// Consumer: 001, Value: 003
	// Consumer: 001, Value: 004
	// Consumer: 001, Value: 005
}

func ExampleManyProducerOneConsumer() {
	var wg sync.WaitGroup
	var counter atomic.Uint64
	ch := make(chan uint64)
	nvals := 5
	nprods := 5

	wg.Add(1)
	go consumer(1, nvals*nprods, &wg, ch)
	for i := range nprods {
		wg.Add(1)
		go producer(i, nvals, &wg, ch, &counter)
	}
	wg.Wait()
	// Output:
	// Consumer: 001, Value: 001
	// Consumer: 001, Value: 002
	// Consumer: 001, Value: 003
	// Consumer: 001, Value: 004
	// Consumer: 001, Value: 005
	// Consumer: 001, Value: 006
	// Consumer: 001, Value: 007
	// Consumer: 001, Value: 008
	// Consumer: 001, Value: 009
	// Consumer: 001, Value: 010
	// Consumer: 001, Value: 011
	// Consumer: 001, Value: 012
	// Consumer: 001, Value: 013
	// Consumer: 001, Value: 014
	// Consumer: 001, Value: 015
	// Consumer: 001, Value: 016
	// Consumer: 001, Value: 017
	// Consumer: 001, Value: 018
	// Consumer: 001, Value: 019
	// Consumer: 001, Value: 020
	// Consumer: 001, Value: 021
	// Consumer: 001, Value: 022
	// Consumer: 001, Value: 023
	// Consumer: 001, Value: 024
	// Consumer: 001, Value: 025
}

func ExampleOneProducerManyConsumer() {
	var wg sync.WaitGroup
	var counter atomic.Uint64
	ch := make(chan uint64)
	nvals := 5
	ncons := 5

	for i := range ncons {
		wg.Add(1)
		go consumer(i, nvals, &wg, ch)
	}
	wg.Add(1)
	go producer(1, nvals*ncons, &wg, ch, &counter)
	wg.Wait()
	// Output:
	// Consumer: 000, Value: 001
	// Consumer: 000, Value: 006
	// Consumer: 000, Value: 013
	// Consumer: 000, Value: 016
	// Consumer: 000, Value: 021
	// Consumer: 001, Value: 002
	// Consumer: 001, Value: 007
	// Consumer: 001, Value: 012
	// Consumer: 001, Value: 017
	// Consumer: 001, Value: 022
	// Consumer: 003, Value: 003
	// Consumer: 003, Value: 008
	// Consumer: 003, Value: 011
	// Consumer: 003, Value: 018
	// Consumer: 003, Value: 023
	// Consumer: 004, Value: 005
	// Consumer: 004, Value: 010
	// Consumer: 004, Value: 015
	// Consumer: 004, Value: 019
	// Consumer: 004, Value: 024
	// Consumer: 002, Value: 004
	// Consumer: 002, Value: 009
	// Consumer: 002, Value: 014
	// Consumer: 002, Value: 020
	// Consumer: 002, Value: 025
}

func producertobuffer(id int, nvals int, wg *sync.WaitGroup, outch chan uint64, counter *atomic.Uint64) {
	for range nvals {
		counter.Add(1)
		outch <- counter.Load()
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
	wg.Done()
}

func consumerfrombuffer(id int, nvals int, wg *sync.WaitGroup, inch chan uint64) {
	var vals []string
	for range nvals {
		val := <-inch
		vals = append(vals, fmt.Sprintf("Consumer: %03d, Value: %03d", id, val))
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
	sort.Strings(vals)
	for _, s := range vals {
		fmt.Println(s)
	}
	wg.Done()
}

// In the previous examples we implicitly used a fixed buffer of size "1".
// We could have changed the channel to be an N elem bounded channel.
// Now we will look at an infinite buffer.  Producers keep dumping into a queue
// and consumers will just read of this queue.
func ExampleInfiniteBuffer() {
	var wg sync.WaitGroup
	var counter atomic.Uint64
	ch := make(chan uint64)
	nvals := 5
	ncons := 5

	for i := range ncons {
		wg.Add(1)
		go consumer(i, nvals, &wg, ch)
	}
	wg.Add(1)
	go producer(1, nvals*ncons, &wg, ch, &counter)
	wg.Wait()
	// Output:
	// Consumer: 000, Value: 001
	// Consumer: 000, Value: 006
	// Consumer: 000, Value: 013
	// Consumer: 000, Value: 016
	// Consumer: 000, Value: 021
	// Consumer: 001, Value: 002
	// Consumer: 001, Value: 007
	// Consumer: 001, Value: 012
	// Consumer: 001, Value: 017
	// Consumer: 001, Value: 022
	// Consumer: 003, Value: 003
	// Consumer: 003, Value: 008
	// Consumer: 003, Value: 011
	// Consumer: 003, Value: 018
	// Consumer: 003, Value: 023
	// Consumer: 004, Value: 005
	// Consumer: 004, Value: 010
	// Consumer: 004, Value: 015
	// Consumer: 004, Value: 019
	// Consumer: 004, Value: 024
	// Consumer: 002, Value: 004
	// Consumer: 002, Value: 009
	// Consumer: 002, Value: 014
	// Consumer: 002, Value: 020
	// Consumer: 002, Value: 025
}
