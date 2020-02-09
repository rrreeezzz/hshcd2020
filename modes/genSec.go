package modes

import (
	"runtime"
	"sync"
)

// GenSec mode
type GenSec struct {
	ModeBase

	rchan chan bestTuple
}

type bestTuple struct {
	slices  int
	pizzOut []int
}

// NewGenSec allocates new GenSec
func NewGenSec() *GenSec {
	return &GenSec{
		rchan: make(chan bestTuple),
	}
}

// Name returns name of mode
func (m *GenSec) Name() string {
	return "GenSec"
}

func (m *GenSec) worker(wg *sync.WaitGroup, max, num, start, step int, pizSizes []int) {
	defer wg.Done()

	// use a map, no duplicate this way
	prevGen := make(map[int][]int)
	nextGen := make(map[int][]int)
	// generation 1 + steps
	newPizSizes := []int{}
	for i := start; i < num; i += step {
		newPizSizes = append(newPizSizes, i)
		if _, ok := prevGen[pizSizes[i]]; ok {
			continue
		}
		prevGen[pizSizes[i]] = append(prevGen[pizSizes[i]], i)
	}

	for it := 1; it < len(newPizSizes)-1; it++ {
		first := true
		// go through all previous gen and add pizzas
		for sum, idxs := range prevGen {
			for i, v := range newPizSizes {
				// edge cases
				if contains(idxs, i) {
					continue
				}
				tmpsum := sum + v
				if tmpsum > max {
					continue
				}
				if first {
					first = false
					nextGen = make(map[int][]int)
				}
				if _, exists := nextGen[tmpsum]; exists {
					continue
				}
				nextGen[tmpsum] = append(prevGen[sum], i)
			}
		}
		prevGen = nextGen
	}

	// Select best tuple (slices, [pizzas])
	bestTuple := bestTuple{}
	for sum, idx := range nextGen {
		if sum > bestTuple.slices {
			bestTuple.slices = sum
			bestTuple.pizzOut = idx
		}
	}

	m.rchan <- bestTuple
}

// Run runs
func (m *GenSec) Run(max, num int, pizSizes []int) (int, []int) {

	var wg sync.WaitGroup

	// Set GOMAXPROCS
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	for grs := 0; grs < cpus; grs++ {
		wg.Add(1)
		go m.worker(&wg, max, num, grs, cpus, pizSizes)
	}

	wg.Wait()

	finalTuple := &bestTuple{}
	select {
	case tmp := <-m.rchan:
		finalTuple.slices += tmp.slices
		finalTuple.pizzOut = append(finalTuple.pizzOut, tmp.pizzOut...)
	default:
		break
	}

	// TODO: refactor, uses goroutines, channels... make it faster
	// there should be maximum num-1 generations
	// not reproducible, because maps not ordered

	return finalTuple.slices, finalTuple.pizzOut
}
