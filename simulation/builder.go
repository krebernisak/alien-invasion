package simulation

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"alien-invasion/simulation/types"
)

// WorldMapFile is a list of cities as read from the input file
type WorldMapFile []*City

const (
	// RandAlienNameLen is a constant used to normalize number chosen as Alien name
	RandAlienNameLen = 10
)

// String representation of a WorldMapFile, used to display output in same format
func (in WorldMapFile) String() string {
	var out string
	for _, city := range in {
		if city.IsDestroyed() {
			continue
		}
		out += fmt.Sprintf("%s", city.Name)
		for _, r := range city.Links {
			n := city.Nodes[r.Key]
			c := City{Node: *n}
			if c.IsDestroyed() {
				continue
			}
			// TODO: Avoid double Roads ?roadName != ""
			if roadName := city.RoadNames[r.Key]; true {
				if roadName == "" {
					roadName = "?"
				}
				out += fmt.Sprintf(" %s=%s", roadName, c.Name)
			}
		}
		out += fmt.Sprintln()
	}
	return out
}

// RandAliens creates N new Aliens with random names
func RandAliens(n int, r *rand.Rand) []*Alien {
	out := make([]*Alien, 0)
	for n > 0 {
		name := strconv.Itoa(r.Int())[:RandAlienNameLen]
		alien := types.NewAlien(name)
		out = append(out, &alien)
		n--
	}
	return out
}

// ReadWorldMapFile takes in a file and constructs a World map
func ReadWorldMapFile(file string) (World, WorldMapFile, error) {
	// Open and close file
	f, err := os.Open(file)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	// Init scanner to scan lines
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	// Prepare data structures
	w := make(World)
	input := make(WorldMapFile, 0)
	for scanner.Scan() {
		sections := strings.Split(scanner.Text(), " ")
		// Add new City to the world map
		city := w.AddNewCity(sections[0])
		// Connect nearby Cities
		connections := sections[1:]
		for _, c := range connections {
			link := strings.Split(c, "=")
			roadName, cityName := link[0], link[1]
			other, exists := w[cityName]
			if !exists {
				other = w.AddNewCity(cityName)
			}
			// Discovered a Road => Link Cities
			road := city.Connect(&other.Node)
			city.RoadNames[road.Key] = roadName
			other.ConnectVia(road, &city.Node)
		}
		input = append(input, city)
		fmt.Printf("Reading... %s\n", city)
	}

	return w, input, nil
}
