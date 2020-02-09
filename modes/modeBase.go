package modes

import "reflect"

// ModeBase basic mode struct
type ModeBase struct {
	Mode
}

// Name returns name of mode
func (m *ModeBase) Name() string {
	return reflect.TypeOf(m).String()
}

// Init fake for base struct
func (m *ModeBase) Init() {
	return
}
