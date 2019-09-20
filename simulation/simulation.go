package simulation

import (
	"fmt"
	"math/rand"
)

// Aliens is a collection of all Aliens
type Aliens []*Alien

// AlienOccupation maps all Aliens by name
type AlienOccupation map[string]*Alien

// CityDefence maps Aliens by City
type CityDefence map[string]AlienOccupation

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
			if err  := s.MoveAlien(s.Aliens[p]); err != nil {
				return err
			}
		}
		s.Iteration++
	}

	return nil
}

// MoveAlien moves the Alien position in the simulation
func (s *Simulation) MoveAlien(alien *Alien) error {
	fmt.Printf("Moving Alien: %s", alien)
	if alien.IsDead() || alien.IsTrapped() {
		// no-op
		return nil
	}

	from := alien.City
	to := s.PickConnectedCity(alien)
	if from == nil && to == nil {
		// At the beginning
		to = s.PickAnyCity()
	}
	alien.City = to
	if from != nil {
		// Move from City
		delete(s.CityDefence[from.Name], alien.Name)
	}
	// Init city defence
	if (s.CityDefence[to.Name] == nil) {
		s.CityDefence[to.Name] = make(AlienOccupation)
	}
	// Move to City
	s.CityDefence[to.Name][alien.Name] = alien
	if len(s.CityDefence[to.Name]) > 1 {
		for _, a := range s.CityDefence[to.Name] {
			a.Kill();
		}
		to.Destroy()
	}
	return nil
}

// PickConnectedCity picks a random road to undestroyed City
func (s *Simulation) PickConnectedCity(alien *Alien) *City {
	if !alien.IsInvading() {
		return nil
	}
	for _, c := range alien.City.Roads {
		if (!c.IsDestroyed()) {
			return c
		}
	}
	return nil
}

// PickAnyCity picks a random road to undestroyed City
func (s *Simulation) PickAnyCity() *City {
	// TODO: pick random city deterministically
	for _, c := range s.World {
		return c
	}
	return nil
}
