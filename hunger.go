package main

import (
	"fmt"
)

// HungerComponent tracks how hungry an entity is.
type HungerComponent struct {
	Hunger int
}

type HungerSystem struct {
	BaseSystem
}

func (hs *HungerSystem) Update(dt float32) {
	for i := 0; i <= len(hs.entities)-1; i++ {
		fmt.Printf("current hunger for %s: %d\n", hs.entities[i].Name, hs.entities[i].Hunger)
		fmt.Printf("incrementing hunger\n")
		hs.entities[i].Hunger++
	}
}
