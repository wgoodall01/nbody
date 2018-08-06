package sim

// Simulation defines an n-body simulation, and methods to tick it forward
// through time.
type Simulation interface {
	// Tick advances the simulation by an amount of time specified by t.
	Tick(deltaTime float64)

	// Add adds a mass to the simulation.
	Add(m Mass)

	// WriteFrame calls the visitor function once with each Mass, stopping if it returns false.
	WriteFrame(func(Mass) bool)
}

// Mass represents an object, with a position, velocity, and mass.
type Mass struct {
	// Pos contains the position of the mass
	Pos Vec2

	// Vel contains the velocity of the mass. Always {0,0} for static masses.
	Vel Vec2

	// Mass of the object
	Mass float64

	// Static = true if the mass is fixed in space
	Static bool
}
