package aggregates

import (
	"time"
	"errors"
	"fmt"
	"es/events"
)

type Aggregate interface {
	On(evt events.Event) error
	Apply (history []events.Event) error
}

type Ingredient struct {
	Name string
	Amount float64
	Unit string
}

type Recipe struct {
	Aggregate
	Id string
	CreatedAt time.Time
	Name string
	Ingredients []Ingredient
}

func (r *Recipe) Apply(history []events.Event) error {
	for _, v := range history {
		err := r.On(v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Recipe) On(evt events.Event) error {
	switch v := evt.(type) {
	case events.RecipeCreated:
		r.Id = v.AggregateID()
		r.CreatedAt = v.At()
		r.Name = v.Name
		r.Ingredients = []Ingredient{}
	case events.IngredientAdded:
		i := Ingredient{Name: v.Name, Amount: v.Amount, Unit: v.Unit}
		r.Ingredients = append(r.Ingredients, i)
	default:
		err := errors.New(fmt.Sprintf("invalid event type: %v", v))
		return err
	}
	return nil
}