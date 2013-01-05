package algos

import (
	"math"
	"../../utils"
)

// brute force: list all possible permutations of the points
func BruteForce(points [][]float64) [][]float64 {
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
