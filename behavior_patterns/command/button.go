package command

// 请求者：button.go

type Button struct {
	Command Command
}

func (b *Button) Press() {
	b.Command.execute()
}
