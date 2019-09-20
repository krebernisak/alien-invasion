package types

import (
	"fmt"
)

const (
	// FlagDestroyed is a flag used to mark destroyed Cities
	FlagDestroyed string = "destroyed"
)

// City has a name and is connected to other Cities via roads
type City struct {
	Name string
	Flags map[string]bool
	Roads map[string]*City
}

// NewCity creates a City with a name and default flags
func NewCity(name string) City {
	flags := map[string]bool{FlagDestroyed: false}
	roads := make(map[string]*City)
	return City{Name: name, Flags: flags, Roads: roads}
}

// IsDestroyed checks if City is destroyed
func (c *City) IsDestroyed() bool {
	return c.Flags[FlagDestroyed];
}

// Destroy City makes City burn in flames
func (c *City) Destroy() {
	c.Flags[FlagDestroyed] = true
}

// String representation for a City
func (c *City) String() string {
	out := fmt.Sprintf("City: %s", c.Name)
	for road, city := range c.Roads {
		out += fmt.Sprintf("\n %s => %s", road, city.Name)
	}
	out += "\n"
	return out
}