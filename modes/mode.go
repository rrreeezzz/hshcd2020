package modes

// Mode describe a mode
type Mode interface {
	// Run runs the mode
	Run(int, int, []int) (int, []int)
}
