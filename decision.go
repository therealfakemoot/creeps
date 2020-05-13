package main

import (
	"fmt"
)

type DecisionSystem struct {
	BaseSystem
	Deciders []Decider
}

func (ds *DecisionSystem) Update(dt float32) {
	for i := 0; i <= len(ds.entities)-1; i++ {
		ds.Decide(ds.entities[i])
	}
}

type Decider func(e *Creep)

func HungerDecider(e *Creep) {
	switch {
	case e.Hunger <= 10:
		return
	case e.Hunger <= 20:
		fmt.Printf("%s is mildly hungry\n", e.Name)
	case e.Hunger <= 30:
		fmt.Printf("%s is hungry\n", e.Name)
	case e.Hunger <= 40:
		fmt.Printf("%s is starving\n", e.Name)
	default:
		fmt.Printf("%s is ready for cannibalism\n", e.Name)

	}
}

func (ds *DecisionSystem) Decide(e *Creep) {
	for _, d := range ds.Deciders {
		d(e)
	}
}
