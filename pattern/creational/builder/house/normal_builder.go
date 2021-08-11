package house

type NormalBuilder struct {
	WindowType, DoorType string
	Floor                int
}

func NewNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) SetWindowType() {
	b.WindowType = "Wooden Window"
}

func (b *NormalBuilder) SetDoorType() {
	b.DoorType = "Wooden Door"
}

func (b *NormalBuilder) SetNumFloor() {
	b.Floor = 2
}

func (b *NormalBuilder) GetHouse() House {
	return House{
		WindowType: b.WindowType,
		DoorType:   b.DoorType,
		Floor:      b.Floor,
	}
}
