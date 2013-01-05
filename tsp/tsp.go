// Travelling Salesman
// given a list of points
// x y \n x y ... // eventually N-dim
// find an order in which 
// future iterations

// error codes:
// 1 - problem with I/O
// 2 - problem with command line options

package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"../utils"
)

////////////////////////////////////////////////////////////////////////////////
// ALGOS

// brute force: list all possible permutations of the points
func brute_force_tsp(points [][]float64) [][]float64 {
	// make a list of indicies to permute
	inds := make([]int, len(points))
	for i, _ := range inds {
		inds[i] = i
	}
	// run through each permutation
	min_inds := make([]int, len(points))
	var min_len float64 = math.Inf(1)
	for {
		// go through, add up the lengths
		var cur_len float64 = 0
		for i := 0; i < len(inds) - 1; i++ {
			cur_len += utils.Dist(points[inds[i]], points[inds[i+1]])
		}
		if cur_len < min_len {
			min_len = cur_len
			copy(min_inds, inds)
		}
		// permute
		inds := utils.PermuteIndicies(inds)
		// check if we're done
		if inds == nil {
			break
		}
	}
	// output the right order of points
	cur_points := make([][]float64, len(points))
	for i, v := range min_inds {
		cur_points[i] = points[v]
	}
	return cur_points
}

// use the nearest neighbor heuristic
func nearest_neighbor_tsp(points [][]float64) [][]float64 {
	var min_pts [][]float64
	var min_len float64 = math.Inf(1)
	// for each point as starting point, run the NN heuristic
	for i := 0; i < len(points); i++ {
		var cur_len float64 = 0
		cur_pts := make([][]float64, 1, len(points))
		set_pts := make([][]float64, len(points))
		// make a copy which we remove from, remove starting point
		copy(set_pts, points)
		cur_pt := set_pts[i]
		cur_pts[0] = cur_pt
		set_pts[i] = set_pts[len(set_pts) - 1]
		set_pts = set_pts[:len(set_pts) - 1]
		// find the shortest path
		for len(set_pts) > 0 {
			next_i := -1
			var next_pt []float64
			var next_dist float64 = math.Inf(1)
			// find next shortest point
			for i, p := range set_pts {
				d := utils.Dist(cur_pt, p)
				if d < next_dist {
					next_pt = p
					next_dist = d
					next_i = i
				}
			}
			cur_len += next_dist
			cur_pt = next_pt
			cur_pts = cur_pts[:len(cur_pts) + 1]
			cur_pts[len(cur_pts) - 1] = next_pt
			// remove the point
			set_pts[next_i] = set_pts[len(set_pts) - 1]
			set_pts = set_pts[:len(set_pts) - 1]
		}
		// check if it's smaller or not
		if cur_len < min_len {
			min_len = cur_len
			min_pts = cur_pts
		}
	}
	return min_pts
}

////////////////////////////////////////////////////////////////////////////////
// Runtime things

// algo types
var flag_algo_brute = flag.Bool("brute", false,
	`Brute force searcher: slow and obvious
	 Time: O(n!) Space: O(n)
		Guaranteed to be correct`)
var flag_algo_nn = flag.Bool("nearest-neighbor", false,
	`Nearest Neighbor heuristic: fast and approximate
	 Time: O(n**3) Space: O(n)
		Path produced is on average 25% longer than optimal:
	 [Johnson, D.S. and McGeoch, L.A.. "The traveling salesman problem:
	 A case study in local optimization", Local search in combinatorial
	 optimization, 1997, 215-310]`)

func main() {
	// parse the command line options
	flag.Parse()

	// get the list of points (maybe also a graph)
	points, err := utils.ReadPointsFromStdin()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if points == nil {
		fmt.Fprintln(os.Stderr, "No points input")
		os.Exit(1)
	}
	// solve it
	switch {
	case *flag_algo_brute:
		points = brute_force_tsp(points)
	case *flag_algo_nn:
		points = nearest_neighbor_tsp(points)
	default:
		fmt.Fprintln(os.Stderr, "No algorithm selected.")
		os.Exit(2)
	}
	// print it out
	fmt.Println(points)
}
