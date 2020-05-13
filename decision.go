package main

import ()

type DecisionSystem struct {
	BaseSystem
	Deciders []Decider
}

func (ds *DecisionSystem) Update(dt float32) {
	for i := 0; i <= len(ds.entities)-1; i++ {
		ds.Decide(ds.entities[i])
	}
}

func (ds *DecisionSystem) Decide(e *Creep) {
	for _, d := range ds.Deciders {
		d(e)
	}
}
