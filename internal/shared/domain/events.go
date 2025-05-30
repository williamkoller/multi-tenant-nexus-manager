package domain

import (
	"time"

	"github.com/google/uuid"
)

type DomainEvent interface {
	GetEventID() string
	GetEventType() string
	GetAggregateID() string
	GetOccurredOn() time.Time
	GetEventData() interface{}
}

type BaseDomainEvent struct {
	EventID     string      `json:"event_id"`
	EventType   string      `json:"event_type"`
	AggregateID string      `json:"aggregate_id"`
	OccurredOn  time.Time   `json:"occurred_on"`
	EventData   interface{} `json:"event_data"`
}

func NewBaseDomainEvent(eventType, aggregateID string, eventData interface{}) BaseDomainEvent {
	return BaseDomainEvent{
		EventID:     uuid.New().String(),
		EventType:   eventType,
		AggregateID: aggregateID,
		OccurredOn:  time.Now().UTC(),
		EventData:   eventData,
	}
}

func (e BaseDomainEvent) GetEventID() string        { return e.EventID }
func (e BaseDomainEvent) GetEventType() string      { return e.EventType }
func (e BaseDomainEvent) GetAggregateID() string    { return e.AggregateID }
func (e BaseDomainEvent) GetOccurredOn() time.Time  { return e.OccurredOn }
func (e BaseDomainEvent) GetEventData() interface{} { return e.EventData }
