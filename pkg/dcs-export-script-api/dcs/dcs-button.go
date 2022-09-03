package DCS

type ButtonState int8

const (
	DefaultUp   ButtonState = 0
	DefaultDown ButtonState = 1
	Unknown     ButtonState = 99
)

type Button struct {
	id            int
	defaultValue  ButtonState
	expectedValue ButtonState
	currentValue  ButtonState
}

// Press send button down signal through the UDP socket until
// receives a button down signal then send button up signal util receives a button is up signal
func (b Button) Press(s ButtonState) {
	b.expectedValue = s
}

func (b Button) Reset() {
	b.expectedValue = b.defaultValue
}
