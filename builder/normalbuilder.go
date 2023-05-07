package builder

type NormalBuilder struct {
	windowsType string
	doorType    string
	floor       int
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) setWindowType() {
	b.windowsType = "Wooden Windows"
}

func (b *NormalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *NormalBuilder) setNumFloor() {
	b.floor = 2
}

func (b *NormalBuilder) getHouse() House {
	return House{
		WindowsType: b.doorType,
		DoorType:    b.windowsType,
		Floor:       b.floor,
	}
}
