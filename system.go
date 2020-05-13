package main

import (
	"github.com/EngoEngine/ecs"
)

// BaseSystem helps centralize common functionality in systems.
type BaseSystem struct {
	entities []*Creep
}

// Add appends an entity to the system
func (bs *BaseSystem) Add(e *Creep) {
	bs.entities = append(bs.entities, e)
}

// Remove pops an entity from the rendering system.
func (bs *BaseSystem) Remove(e ecs.BasicEntity) {
	id := e.ID()

	for idx, ent := range bs.entities {
		if ent.ID() == id {
			bs.entities = append(bs.entities[idx:], bs.entities[idx:]...)
		}
	}
}
