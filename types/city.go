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
	Name     string
	Flags    map[string]bool
	Roads    []*Road
	RoadsMap map[string]*City
}

// NewCity creates a City with a name and default flags
func NewCity(name string) City {
	return City{
		Name:     name,
		Flags:    map[string]bool{FlagDestroyed: false},
		Roads:    make([]*Road, 0),
		RoadsMap: make(map[string]*City),
	}
}

// AddRoad to linked City
func (c *City) AddRoad(road *Road, link *City) {
	// TODO: merge roads if road with same key exists
	c.Roads = append(c.Roads, road)
	c.RoadsMap[road.Key] = link
}

// IsDestroyed checks if City is destroyed
func (c *City) IsDestroyed() bool {
	return c.Flags[FlagDestroyed]
}

// Destroy City makes City burn in flames
func (c *City) Destroy() {
	c.Flags[FlagDestroyed] = true
}

// String representation for a City
func (c *City) String() string {
	out := fmt.Sprintf("name=%s roads=map[", c.Name)
	for k, c := range c.RoadsMap {
		out += fmt.Sprintf("%s:%s ", k, c.Name)
	}
	if len(c.RoadsMap) > 0 {
		return out[:len(out)-1] + "]"
	}
	return out
}
