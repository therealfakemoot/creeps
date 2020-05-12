package main

import (
	"fmt"

	"github.com/EngoEngine/ecs"
)

// RenderSystem prints the state of the rendered Entities to the terminal
type RenderSystem struct {
	ticks    int
	entities []Creep
}

// Add appends an entity to the system
func (rs RenderSystem) Add(e Creep) {
	rs.entities = append(rs.entities, e)
}

// Update loops over
func (rs RenderSystem) Update(dt float32) {
	fmt.Printf("%#v\n", rs)
	for _, e := range rs.entities {
		fmt.Printf("e has id %d", e.ID())
	}
}

// Remove pops an entity from the rendering system.
func (rs RenderSystem) Remove(e ecs.BasicEntity) {
	id := e.ID()

	for idx, ent := range rs.entities {
		b := ent.GetBasicEntity()
		if b.ID() == id {
			rs.entities = append(rs.entities[idx:], rs.entities[idx:]...)
		}
	}
}
