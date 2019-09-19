package types

import (
	"fmt"
)

// World is a map of Cities
type World map[string]*City

// AddCity to a World by name
func (w World) AddCity(name string) *City {
	city := NewCity(name)
	w[city.Name] = &city
	return &city
}

// String representation of a World
func (w World) String() string {
	var out string
	for _, city := range w {
		out += fmt.Sprintf("%s\n", city)
	}
	return out
}

// AlienOccupation maps all Aliens by name
type AlienOccupation map[string]*Alien

// CityDefence maps Aliens by City
type CityDefence map[string][]*Alien
