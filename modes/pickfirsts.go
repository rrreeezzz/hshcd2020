package modes

// Pickfirsts mode
type Pickfirsts struct {
	ModeBase
}

// NewPickfirsts allocates new pickfirsts
func NewPickfirsts() *Pickfirsts {
	return &Pickfirsts{}
}

// Name returns name of mode
func (m *Pickfirsts) Name() string {
	return "Pickfirsts"
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

	return r, pizOut
}
