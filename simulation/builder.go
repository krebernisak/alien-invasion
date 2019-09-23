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

const (
	// RandAlienNameLen is a constant used to normalize number chosen as Alien name
	RandAlienNameLen = 10
)

//  Four corners of the World (cardinal directions)
var compass = map[string]string{
	"north": "south",
	"south": "north",
	"west":  "east",
	"east":  "west",
}

// WorldMapFile is a list of cities as read from the input file
type WorldMapFile []*City

// FilterDestroyed Cities from WorldMapFile
func (in WorldMapFile) FilterDestroyed(world types.World) WorldMapFile {
	out := make(WorldMapFile, 0)
	processed := make(map[string]bool)
	for _, city := range in {
		// If processed continue
		if processed[city.Name] {
			continue
		}
		// If not destroyed process
		if !city.IsDestroyed() {
			out = append(out, city)
			processed[city.Name] = true
			continue
		}
		// If destroyed process links
		for _, link := range city.Links {
			n := city.Nodes[link.Key]
			other := world[n.Name]
			if processed[other.Name] || other.IsDestroyed() {
				continue
			}
			out = append(out, other)
			processed[other.Name] = true
		}
	}
	return out
}

// String representation of a WorldMapFile, used to display output in same format as input
func (in WorldMapFile) String() string {
	var out string
	for _, city := range in {
		// If destroyed print nothing
		if city.IsDestroyed() {
			continue
		}
		// If survived print city name with links
		out += fmt.Sprintf("%s\n", city)
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

// IdentifyAliens reads Alien intel file and names Aliens
func IdentifyAliens(aliens []*Alien, file string) error {
	// Open and close file
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	// Init scanner to scan lines
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	// Give names to aliens until names available
	for i := 0; i < len(aliens) && scanner.Scan(); i++ {
		aliens[i].Name = scanner.Text()
	}
	return nil
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
		for _, c := range sections[1:] {
			linkStr := strings.Split(c, "=")
			roadName, cityName := linkStr[0], linkStr[1]
			other, exists := w[cityName]
			if !exists {
				other = w.AddNewCity(cityName)
			}
			// Discovered a Road => Link Cities
			link := city.Connect(&other.Node)
			other.ConnectVia(link, &city.Node)
			city.RoadNames[link.Key] = roadName
			other.RoadNames[link.Key] = compass[roadName]
		}
		input = append(input, city)
		fmt.Printf("Reading... %s\n", city)
	}

	return w, input, nil
}
