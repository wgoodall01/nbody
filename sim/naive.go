package sim

import (
	"math"
	"sync"
)

type NaiveSimulation struct {
	Simulation
	mu sync.RWMutex

	frontBuffer []Mass
	backBuffer  []Mass

	// Gravitational constant.
	g float64

	// Radius of each mass.
	// This will merge particles inside this range.
	radius float64
}

func NewNaiveSimulation(capacity int, g float64, radius float64) *NaiveSimulation {
	return &NaiveSimulation{
		frontBuffer: make([]Mass, 0, capacity),
		backBuffer:  make([]Mass, 0, capacity),
		g:           g,
		radius:      radius,
	}
}

// swapBuffers promotes the back to the front, then clears the back.
func (sim *NaiveSimulation) swapBuffers() {
	sim.frontBuffer = sim.backBuffer
	sim.backBuffer = sim.backBuffer[:0]
}

// Tick calculates another frame of the simulation.
func (sim *NaiveSimulation) Tick(deltaTime float64) {
	sim.mu.Lock()
	defer sim.mu.Unlock()

	// This is O(n^2) and pretty ehh.
	for _, m1 := range sim.frontBuffer {
		// Don't move any static masses.
		if !m1.Static {
			for _, m2 := range sim.frontBuffer {

				displacement := m2.Pos.Sub(m1.Pos)
				angle := displacement.Normalize()
				r := displacement.Norm() // Distance between the two masses

				// Skip if inside radius.
				if r < sim.radius {
					continue
				}

				magFg := sim.g * (m1.Mass * m2.Mass) / math.Pow(r, 2) // Newton's universal gravitation ( G*m1*m2/r^2 )
				Fg := angle.Scale(magFg)                              // Force vector
				a := Fg.Div(m1.Mass)                                  // Newton's second (F=ma)
				m1.Vel = m1.Vel.Add(a.Scale(deltaTime))               // integrate Fg with m1's velocity, (v = at)
			}
		}

		// Integrate m1 velocity into position (x = vt)
		m1.Pos = m1.Pos.Add(m1.Vel.Scale(deltaTime))

		// Save updated mass to back buffer.
		sim.backBuffer = append(sim.backBuffer, m1)
	}

	sim.swapBuffers()
}

// Add adds a mass to the simulation.
func (sim *NaiveSimulation) Add(m Mass) {
	sim.mu.Lock()
	sim.frontBuffer = append(sim.frontBuffer, m)
	sim.mu.Unlock()
}

// WriteFrame writes a frame of the simulation through visitor.
func (sim *NaiveSimulation) WriteFrame(visitor func(Mass) bool) {
	sim.mu.RLock()
	defer sim.mu.RUnlock()

	for _, mass := range sim.frontBuffer {
		if !visitor(mass) {
			break
		}
	}
}
