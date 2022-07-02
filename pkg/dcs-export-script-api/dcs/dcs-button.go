package DCS

type ButtonState float32

const (
	DefaultUp   ButtonState = 0
	DefaultDown ButtonState = 1
	Unknown     ButtonState = 99
)

type Button struct {
	id           int
	defaultValue ButtonState
	upValue      ButtonState
	downValue    ButtonState
	currentValue ButtonState
}

// Press send a button down signal through the UDP socket until
// receives a button down signal then send a button up signal util receives a button is up signal
func (b Button) Press() {}

func (b Button) Release() {}

func (b Button) Reset() {}
