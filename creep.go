package main

import (
	"github.com/EngoEngine/ecs"
)

// Creep is a hungry boi.
type Creep struct {
	ecs.BasicEntity
	HungerComponent
}

// HungerComponent tracks how hungry an entity is.
type HungerComponent struct {
	hunger int
}

// Get returns the hunger value
func (hc *HungerComponent) Get() int {
	return hc.hunger
}

// Add increases hunger by the given amount
func (hc *HungerComponent) Add(n int) {
	hc.hunger += n
}

// Sub decreases hunger by the given amount
func (hc *HungerComponent) Sub(n int) {
	hc.hunger -= n
}
