package main

import (
	"fmt"
)

type Decider func(e *Creep)

func HungerDecider(e *Creep) {
	switch {
	case e.Hunger < 0:
		fmt.Printf("%s is real sick(%d)\n", e.Name, e.Hunger)
	case e.Hunger <= 10:
		fmt.Printf("%s is satiated(%d)\n", e.Name, e.Hunger)
	case e.Hunger <= 20:
		fmt.Printf("%s is mildly hungry(%d)\n", e.Name, e.Hunger)
	case e.Hunger <= 30:
		fmt.Printf("%s is hungry(%d)\n", e.Name, e.Hunger)
	case e.Hunger <= 40:
		fmt.Printf("%s is starving(%d)\n", e.Name, e.Hunger)
	default:
		fmt.Printf("%s is ready for cannibalism(%d)\n", e.Name, e.Hunger)

	}
}
