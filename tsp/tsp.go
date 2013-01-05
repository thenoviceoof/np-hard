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
	"os"
	"../utils"
	"./algos"
)

// algo types
var flag_algo_brute = flag.Bool("brute", false,
	`Brute force searcher: slow and obvious
	 Time: O(n!) Space: O(n)
		Multithreaded by default, turn off with -single-core`)
var flag_algo_nn = flag.Bool("nearest-neighbor", false,
	`Nearest Neighbor heuristic: fast and approximate
	 Time: O(n**3) Space: O(n)
		Path produced is on average 25% longer than optimal:
	 [Johnson, D.S. and McGeoch, L.A.. "The traveling salesman problem:
	 A case study in local optimization", Local search in combinatorial
	 optimization, 1997, 215-310]`)

// option flags
var flag_single_core = flag.Bool("single-core", false,
	"Disable multithreading")

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
	case *flag_algo_brute && *flag_single_core:
		points = algos.BruteForce(points)
	case *flag_algo_brute:
		points = algos.BruteForceMT(points)
	case *flag_algo_nn:
		points = algos.NearestNeighbor(points)
	default:
		fmt.Fprintln(os.Stderr, "No algorithm selected.")
		os.Exit(2)
	}
	// print it out
	fmt.Println(points)
}
