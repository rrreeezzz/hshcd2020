package modes

// Picklasts mode
type Picklasts struct {
	ModeBase
}

// NewPicklasts allocates new Picklasts
func NewPicklasts() *Picklasts {
	return &Picklasts{}
}

// Name returns name of mode
func (m *Picklasts) Name() string {
	return "Picklasts"
}

// Run runs
func (m *Picklasts) Run(max, num int, pizSizes []int) (int, []int) {

	var r int
	var pizOut []int
	for i := len(pizSizes) - 1; i >= 0; i-- {
		if r+pizSizes[i] > max {
			break
		}
		r = r + pizSizes[i]
		pizOut = append(pizOut, i)
	}

	return r, pizOut
}
