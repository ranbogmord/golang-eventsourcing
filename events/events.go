package events

import "time"

type Event interface {
	AggregateID() string
	At() time.Time
}
type EventModel struct {
	Event
	aggregateId string
	at time.Time
}

func (m EventModel) AggregateID() string {
	return m.aggregateId
}

func (m EventModel) At() time.Time {
	return m.at
}

type RecipeCreated struct {
	EventModel
	Name string
}

func NewRecipeCreatedEvent(name string) RecipeCreated {
	return RecipeCreated{
		EventModel: EventModel{aggregateId: "1234", at: time.Now()},
		Name: name,
	}
}

type IngredientAdded struct {
	EventModel
	Name string
	Amount float64
	Unit string
}

func NewIngredientAddedEvent(aggregateId string, name string, amount float64, unit string) IngredientAdded {
	return IngredientAdded{
		EventModel: EventModel{aggregateId: aggregateId, at: time.Now()},
		Name: name,
		Amount: amount,
		Unit: unit,
	}
}
