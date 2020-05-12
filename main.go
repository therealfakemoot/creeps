package main

import (
	// "log"
	"github.com/EngoEngine/ecs"
	// "time"
)

func main() {
	var w ecs.World
	var rs RenderSystem
	w.AddSystem(rs)

	bob := Creep{BasicEntity: ecs.NewBasic()}
	rs.Add(bob)

	// run the world for 10 seconds
	for i := 0; i < 10; i++ {
		w.Update(1)
	}

}
