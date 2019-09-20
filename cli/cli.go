package cli

import (
	"fmt"
	"os"
	"flag"
	"math/rand"
	"crypto/sha256"
	"encoding/binary"
	"time"

	"alien-invasion/simulation"
)

var (
	entropy int64
	iterations, alienNumber int
	simulationName, worldFile, intel string
)

func init() {
	flag.Int64Var(&entropy, "entropy", 0, "random number used as entropy seed")
	flag.IntVar(&iterations, "iterations", simulation.DefaultIterations, "number of iterations")
	flag.IntVar(&alienNumber, "aliens", 0, "number of aliens invading")
	flag.StringVar(&simulationName, "simulation", "", "name hashed and used as entropy seed")
	flag.StringVar(&worldFile, "world", "", "file used as world map input")
	flag.StringVar(&intel, "intel", "", "file used to identify aliens")
	flag.Parse()
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	var source rand.Source
	if entropy != 0 {
		source = rand.NewSource(entropy)
		fmt.Printf("Entropy: using --entropy flag\n")
		fmt.Printf("Entropy: %v\n", entropy)
	} else if len(simulationName) > 0 {
		hash := sha256.Sum256([]byte(simulationName))
		seed := int64(binary.BigEndian.Uint64(hash[:8]))
		source = rand.NewSource(seed)
		fmt.Printf("Entropy: using first 8 bytes of sha256(\"%v\")\n", simulationName)
		fmt.Printf("Entropy: %v\n", seed)
	} else {
		now := time.Now().UnixNano();
		source = rand.NewSource(now)
		fmt.Printf("Entropy: using current unix time as a random source\n")
		fmt.Printf("Entropy: %v\n", now)
	}

	world, _, err := simulation.ReadWorldMapFile(worldFile)
	if err != nil {
		fmt.Printf("Could not read world from map file \"%s\" with error: %s\n", world, err)
		os.Exit(1)
	}

	r := rand.New(source)
	aliens := simulation.RandAliens(alienNumber, r)
	sim := simulation.NewSimulation(r, iterations, world, aliens);

	if err := sim.Start(); err != nil {
		fmt.Printf("Error while running simulation: %s\n", err)
		os.Exit(1)
	}
}
