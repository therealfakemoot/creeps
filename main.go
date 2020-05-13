package main

import (
	"github.com/EngoEngine/ecs"
	// "log"
	// "time"
)

func main() {
	var w ecs.World
	rs := &RenderSystem{}
	hs := &HungerSystem{}
	w.AddSystem(rs)
	w.AddSystem(hs)

	bob := Creep{
		BasicEntity:       ecs.NewBasic(),
		HungerComponent:   HungerComponent{Hunger: 5},
		IdentityComponent: IdentityComponent{Name: "Bob"},
	}
	rs.Add(&bob)
	hs.Add(&bob)

	// run the world for 10 seconds
	for i := 0; i < 10; i++ {
		w.Update(1)
	}

}
