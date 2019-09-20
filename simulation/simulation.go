package simulation

import (
	"fmt"
	"errors"
	"math/rand"
)

// Aliens is a collection of all Aliens
type Aliens []*Alien

// AlienOccupation maps all Aliens by name
type AlienOccupation map[string]*Alien

// CityDefence maps Aliens by City
type CityDefence map[string][]*Alien

// Simulation struct represents a running simulation
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
		R: r,
		Iteration: 0,
		EndIteration: endIteration,
		World: world,
		Aliens: aliens,
		AlienOccupation: make(AlienOccupation),
		CityDefence: make(CityDefence),
	}
}

// Start the simulation
func (s *Simulation) Start() error {
	fmt.Printf("Running simulation for %d iterations\n", s.EndIteration)
	for s.Iteration < s.EndIteration {
		fmt.Printf("\nStart Iteration: %d\n", s.Iteration)
		fmt.Print("----------------\n")
		picks := MakeRange(0, len(s.Aliens))
		Shuffle(picks, s.R)
		for _, p := range picks {
			s.MoveAlien(s.Aliens[p])
		}
		s.Iteration++
	}

	return errors.New("Not implemented (Yet!)")
}

// MoveAlien moves the Alien position in the simulation
func (s *Simulation) MoveAlien(alien *Alien) error {
	fmt.Printf("Moving Alien: %s\n", alien)
	return errors.New("Not implemented (Yet!)")
}
