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

const (
	// RandAlienNameLen is a constant used to normalize number choosen as Alien name
	RandAlienNameLen = 10
)

// RandAliens creates N new Alien objects with random names
func RandAliens(n int, source rand.Source) []*Alien {
	out := make([]*Alien, 0)
	r := rand.New(source)
	for n > 0 {
		name := strconv.Itoa(r.Int())
		alien := types.NewAlien(name[:RandAlienNameLen])
		out = append(out, &alien)
		n--
	}
	return out
}

// ReadWorldMapFile takes in a file and constructs a World map
func ReadWorldMapFile(file string) World {
	w := make(World)
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		sections := strings.Split(scanner.Text(), " ")
		// create and add city to the world map
		city := w.AddCity(sections[0])
		// create and link connections
		connections := sections[1:]
		for _, c := range connections {
			link := strings.Split(c, "=")
			roadName, cityName := link[0], link[1]
			linkedCity, exists := w[cityName];
			if !exists {
				linkedCity = w.AddCity(cityName)
			}
			city.Roads[roadName] = linkedCity
			if _, ok := linkedCity.Roads[roadName]; !ok {
				linkedCity.Roads[roadName] = city
			}
		}
		fmt.Printf("Reading... %s\n", city)
	}

	return w
}
