package mediator

type Mediator interface {
	CanArrive(train Train) bool
	NotifyAboutDeparture()
}
