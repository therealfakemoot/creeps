package main

import ()

// HungerComponent tracks how hungry an entity is.
type HungerComponent struct {
	Hunger       int
	HungerFactor int
}

// HungerSystem handles actions based on hunger.
type HungerSystem struct {
	BaseSystem
}

// Update will increment hunger for all entitites
func (hs *HungerSystem) Update(dt float32) {
	for i := 0; i <= len(hs.entities)-1; i++ {
		hs.entities[i].Hunger += hs.entities[i].HungerFactor
	}
}
