package cli

import (
	"fmt"
	"os"
	"flag"
	"math/rand"
	"crypto/sha256"
	"encoding/binary"
	"time"

	simulator "alien-invasion/simulation"
)

var (
	entropy int64
	iterations, aliens int
	simulation, world, intel, log string
)

func init() {
	flag.Int64Var(&entropy, "entropy", 0, "random number used as entropy seed")
	flag.IntVar(&iterations, "iterations", simulator.DefaultIterations, "number of iterations")
	flag.IntVar(&aliens, "aliens", 0, "number of aliens invading")
	flag.StringVar(&simulation, "simulation", "", "name hashed and used as entropy seed")
	flag.StringVar(&world, "world", "", "file used as world map input")
	flag.StringVar(&intel, "intel", "", "file used to identify aliens")
	flag.StringVar(&log, "log", "debug", "log level used in debugging")
	flag.Parse()
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	var source rand.Source
	if entropy != 0 {
		source = rand.NewSource(entropy)
		fmt.Printf("Entropy: using --entropy flag\n")
		fmt.Printf("Entropy: %v\n", entropy)
	} else if len(simulation) > 0 {
		hash := sha256.Sum256([]byte(simulation))
		seed := int64(binary.BigEndian.Uint64(hash[:8]))
		source = rand.NewSource(seed)
		fmt.Printf("Entropy: using first 8 bytes of sha256(\"%v\")\n", simulation)
		fmt.Printf("Entropy: %v\n", seed)
	} else {
		now := time.Now().UnixNano();
		source = rand.NewSource(now)
		fmt.Printf("Entropy: using current unix time as a random source\n")
		fmt.Printf("Entropy: %v\n", now)
	}

	r := rand.New(source)
	fmt.Println(r.Int())

	if err := simulator.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
