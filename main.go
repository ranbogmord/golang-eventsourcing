package main

import (
	"fmt"
	"es/events"
	"es/aggregates"
	"es/repositories"
)

func main() {
	repo := repositories.MemoryEventRepository{Events: []events.Event{}}

	createEvt := events.NewRecipeCreatedEvent("Test recipe")
	repo.Store(createEvt)
	repo.Store(events.NewIngredientAddedEvent(createEvt.AggregateID(), "Ris", 3, "dl"))
	repo.Store(events.NewIngredientAddedEvent(createEvt.AggregateID(), "Kyckling", 300, "grams"))
	repo.Store(events.NewIngredientAddedEvent(createEvt.AggregateID(), "Wookmix", 200, "grams"))

	r := &aggregates.Recipe{}
	rr, err := repo.Load(r, "1234")
	if err != nil {
		panic(err)
	}

	r = rr.(*aggregates.Recipe)

	fmt.Printf(`ID: %s
Name: %s
Created at: %s
Ingredients:
%v
`, r.Id, r.Name, r.CreatedAt.Format("2006-01-02 15:04:05"), r.Ingredients)
}