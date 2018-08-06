package sim

import (
	"math"
)

type DemoSimulation struct {
	Simulation

	time float64
}

func NewDemoSimulation() *DemoSimulation {
	return &DemoSimulation{}
}

// Tick advances time slightly
func (sim *DemoSimulation) Tick(deltaTime float64) {
	sim.time += deltaTime
}

// Add is noop
func (sim *DemoSimulation) Add(m Mass) {
	return
}

// Masses returns some cool looking stuff
func (sim *DemoSimulation) WriteFrame(visitor func(Mass) bool) {
	for x := -200.0; x < 200; x += 1 {
		for ofs := 0.0; ofs < 4; ofs += 0.04 {
			val := math.Sin(sim.time + ofs + (x / 100.0))
			m := Mass{Pos: Vec2{x, val * 100}}

			if !visitor(m) {
				return
			}
		}
	}
}
