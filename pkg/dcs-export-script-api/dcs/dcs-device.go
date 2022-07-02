package dcs

type DcsDevice struct {
	id   int
	name string
}

func (d DcsDevice) PressButton(id int, v float32) {
	PressButton(d.id, id, v)
}
