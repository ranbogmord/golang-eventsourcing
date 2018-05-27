package repositories

import (
	"es/events"
	"es/aggregates"
)

type Repository interface {
	Store(events.Event)
	Load(agg aggregates.Aggregate, id string) (aggregates.Aggregate, error)
}

type MemoryEventRepository struct {
	Events []events.Event
}

func (r *MemoryEventRepository) Store(evt events.Event) {
	r.Events = append(r.Events, evt)
}

func (r *MemoryEventRepository) Load(agg aggregates.Aggregate, id string) (aggregates.Aggregate, error) {
	for _, e := range r.Events {
		if e.AggregateID() == id {
			err := agg.On(e)
			if err != nil {
				return nil, err
			}
		}
	}

	return agg, nil
}
