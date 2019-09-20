package simulation

import (
	"errors"
	"math/rand"
)

const (
	// DefaultIterations is the default used if number of iteration is not othervise specified
	DefaultIterations int = 10000
)

// Aliens is a collection of all Aliens
type Aliens []*Alien

// AlienOccupation maps all Aliens by name
type AlienOccupation map[string]*Alien

// CityDefence maps Aliens by City
type CityDefence map[string][]*Alien

// Simulation represents a running simulation
type Simulation struct {
	R *rand.Rand
	Iteration int
	EndIteration int

	World
	Aliens
	AlienOccupation
	CityDefence
}

// NewSimulation inits a new Simulation instance
func NewSimulation(r *rand.Rand, endIteration int, world World, aliens Aliens) Simulation {
	return Simulation{
		R: r, Iteration: 0,
		EndIteration: endIteration,
		World: world,
		Aliens: aliens,
		AlienOccupation: make(AlienOccupation),
		CityDefence: make(CityDefence),
	}
}

// Start the simulation
func (in Simulation) Start() error {
	return errors.New("Not implemented (Yet!)")
}

// MoveAlien moves the Alien position in the simulation
func (in Simulation) MoveAlien(alien Alien) error {
	return errors.New("Not implemented (Yet!)")
}
