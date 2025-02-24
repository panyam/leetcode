package conc

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func plainFizzBuzz(N int) (out []any) {
	for i := 1; i <= N; i++ {
		if i%15 == 0 {
			out = append(out, "fizzbuzz")
		} else if i%3 == 0 {
			out = append(out, "fizz")
		} else if i%5 == 0 {
			out = append(out, "buzz")
		} else {
			out = append(out, i)
		}
	}
	return
}

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
	assert.Equal(t, plainFizzBuzz(5), (ParallelFizzBuzz(5)))
	assert.Equal(t, plainFizzBuzz(20), (ParallelFizzBuzz(20)))
	assert.Equal(t, plainFizzBuzz(100), (ParallelFizzBuzz(100)))

	assert.Equal(t, plainFizzBuzz(5), (OrchestratingFizzBuzz(5)))
	assert.Equal(t, plainFizzBuzz(20), (OrchestratingFizzBuzz(20)))
	assert.Equal(t, plainFizzBuzz(100), (OrchestratingFizzBuzz(100)))
}

// In the parallel version each goroutine wakes up, checks if current number matches its
// "ability" and acts otherwise continues.  Each goroutine is unaware of the capability
// of other goroutines
//
// General pattern here is - you have N threads - than can only do N DIFFERENT kinds of
// work so a job queue strategy wont work.  Instead each thread needs to check repeatedly
// if the current work is of its type - the matching thread would do the work, then
// "modify" the current work to the next state (so the next thread that can match it will
// pick it up when it wakes up
func ParallelFizzBuzz(N int) (out []any) {
	var wg sync.WaitGroup
	wg.Add(4)

	i := 1
	var counterLock sync.Mutex

	// only calls Fizz

	doIf := func(name string, matcher func(curr int) (bool, any)) bool {
		// log.Println("Started: ", name)
		defer func() {
			// log.Printf("Quitting: Runner: %s, i: %d", name, i)
			wg.Done()
		}()
		for {
			var output any
			matched := false
			counterLock.Lock()
			done := i > N
			if !done {
				matched, output = matcher(i)
				if matched {
					// log.Printf("%s matched: %d, OK: %t", name, i, matched)
					out = append(out, output)
					i += 1
				}
			}
			counterLock.Unlock()
			if done {
				return false
			}
		}
	}

	go doIf("number", func(curr int) (bool, any) {
		return curr%3 != 0 && curr%5 != 0, curr
	})

	go doIf("buzzer", func(curr int) (bool, any) {
		return curr%3 != 0 && curr%5 == 0, "buzz"
	})

	go doIf("fizzer", func(curr int) (bool, any) {
		return curr%3 == 0 && curr%5 != 0, "fizz"
	})

	go doIf("fizzbuzzer", func(curr int) (bool, any) {
		return curr%3 == 0 && curr%5 == 0, "fizzbuzz"
	})

	wg.Wait()
	return
}

// In this version, we can use an orchestrator goroutine - for the "number"s
// which sends the signal to the "next" goroutine based on the number
// and each goroutine sends a response back (via channels)
func OrchestratingFizzBuzz(N int) (out []any) {
	mainChan := make(chan int)
	fbChan := make(chan int)
	fChan := make(chan int)
	bChan := make(chan int)
	donechan := make(chan bool, 3)

	var wg sync.WaitGroup
	wg.Add(4)

	doIt := func(name string, inchan chan int, donechan chan bool) {
		// log.Println("Starting: ", name)
		// defer log.Println("Quitting: ", name)
		defer wg.Done()
		for {
			// log.Printf("%s reading : %t", name, true)
			select {
			case <-donechan:
				return
			case i := <-inchan:
				// log.Printf("%s received: %d", name, i)
				out = append(out, name)
				mainChan <- i + 1
			}
		}
	}

	main := func() {
		defer func() {
			// log.Println("Quiting main...")
			donechan <- true
			donechan <- true
			donechan <- true
			wg.Done()
		}()
		i := 1
		for i <= N {
			if i%3 == 0 && i%5 == 0 {
				fbChan <- i
				i = <-mainChan
			} else if i%3 == 0 {
				fChan <- i
				i = <-mainChan
			} else if i%5 == 0 {
				bChan <- i
				i = <-mainChan
			} else {
				out = append(out, i)
				i += 1
			}
		}
	}

	go main()
	go doIt("fizz", fChan, donechan)
	go doIt("buzz", bChan, donechan)
	go doIt("fizzbuzz", fbChan, donechan)
	wg.Wait()
	return
}
