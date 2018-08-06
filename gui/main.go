package main

import (
	"math/rand"
	"time"

	. "github.com/wgoodall01/nbody/sim"
)

func randRange(min, max float64) float64 {
	diff := max - min
	return min + rand.Float64()*diff
}

func main() {
	sim := NewNaiveSimulation(10000, 10000, 1)

	ticker := time.Tick(time.Second / 10)
	go func() {
		for range ticker {
			sim.Add(Mass{
				Pos:  Vec2{randRange(-500, 500), randRange(-500, 500)},
				Mass: randRange(0, 10),
			})
		}
	}()

	gui := NewGui(sim)
	gui.Start()
}
