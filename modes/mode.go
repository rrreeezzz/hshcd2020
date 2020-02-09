package modes

// Mode describe a mode
type Mode interface {
	// Run runs the mode
	// returns resulting number of slices and list of selected pizzas
	Run(int, int, []int) (int, []int)

	// Name returns name of the mode
	Name() string

	//TODO: time function
	Init()
}
