package simulation

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"math/rand"

	"alien-invasion/types"
)

// InputCityList is a list of cities as read from the input file lines
type InputCityList []*City

const (
	// RandAlienNameLen is a constant used to normalize number choosen as Alien name
	RandAlienNameLen = 10
)

// String representation of a InputCityList, used to display output in same format
func (in InputCityList) String() string {
	var out string
	for _, city := range in {
		if city.IsDestroyed() {
			continue
		}
		out += fmt.Sprintf("%s", city.Name)
		for _, r := range city.Roads {
			c := city.RoadsMap[r.Key]
			if (c.IsDestroyed()) {
				continue
			}
			// TODO: Avoid double Roads ?roadName != ""
			if roadName := r.Names[c.Name]; true {
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

// RandAliens creates N new Alien objects with random names
func RandAliens(n int, r *rand.Rand) []*Alien {
	out := make([]*Alien, 0)
	for n > 0 {
		name := strconv.Itoa(r.Int())
		alien := types.NewAlien(name[:RandAlienNameLen])
		out = append(out, &alien)
		n--
	}
	return out
}

// ReadWorldMapFile takes in a file and constructs a World map
func ReadWorldMapFile(file string) (World, InputCityList, error) {
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
	input := make(InputCityList, 0)
	for scanner.Scan() {
		sections := strings.Split(scanner.Text(), " ")
		// Add new City to the world map
		city := w.AddNewCity(sections[0])
		// Connect nearby Cities
		connections := sections[1:]
		for _, c := range connections {
			link := strings.Split(c, "=")
			roadName, cityName := link[0], link[1]
			linkedCity, exists := w[cityName];
			if !exists {
				linkedCity = w.AddNewCity(cityName)
			}
			// Discovered a Road
			road := types.NewRoad(city.Name, cityName)
			road.PutName(cityName, roadName)
			// Link Cities
			city.AddRoad(&road, linkedCity);
			linkedCity.AddRoad(&road, city);
		}
		input = append(input, city)
		fmt.Printf("Reading... %s\n", city)
	}

	return w, input, nil
}
