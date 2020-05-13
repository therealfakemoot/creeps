package main

import (
// "fmt"
)

// RenderSystem prints the state of the rendered Entities to the terminal
type RenderSystem struct {
	BaseSystem
	ticks int
}

// Update loops over
func (rs *RenderSystem) Update(dt float32) {
	for i := 0; i <= len(rs.entities)-1; i++ {
		// fmt.Printf("%s has hunger %d\n", rs.entities[i].Name, rs.entities[i].Hunger)
	}
}
