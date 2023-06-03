package cor

// Department 处理者接口
type Department interface {
	Execute(*Patient)
	SetNext(department Department)
}
