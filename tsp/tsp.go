// Travelling Salesman
// given a list of points
// x y \n x y ... // eventually N-dim
// find an order in which 
// future iterations

package main

import (
	"flag"
	"fmt"
	"io"
	"math"
	"os"
)

////////////////////////////////////////////////////////////////////////////////
// UTILITIES

// this will probably eventually go in a utility file
func read_points_from_file(file io.Reader) ([][]float64, error) {
	var x, y float64
	points := make([][]float64, 0, 100)
	for i := 0;; i++ {
		// read pairs of floats for now
		n, err := fmt.Fscanf(file, "%f %f", &x, &y)
		if err == io.EOF {
			break
		}
		if err != nil || n != 2 {
			return nil, err
		}
		// check the capacity
		if len(points) == cap(points) {
			more_points := make([][]float64, len(points), 2*cap(points))
			copy(more_points, points)
			points = more_points
		}
		points = points[:i+1]
		points[i] = []float64{x, y}
	}
	return points, nil
}

func read_points_from_stdin() ([][]float64, error) {
	return read_points_from_file(os.Stdin)
}

////////////////////////////////////////////////////////////////////////////////
// ALGOS

// takes a pair of N-dim points, outputs the euclidean distance
func dist(x, y []float64) float64 {
	var sum float64 = 0
	for i, _ := range x {
		sum += math.Pow(x[i] - y[i], 2)
	}
	return math.Sqrt(sum)
}

// from http://en.wikipedia.org/wiki/Permutation#Generation_in_lexicographic_order
func permute_indicies(inds []int) []int {
	var k, l int
	// find the largest k a[k] < a[k+1]; none means done
	for k = len(inds) - 2; k >= 0; k-- {
		if inds[k] < inds[k + 1] {
			break
		}
	}
	if k == -1 {
		return nil
	}
	// find largest l a[k] < a[l]
	for l = len(inds) - 1; k < l ; l-- {
		if inds[k] < inds[l] {
			break
		}
	}
	// swap k/l
	tmp := inds[k]
	inds[k] = inds[l]
	inds[l] = tmp
	// reverse a[k+1] until end
	for j := 0; j < (len(inds) - k + 1)/2; j++ {
		tmp = inds[k+1 + j]
		inds[k+1 + j] = inds[len(inds) - 1 - j]
		inds[len(inds) - 1 - j] = tmp
	}
	return inds
}

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
			cur_len += dist(points[inds[i]], points[inds[i+1]])
		}
		if cur_len < min_len {
			min_len = cur_len
			copy(min_inds, inds)
		}
		// permute
		inds := permute_indicies(inds)
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

////////////////////////////////////////////////////////////////////////////////
// Runtime things

// algo types
var brute_flag = flag.Bool("brute", false, `Brute force searcher
	Time: O(n!) Space: O(n)`)

func main() {
	// parse the command line options
	flag.Parse()
	// get the list of points (maybe also a graph)
	points, err := read_points_from_stdin()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if points == nil {
		fmt.Println("No points input")
		os.Exit(2)
	}
	// solve it
	switch {
	case *brute_flag:
		points = brute_force_tsp(points)
	default:
		fmt.Println("No algorithm selected.")
		os.Exit(3)
	}
	// print it out
	fmt.Println(points)
}
