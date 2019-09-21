package types

import (
	"fmt"
	"sort"
	"strings"
)

// Road type representing connection between Cities
type Road struct {
	Key    string
	Cities []string
	Names  map[string]string
}

// NewRoad creates a Road with sorted key
func NewRoad(cities ...string) Road {
	sort.Strings(cities)
	key := strings.Join(cities, "_")
	return Road{key, cities, make(map[string]string)}
}

// PutName road name for a City
func (r Road) PutName(city, name string) {
	r.Names[city] = name
}

// String representation of a Road
func (r Road) String() string {
	return fmt.Sprintf("key=%s cities=%s\n", r.Key, r.Cities)
}
