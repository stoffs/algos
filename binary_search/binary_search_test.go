package main

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {

	e := 100000
	seed := make([]int, 0, e)

	for ctr := 0; ctr <= e; ctr++ {
		seed = append(seed, ctr)
	}

	n := 100

	shouldPosition := n
	realPosition, err := bs(seed, n, 0, len(seed))

	if err != nil {
		fmt.Printf("Number not found: %s\n", err)
	}

	if shouldPosition != realPosition {
		t.Error(
			"expected:", shouldPosition,
			"got:", realPosition,
		)
	}
}

func BenchmarkBinarySearch(b *testing.B) {

	e := 1000
	seed := make([]int, 0, e)

	for ctr := 0; ctr <= e; ctr++ {
		seed = append(seed, ctr)
	}

	n := 100

	shouldPosition := n
	realPosition, _ := bs(seed, n, 0, len(seed))

	if shouldPosition != realPosition {
		b.Error(
			"expected:", shouldPosition,
			"got:", realPosition,
		)
	}
}

func TestAvg(t *testing.T) {
	seed := map[int][]int{
		2:  []int{2, 2},
		4:  []int{4, 5},
		7:  []int{10, 5},
		50: []int{90, 10},
	}

	for k, v := range seed {

		r := avg(v[0], v[1])
		expected := k

		if r != expected {
			t.Error(
				"expected:", expected,
				"got:", r,
			)
		}

	}
}
