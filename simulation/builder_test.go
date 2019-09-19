package simulation

import (
	"fmt"
	"testing"
	"math/rand"
)

func TestWorldReadFromFile(t *testing.T) {
	w := ReadWorldMapFile("../test/example.txt")
	fmt.Printf("World:\n%s", w)
}

func testRandAliens(t *testing.T, n int, seed int64) {
	source := rand.NewSource(seed)
	aliens := RandAliens(n, source)
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