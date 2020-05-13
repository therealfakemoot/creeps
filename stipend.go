package main

import (
// "fmt"
)

type Stipend func(e *Creep)

type StipendSystem struct {
	BaseSystem
	Stipends []Stipend
}

func (ss *StipendSystem) Update(dt float32) {
	for i := 0; i <= len(ss.entities)-1; i++ {
		for _, s := range ss.Stipends {
			s(ss.entities[i])
		}
	}
}

func GoldStipend(e *Creep) {

}

func FoodStipend(e *Creep) {
	e.Hunger--
}
