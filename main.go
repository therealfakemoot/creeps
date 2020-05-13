package main

import (
	"flag"

	"github.com/EngoEngine/ecs"
	// "log"
	// "time"
)

func main() {
	var w ecs.World
	rs := &RenderSystem{}
	hs := &HungerSystem{}
	ds := &DecisionSystem{}
	w.AddSystem(rs)
	w.AddSystem(hs)
	w.AddSystem(ds)

	bob := Creep{
		BasicEntity:       ecs.NewBasic(),
		HungerComponent:   HungerComponent{Hunger: 5, HungerFactor: 2},
		IdentityComponent: IdentityComponent{Name: "Bob"},
	}
	rs.Add(&bob)
	hs.Add(&bob)
	ds.Add(&bob)

	john := Creep{
		BasicEntity:       ecs.NewBasic(),
		HungerComponent:   HungerComponent{Hunger: 5, HungerFactor: 3},
		IdentityComponent: IdentityComponent{Name: "John"},
	}
	rs.Add(&john)
	hs.Add(&john)
	ds.Add(&john)

	var (
		ticks int
	)

	flag.IntVar(&ticks, "ticks", 10, "number of ticks to simulate")
	flag.Parse()
	// run the world for 10 seconds
	for i := 0; i < ticks; i++ {
		w.Update(1)
	}

}
