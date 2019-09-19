package types

import (
	"fmt"
)

const (
	// FlagDead is a flag used to mark dead Aliens
	FlagDead string = "dead"
)

// Alien can be dead or alive and occupating a City
type Alien struct {
	Name string
	Flags map[string]bool
	City *City
}

// NewAlien creates an Alien with a name and default flags
func NewAlien(name string) Alien {
	flags := map[string]bool{FlagDead: false}
	return Alien{Name: name, Flags: flags}
}

// IsDead checks if Alien died
func (a *Alien) IsDead() bool {
	return a.Flags[FlagDead];
}

// Kill Alien makes it dead
func (a *Alien) Kill() {
	a.Flags[FlagDead] = true
}

// IsTrapped checks if Alien is trapped in a City with no roads out
func (a *Alien) IsTrapped() bool {
	var roads int
	for _, c := range a.City.Roads {
		if (!c.IsDestroyed()) {
			roads++
		}
	}
	return roads <= 0;
}

// String representation for an Alien
func (a *Alien) String() string {
	return fmt.Sprintf("Alien: %s city=%s flags=%v\n", a.Name, a.City, a.Flags)
}
