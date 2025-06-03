package domain

type AggregateRoot interface {
	GetID() string
	GetVersion() int64
	GetDomainEvents() []DomainEvent
	ClearDomainEvents()
	RaiseDomainEvent(event DomainEvent)
}

type BaseAggregateRoot struct {
	BaseEntity
	domainEvents []DomainEvent
}

func (a *BaseAggregateRoot) GetDomainEvents() []DomainEvent {
	return a.domainEvents
}

func (a *BaseAggregateRoot) ClearDomainEvents() {
	a.domainEvents = nil
}

func (a *BaseAggregateRoot) RaiseDomainEvent(event DomainEvent) {
	a.domainEvents = append(a.domainEvents, event)
}
