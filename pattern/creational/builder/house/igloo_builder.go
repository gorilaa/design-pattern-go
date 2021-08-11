package house

type IglooBuilder struct {
	WindowType, DoorType string
	Floor                int
}

func NewIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) SetWindowType() {
	b.WindowType = "Wooden Window"
}

func (b *IglooBuilder) SetDoorType() {
	b.DoorType = "Wooden Door"
}

func (b *IglooBuilder) SetNumFloor() {
	b.Floor = 2
}

func (b *IglooBuilder) GetHouse() House {
	return House{
		WindowType: b.WindowType,
		DoorType:   b.DoorType,
		Floor:      b.Floor,
	}
}
