package dcs

type ButtonState float32

const (
	DefaultUp   ButtonState = 0
	DefaultDown ButtonState = 1
	Unknown     ButtonState = 99
)

type DcsButton struct {
	id           int
	defaultValue ButtonState
	upValue      ButtonState
	downValue    ButtonState
	currentValue ButtonState
}

// Press send a button down signal through the UDP socket until
// receives a button down signal then send a button up signal util receives a button is up signal
func (b DcsButton) Press() {}

func (b DcsButton) Release() {}

func (b DcsButton) Reset() {}
