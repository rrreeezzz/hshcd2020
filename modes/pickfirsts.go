package modes

// Pickfirsts mode
type Pickfirsts struct {
}

// NewPickFirsts allocates new pickfirsts
func NewPickFirsts() *Pickfirsts {
	return &Pickfirsts{}
}

// Run runs
func (m *Pickfirsts) Run(max, num int, pizSizes []int) (int, []int) {

	var r int
	var pizOut []int
	for i, p := range pizSizes {
		if r+p > max {
			break
		}
		r = r + p
		pizOut = append(pizOut, i)
	}

	return len(pizOut), pizOut
}
