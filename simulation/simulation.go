package simulation

import (
	"fmt"
	"math/rand"

	"alien-invasion/types"
)

type (
	// Alien attacking City
	Alien = types.Alien
	// City connected to other Cities with Roads
	City = types.City
	// World map of Cities
	World = types.World
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

// NoOpError when next move can not be made
type NoOpError struct {
	reason int
}

const (
	// NoOpAlienDisabled when Alien Dead or Trapped
	NoOpAlienDisabled int = 1
	// NoOpWorldDestroyed when 
	NoOpWorldDestroyed int = 2
)

// Error string representation
func (err *NoOpError) Error() string {
	return fmt.Sprintf("Simulator no-op with reason: %d", err.reason)
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
		fmt.Printf("\n\n")
		fmt.Printf("Iteration: %d\n", s.Iteration)
		fmt.Print("----------\n")
		// Shuffle cards every iteration
		picks := MakeRange(0, len(s.Aliens))
		Shuffle(picks, s.R)
		// Aliens make their moves
		noOpMoves := 0
		for _, p := range picks {
			if err := s.MoveAlien(s.Aliens[p]); err != nil {
				// Count no-op
				if _, ok := err.(*NoOpError); ok {
					noOpMoves++
					continue;
				}
				return err
			}
		}
		// Check if last iteration was empty (no moves)
		if noOpMoves == len(s.Aliens) {
			fmt.Printf("\n")
			fmt.Printf("Simulation ended early on iteration %d: No Available Moves", s.Iteration)
			return nil
		}
		// Next round
		s.Iteration++
	}
	// Game Over
	return nil
}

// MoveAlien moves the Alien position in the simulation
func (s *Simulation) MoveAlien(alien *Alien) error {
	fmt.Printf("Moving Alien: %s", alien)
	if alien.IsDead() || alien.IsTrapped() {
		// no-op
		fmt.Print(" => To: NO move! Alien Dead or Trapped.\n")
		return &NoOpError{NoOpAlienDisabled}
	}
	// Move
	from := alien.City
	to := s.PickConnectedCity(alien)
	if from == nil && to == nil {
		// At the beginning
		to = s.PickAnyCity()
		if to == nil {
			// no-op
			fmt.Print(" => To: NO move! World is destroyed.\n")
			return &NoOpError{reason: NoOpWorldDestroyed}
		}
	}
	fmt.Printf(" => To: %s\n", to)
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
		to.Destroy()
		// Kill Aliens and notify
		out := fmt.Sprintf(" || %s has been destroyed by ", to.Name)
		for _, a := range s.CityDefence[to.Name] {
			out += fmt.Sprintf("alien %s and ", a.Name)
			a.Kill();
		}
		out = out[:len(out) - 5] + "!\n"
		fmt.Print(out)
	}
	// Done
	return nil
}

// PickConnectedCity picks a random road to undestroyed City
func (s *Simulation) PickConnectedCity(alien *Alien) *City {
	// Nil if still not invading
	if !alien.IsInvading() {
		return nil
	}
	// Any undestroyed connected city
	// TODO: pick connected city deterministically
	for _, c := range alien.City.Roads {
		if (!c.IsDestroyed()) {
			return c
		}
	}
	return nil
}

// PickAnyCity picks a random road to undestroyed City
func (s *Simulation) PickAnyCity() *City {
	// Any undestroyed city
	// TODO: pick random city deterministically
	for _, c := range s.World {
		if (!c.IsDestroyed()) {
			return c
		}
	}
	return nil
}
