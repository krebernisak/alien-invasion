package types

import (
	"alien-invasion/types"
)

const (
	// FlagDestroyed is a flag used to mark destroyed Cities
	FlagDestroyed string = "destroyed"
)

// City has a name and is connected to other Cities via roads
type City struct {
	types.Node
	RoadNames map[string]string
}

// NewCity creates a City with a name and default flags
func NewCity(name string) City {
	// FlagDestroyed default is false
	return City{types.NewNode(name), make(map[string]string)}
}

// IsDestroyed checks if City is destroyed
func (c *City) IsDestroyed() bool {
	return c.Flags[FlagDestroyed]
}

// Destroy City makes City burn in flames
func (c *City) Destroy() {
	c.Flags[FlagDestroyed] = true
}
