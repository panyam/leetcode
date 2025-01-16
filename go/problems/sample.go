package main

import (
	"fmt"
	"sync"
)

const numWorkers = 5

type Counter struct {
	mutex sync.RWMutex
	value int
}

// Runs the boot strap of the ith stage for a given worker
func bootstrapper(loopIndex int) {
}

func job(loopIndex int) {
}

func worker(i int) {
	bootstrap()
	job()
}

func main() {
	for i := range numWorkers {
		go startWorker()
	}
	fmt.Println("vim-go")
}
