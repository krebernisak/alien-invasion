package simulation

import (
	"fmt"
	"testing"
	"math/rand"
)

func TestWorldReadFromFile(t *testing.T) {
	w, input, err := ReadWorldMapFile("../test/example.txt")
	if err != nil {
		t.Errorf("%v: could not read file", err)
	}
	fmt.Printf("Input:\n%s\n\n", input)
	fmt.Printf("World:\n%s\n\n", w)
}

func testRandAliens(t *testing.T, n int, seed int64) {
	source := rand.NewSource(seed)
	r := rand.New(source)
	aliens := RandAliens(n, r)
	if (len(aliens) != n) {
		t.Errorf("len(RandAliens(%d, 0xffffffff)) = %d; want %d", n, len(aliens), n)
		return
	}
	fmt.Printf("Aliens:\n%s", aliens)
}

func TestRandAliens(t *testing.T) {
	testRandAliens(t, 10, 0xffffffff);
}

func TestRandAliensDouble(t *testing.T) {
	testRandAliens(t, 20, 0xffffffff);
}