// currently, test the utils and TSP brute forcer

package main

import (
	"./algos"
	"testing"
	"../utils"
)

func TestSimplePath(t *testing.T) {
	var path_length float64 = 3
	points := [][]float64{
		{0, 0},
		{3, 0},
		{1, 0},
		{2, 0},
	}

	pts := algos.BruteForceTSP(points)
	if utils.PathLength(pts) != path_length {
		t.Errorf("Brute: Got %v with path length %d/%d", pts,
			utils.PathLength(pts),
			path_length)
	}

	pts = algos.NearestNeighborTSP(points)
	if utils.PathLength(pts) != path_length {
		t.Errorf("NN: Got %v with path length %d/%d", pts,
			utils.PathLength(pts),
			path_length)
	}
}
