package simulation

import (
	"testing"
	"math/rand"
)

func TestMakeRange(t *testing.T) {
	res := MakeRange(10, 20)
	if len(res) != 10 {
		t.Errorf("Range does not have correct len: %d != 10", len(res))
	}
	sum := Sum(res...)
	if sum != 145 {
		t.Errorf("Range does not have correct sum: %d != 10", sum)
	}
}

func TestShuffle(t *testing.T) {
	res := MakeRange(10, 20)
	sum := Sum(res...)

	r := rand.New(rand.NewSource(0xffffffff))
	Shuffle(res, r)
	if len(res) != 10 {
		t.Errorf("Shuffled array does not have correct len: %d != 10", len(res))
	}

	sSum := Sum(res...)
	if sum != sSum {
		t.Errorf("Shuffled array does not have correct sum: %d != %d", sum, sSum)
	}
}
