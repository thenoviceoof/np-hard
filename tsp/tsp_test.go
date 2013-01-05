// currently, test the utils and TSP brute forcer

package main

import (
	"fmt"
	"testing"
	"./algos"
)

// func Test_read_points_from_stdin() {
// 	p1s := [][]float64{{1, 2}, {2, 3}}
// 	// TODO
// }

func TestSimplePath(t *testing.T) {
	points := [][]float64{
		{0, 0},
		{3, 0},
		{1, 0},
		{2, 0},
	}
	result := [][]float64{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
	}
	pts := algos.BruteForceTSP(points)
	for i, v := range pts {
		if fmt.Sprintf("%v", v) != fmt.Sprintf("%v", result[i]) {
			t.Errorf("Got %v instead of %v", pts, result)
			break
		}
	}

	pts = algos.NearestNeighborTSP(points)
	for i, v := range pts {
		if fmt.Sprintf("%v", v) != fmt.Sprintf("%v", result[i]) {
			t.Errorf("Got %v instead of %v", pts, result)
			break
		}
	}
}
