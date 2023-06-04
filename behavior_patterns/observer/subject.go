package observer

//主体

type Subject interface {
	Register(observer Observer)
	deregister(observer Observer)
	notifyAll()
}
