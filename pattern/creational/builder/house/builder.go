package house

type Builder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() House
}

func GetBuilder(builderType string) Builder {
	if builderType == "normal" {
		return &NormalBuilder{}
	}

	if builderType == "igloo" {
		return &IglooBuilder{}
	}
	return nil
}
