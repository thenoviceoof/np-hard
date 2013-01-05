// currently, test the utils and TSP brute forcer

package main

import (
	"./algos"
	"testing"
	"../utils"
)

func run_TSP_fn(name string, t *testing.T, points [][]float64,
	f func([][]float64) [][]float64, optimal_length float64) {
	pts := f(points)
	if utils.PathLength(pts) != optimal_length {
		if len(pts) < 5 {
			t.Errorf("%s: Got %v with path length %d/%d", name, pts,
				utils.PathLength(pts),
				optimal_length)
			return
		}
		t.Errorf("%s: Got length %d/%d", name,
			utils.PathLength(pts),
			optimal_length)
	}
}

func run_approx_TSP_fn(name string, t *testing.T, points [][]float64,
	f func([][]float64) [][]float64, optimal_length float64) {
	pts := f(points)
	if utils.PathLength(pts) < optimal_length {
		if len(pts) < 5 {
			t.Errorf("%s: Got %v with path length %d < %d", name, pts,
				utils.PathLength(pts),
				optimal_length)
			return
		}
		t.Errorf("%s: Got length %d < %d", name,
			utils.PathLength(pts),
			optimal_length)
	}
}

func TestSimplePath(t *testing.T) {
	var path_length float64 = 3
	points := [][]float64{
		{0, 0},
		{3, 0},
		{1, 0},
		{2, 0},
	}

	run_TSP_fn("Brute", t, points, algos.BruteForce, path_length)
	run_TSP_fn("Brute SMP", t, points, algos.BruteForceMT, path_length)
	run_approx_TSP_fn("NN", t, points, algos.NearestNeighbor, path_length)
}

func TestRandomPath(t *testing.T) {
	var path_length float64 = 2.092894936771131
	points := [][]float64{
		{0.4273942341481516, 0.9379795141987419},
		{0.2749888745447667, 0.5756271528388281},
		{0.3800397429906439, 0.6144971508938101},
		{0.5193031806826659, 0.9721129741905882},
		{0.38923549084062004, 0.8525876796184378},
		{0.008802883350379598, 0.8943844427430961},
		{0.24380872115156227, 0.32334453201318303},
		{0.812782456554204, 0.43835018371386203},
		{0.4762686345868331, 0.8538090652413308},
		{0.636409196154668, 0.8744084651387656},
	}

	run_TSP_fn("Brute", t, points, algos.BruteForce, path_length)
	run_TSP_fn("Brute SMP", t, points, algos.BruteForceMT, path_length)
	run_approx_TSP_fn("NN", t, points, algos.NearestNeighbor, path_length)
}
