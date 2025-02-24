package conc

import (
	"log"
	"sort"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Process an input array with N workers.
// You are given an array of inputs and N denoting how many workers should be created
// You are also given M Mapper functions.
// The Outputs[i] = Mappers[0](Mapper[1](Mapper[2]....(Mapper[M](Inputs[i]))))

// The caveat of this is that All workers must finish with a stage (a mapper) before
// proceeding to the next stage

type PAPDebug struct {
	stage int
	index int
	wid   int // <- not needed can be infered from index
}

func (d *PAPDebug) Less(another PAPDebug) bool {
	if d.stage == another.stage {
		return d.index < another.index
	}
	return d.stage < another.stage
}

type ParallelArrayProcessor struct {
	Inputs     []any
	Mappers    []func(any) any
	NumWorkers int
}

func (p *ParallelArrayProcessor) Run() (outputs []any, collVals []PAPDebug) {
	perworker := int(len(p.Inputs) / p.NumWorkers)
	if p.NumWorkers*perworker < len(p.Inputs) {
		p.NumWorkers++
	}

	// Initialize outputs first
	for i := range len(p.Inputs) {
		outputs = append(outputs, p.Inputs[i])
	}

	var wg sync.WaitGroup

	barrier := NewBarrier(p.NumWorkers)

	// Collects which ID has worked on which value in which stage
	// There should be numStages * len(p.Inputs) entries
	// So we if we had 3 workers (we would do 4 since 10 % 3 !0 0) and 10 items with 3
	// stages -
	// respective worker would do 3, 3, 3, 1 chunks
	//
	// We should get [0/0, 0/1, 0/2 .. 0/N] [1/0, .. 1/N], [2/0... 2/N], [3/0..3/N]
	collChan := make(chan PAPDebug)
	go func(nvals int) {
		for range nvals {
			val := <-collChan
			collVals = append(collVals, val)
		}
	}(p.NumWorkers * len(p.Inputs))

	worker := func(wid int) {
		defer wg.Done()
		start := wid * perworker
		end := (wid+1)*perworker - 1
		if end >= len(p.Inputs) {
			end = len(p.Inputs) - 1
		}

		// first copy outputs
		for stage, mapper := range p.Mappers {
			barrier.Before() // wait for all threads to arrive here
			for i := start; i <= end; i++ {
				outputs[i] = mapper(outputs[i])
				collChan <- PAPDebug{stage, i, wid}
			}

			barrier.After() // wait for all threads to arrive here
		}
	}

	for wid := range p.NumWorkers {
		wg.Add(1)
		go worker(wid)
	}
	wg.Wait()
	return
}

// Som tests
func TestBasic(t *testing.T) {
	runTest(t, 3, 10, func(i any) any { return i.(int) * 2 }, func(i any) any { return i.(int) + 2 })
	runTest(t, 3, 20, func(i any) any { return i.(int) * 2 }, func(i any) any { return i.(int) + 2 }, func(i any) any { return i.(int) + 5 }, func(i any) any { return i.(int) * 5 })
	runTest(t, 5, 98, func(i any) any { return i.(int) * 2 }, func(i any) any { return i.(int) + 2 }, func(i any) any { return i.(int) + 5 }, func(i any) any { return i.(int) * 5 })
}

func runTest(t *testing.T, numWorkers int, maxInputs int, mappers ...(func(any) any)) {
	var inputs []any
	for i := range maxInputs {
		inputs = append(inputs, i+1)
	}
	log.Println("=============== NumWorkers, NumInputs, NumMappers: ", numWorkers, maxInputs, len(mappers))

	// Run the job serially as well so we have a source of truth

	p := &ParallelArrayProcessor{NumWorkers: numWorkers, Inputs: inputs, Mappers: mappers}
	expected, expdebug := p.SerialRun()
	log.Println("Expected: ", expected)
	log.Println("Expected, Debug: ", expdebug)

	// Do the same with parallel
	found, founddebug := p.Run()
	log.Println("Found: ", found)
	log.Println("found, Debug: ", founddebug)

	assert.Equal(t, expected, found)
	for stage := range len(mappers) {
		expSlice := expdebug[stage*maxInputs : (stage+1)*maxInputs]
		foundSlice := founddebug[stage*maxInputs : (stage+1)*maxInputs]
		sort.Slice(foundSlice, func(i, j int) bool {
			return foundSlice[i].Less(foundSlice[j])
		})
		if !assert.Equal(t, expSlice, foundSlice) {
			log.Println("Stage: ", stage)
			log.Println("ExpectedDebug: ", expSlice)
			log.Println("FoundDebug: ", foundSlice)

		}
	}
}

func (p *ParallelArrayProcessor) SerialRun() (outputs []any, collVals []PAPDebug) {
	numInputs := len(p.Inputs)
	for i := range len(p.Inputs) {
		outputs = append(outputs, p.Inputs[i])
	}

	perworker := int(numInputs / p.NumWorkers)
	for stage, mapper := range p.Mappers {
		for i := range numInputs {
			wid := int(i / perworker)
			outputs[i] = mapper(outputs[i])
			collVals = append(collVals, PAPDebug{stage, i, wid})
		}
	}
	return
}
