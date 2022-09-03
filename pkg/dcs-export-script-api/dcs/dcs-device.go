package DCS

type Device struct {
	id      int
	name    string
	buttons map[int]Button
}

func (d Device) PressButton(id int, v float32) {
	PressButton(d.id, id, v)
}
