package types

import (
	"fmt"
)

// World is a map of Cities
type World map[string]*City

// AddCity to a World
func (w World) AddCity(city City) *City {
	w[city.Name] = &city
	return &city
}

// AddNewCity to a World with name
func (w World) AddNewCity(name string) *City {
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