package DCS

type Device struct {
	id   int
	name string
}

func (d Device) PressButton(id int, v float32) {
	PressButton(d.id, id, v)
}
