package conc

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

func init() {
	rand.Seed(uint64(time.Now().Unix()))
}

type counter struct {
	c int
	sync.Mutex
}

func (c *counter) Incr() {
	c.Lock()
	c.c += 1
	c.Unlock()
}
func (c *counter) Get() (res int) {
	c.Lock()
	res = c.c
	c.Unlock()
	return
}
func worker(c *counter, before *Barrier, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		before.Before()
		c.Incr()
		before.After()
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Println(c.Get())
	}
	wg.Done()
}

func ExampleBarrier() {
	var wg sync.WaitGroup
	workers := 5
	before := NewBarrier(int(workers))
	c := counter{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go worker(&c, before, &wg)
	}
	wg.Wait()
	// Output:
	// 5
	// 5
	// 5
	// 5
	// 5
	// 10
	// 10
	// 10
	// 10
	// 10
	// 15
	// 15
	// 15
	// 15
	// 15
}
