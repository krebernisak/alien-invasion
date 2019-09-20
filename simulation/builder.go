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

type (
	// Alien attacking City
	Alien = types.Alien
	// City connected to other Cities with Roads
	City = types.City
	// World map of Cities
	World = types.World
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
		out += fmt.Sprintf("%s", city.Name)
		for k, c := range city.Roads {
			out += fmt.Sprintf(" %s=%s", k, c.Name)
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
	w := make(World)
	f, err := os.Open(file)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	input := make(InputCityList, 0)
	for scanner.Scan() {
		sections := strings.Split(scanner.Text(), " ")
		// create and add city to the world map
		city := w.AddNewCity(sections[0])
		// create and link connections
		connections := sections[1:]
		for _, c := range connections {
			link := strings.Split(c, "=")
			roadName, cityName := link[0], link[1]
			linkedCity, exists := w[cityName];
			if !exists {
				linkedCity = w.AddNewCity(cityName)
			}
			city.Roads[roadName] = linkedCity
			if _, ok := linkedCity.Roads[roadName]; !ok {
				linkedCity.Roads[roadName] = city
			}
		}
		input = append(input, city)
		fmt.Printf("Reading... %s\n", city)
	}

	return w, input, nil
}