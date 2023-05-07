package builder

type IglooBuilder struct {
	windowsType string
	doorType    string
	floor       int
}

func NewIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) setWindowType() {
	b.windowsType = "Snow Windows"
}

func (b *IglooBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *IglooBuilder) setNumFloor() {
	b.floor = 1
}

func (b *IglooBuilder) getHouse() House {
	return House{
		WindowsType: b.doorType,
		DoorType:    b.windowsType,
		Floor:       b.floor,
	}
}
