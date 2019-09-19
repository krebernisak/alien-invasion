package simulation

import (
	"errors"
)

const (
	// DefaultIterations is the default used if number of iteration is not othervise specified
	DefaultIterations int = 10000
)

// Simulation represents a running simulation
type Simulation struct {
	Entropy int64
	Iteration int
	EndIteration int
}

// Start the simulation
func Start() error {
	return errors.New("Not implemented (Yet!)")
}
