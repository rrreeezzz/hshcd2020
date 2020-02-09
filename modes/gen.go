package modes

// Gen mode
type Gen struct {
	ModeBase
}

// NewGen allocates new Gen
func NewGen() *Gen {
	return &Gen{}
}

// Name returns name of mode
func (m *Gen) Name() string {
	return "Gen"
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Run runs
func (m *Gen) Run(max, num int, pizSizes []int) (int, []int) {

	var slices int
	var pizOut []int

	// use a map, no duplicate this way
	prevGen := make(map[int][]int)
	nextGen := make(map[int][]int)
	// generation 1
	for i, v := range pizSizes {
		if _, ok := prevGen[v]; ok {
			continue
		}
		prevGen[v] = append(prevGen[v], i)
	}

	// TODO: refactor, uses goroutines, channels... make it faster
	// there should be maximum num-1 generations
	// not reproducible, because maps not ordered
	for it := 1; it < num-1; it++ {
		first := true
		// go through all previous gen and add pizzas
		for sum, idxs := range prevGen {
			for i, v := range pizSizes {
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
	slices = 0
	for sum, idx := range nextGen {
		if sum > slices {
			slices = sum
			pizOut = idx
		}
	}

	return slices, pizOut
}
