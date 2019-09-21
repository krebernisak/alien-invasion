package cli

import (
	"crypto/sha256"
	"encoding/binary"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"alien-invasion/simulation"
)

const (
	// DefaultIterations used if number of iteration is not othervise specified
	DefaultIterations int = 10000
	// DefaultNumberOfAliens used if number of Aliens is not othervise specified
	DefaultNumberOfAliens int = 10
	// DefaultWorldFile used if World file is not othervise specified
	DefaultWorldFile = "./test/example.txt"
)

var (
	entropy                          int64
	iterations, alienNumber          int
	simulationName, worldFile, intel string
)

func init() {
	flag.Int64Var(&entropy, "entropy", 0, "random number used as entropy seed")
	flag.IntVar(&iterations, "iterations", DefaultIterations, "number of iterations")
	flag.IntVar(&alienNumber, "aliens", DefaultNumberOfAliens, "number of aliens invading")
	flag.StringVar(&simulationName, "simulation", "", "name hashed and used as entropy seed")
	flag.StringVar(&worldFile, "world", DefaultWorldFile, "file used as world map input")
	flag.StringVar(&intel, "intel", "", "file used to identify aliens")
	flag.Parse()
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	var source rand.Source
	if entropy != 0 {
		source = rand.NewSource(entropy)
		fmt.Printf("Entropy: using --entropy flag\n")
		fmt.Printf("Entropy Seed: %v\n", entropy)
	} else if len(simulationName) > 0 {
		hash := sha256.Sum256([]byte(simulationName))
		seed := int64(binary.BigEndian.Uint64(hash[:8]))
		source = rand.NewSource(seed)
		fmt.Printf("Entropy: using first 8 bytes of sha256(\"%v\")\n", simulationName)
		fmt.Printf("Entropy Seed: %v\n", seed)
	} else {
		now := time.Now().UnixNano()
		source = rand.NewSource(now)
		fmt.Printf("Entropy: using current unix time as a random source\n")
		fmt.Printf("Entropy Seed: %v\n", now)
	}
	// Read input file
	fmt.Printf("Reading world map from file: %s\n", worldFile)
	world, list, err := simulation.ReadWorldMapFile(worldFile)
	if err != nil {
		fmt.Printf("Could not read world from map file \"%s\" with error: %s\n", world, err)
		os.Exit(1)
	}
	// Build simulator
	r := rand.New(source)
	fmt.Printf("Generating %d random Aliens...\n", alienNumber)
	aliens := simulation.RandAliens(alienNumber, r)
	sim := simulation.NewSimulation(r, iterations, world, aliens)
	// Start simulation and print any errors
	if err := sim.Start(); err != nil {
		fmt.Printf(formatImportantMessage("Error while running simulation: %s"), err)
		os.Exit(1)
	}
	// Success
	fmt.Printf(formatImportantMessage("Simulation Success"))
	fmt.Print(list)
}

func formatImportantMessage(msg string) string {
	out := fmt.Sprintf("\n\n")
	out += fmt.Sprintf("==================\n")
	out += fmt.Sprintf("%s\n", msg)
	out += fmt.Sprintf("==================\n")
	return out
}
