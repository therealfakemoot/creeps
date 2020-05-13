package main

import (
	"github.com/EngoEngine/ecs"
)

// Creep is a hungry boi.
type Creep struct {
	ecs.BasicEntity
	HungerComponent
	IdentityComponent
}

// IdentityComponent contains name and psychology data for an entity.
// My thinking is that PsychProfile can just have integers ( positive and negative ) to indicate inclinations towards whatever behaviors.
type IdentityComponent struct {
	Name         string
	PsychProfile map[string]int
}
