package simulation

import (
	"fmt"
	"sort"
	"math/rand"

	"alien-invasion/types"
)

type (
	// Alien attacking City
	Alien = types.Alien
	// City connected to other Cities with RoadsMap
	City = types.City
	// World map of Cities
	World = types.World
)

// Aliens is a collection of all Aliens
type Aliens []*Alien

// AlienOccupation maps all Aliens by name
type AlienOccupation map[string]*Alien

// CityDefense maps Aliens by City
type CityDefense map[string]AlienOccupation

// Simulation struct represents a running simulation
type Simulation struct {
	// Simulation config
	R *rand.Rand
	Iteration int
	EndIteration int
	// World state
	World
	Aliens
	CityDefense
}

// NoOpError when next move can not be made
type NoOpError struct {
	reason int
}

const (
	// NoOpAlienDisabled when Alien Dead or Trapped
	NoOpAlienDisabled int = 1
	// NoOpWorldDestroyed when World destroyed
	NoOpWorldDestroyed int = 2
	// NoOpMessage when no-op
	NoOpMessage = " => To: NO move! %s\n"
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
		CityDefense: make(CityDefense),
	}
}

// Start the simulation
func (s *Simulation) Start() error {
	fmt.Printf("Running simulation for %d iterations\n", s.EndIteration)
	for ; s.Iteration < s.EndIteration; s.Iteration++ {
		fmt.Printf("\n\n")
		fmt.Printf("Iteration: %d\n", s.Iteration)
		fmt.Print("----------\n")
		// Shuffle cards every iteration
		picks := ShuffleLen(len(s.Aliens), s.R)
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
	}
	// Game Over
	return nil
}

// MoveAlien moves the Alien position in the simulation
func (s *Simulation) MoveAlien(alien *Alien) error {
	fmt.Printf("Moving Alien: %s", alien)
	// Check if dead or trapped
	if err := checkAlien(alien); err != nil {
		return err
	}
	// Move
	from := alien.City
	to := s.pickConnectedCity(alien)
	if from == nil && to == nil {
		// At the beginning
		to = s.pickAnyCity()
		if to == nil {
			// no-op
			fmt.Printf(NoOpMessage, "World is destroyed.")
			return &NoOpError{reason: NoOpWorldDestroyed}
		}
	}
	fmt.Printf(" => To: %s\n", to)
	alien.City = to
	if from != nil {
		// Move from City
		delete(s.CityDefense[from.Name], alien.Name)
	}
	// Init city defense
	if (s.CityDefense[to.Name] == nil) {
		s.CityDefense[to.Name] = make(AlienOccupation)
	}
	// Move to City
	s.CityDefense[to.Name][alien.Name] = alien
	if len(s.CityDefense[to.Name]) > 1 {
		to.Destroy()
		// Kill Aliens and notify
		out := fmt.Sprintf(" || %s has been destroyed by ", to.Name)
		for _, a := range s.CityDefense[to.Name] {
			out += fmt.Sprintf("alien %s and ", a.Name)
			a.Kill();
		}
		out = out[:len(out) - 5] + "!\n"
		fmt.Print(out)
	}
	// Done
	return nil
}

// checkAlien returns NoOpError if Alien dead or trapped
func checkAlien(alien *Alien) *NoOpError {
	if alien.IsDead() {
		// no-op
		fmt.Printf(NoOpMessage, "Alien Dead.")
		return &NoOpError{NoOpAlienDisabled}
	}
	if alien.IsTrapped() {
		// no-op
		fmt.Printf(NoOpMessage, "Alien Trapped.")
		return &NoOpError{NoOpAlienDisabled}
	}
	return nil
}

// pickConnectedCity picks a random road to undestroyed City
func (s *Simulation) pickConnectedCity(alien *Alien) *City {
	// Nil if still not invading
	if !alien.IsInvading() {
		return nil
	}
	// Shuffle roads every pick
	picks := ShuffleLen(len(alien.City.Roads), s.R)
	// Any undestroyed connected city
	for _, p := range picks {
		roadKey := alien.City.Roads[p].Key
		c := alien.City.RoadsMap[roadKey]
		if (!c.IsDestroyed()) {
			return c
		}
	}
	// No connected undestroyed City
	return nil
}

// pickAnyCity picks any undestroyed City in the World
func (s *Simulation) pickAnyCity() *City {
	// Any undestroyed city, pick deterministically
	// TODO: optimize not to sort every pick
	var keys []string
	for k := range s.World {
		if c := s.World[k]; !c.IsDestroyed() {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	for _, k := range keys {
		c := s.World[k]
		if (!c.IsDestroyed()) {
			return c
		}
	}
	// All Cities destroyed
	return nil
}
