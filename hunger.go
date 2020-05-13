package main

import ()

// HungerComponent tracks how hungry an entity is.
type HungerComponent struct {
	Hunger int
}

type HungerSystem struct {
	BaseSystem
}

func (hs *HungerSystem) Update(dt float32) {
	for i := 0; i <= len(hs.entities)-1; i++ {
		hs.entities[i].Hunger++
	}
}
