package algos

import (
	"math"
	"runtime"
	"../../utils"
)

func PathLengthInds(points [][]float64, inds []int) float64 {
	var cur_len float64 = 0
	for i := 0; i < len(inds) - 1; i++ {
		cur_len += utils.Dist(points[inds[i]], points[inds[i+1]])
	}
	return cur_len
}

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
		cur_len := PathLengthInds(points, inds)
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

//------------------------------------------------------------------------------
// multithreaded version

// permutation generator
func permutation_generator(length int, ind_chan chan []int) {
	// make a list of indicies to permute
	inds := make([]int, length)
	for i, _ := range inds {
		inds[i] = i
	}
	// put all the permutations in a channel
	for inds != nil {
		// copy out the indicies slice to prevent mem access shenanigans
		tmp_inds := make([]int, len(inds))
		copy(tmp_inds, inds)
		ind_chan <- tmp_inds
		// permute
		inds = utils.PermuteIndicies(inds)
	}
	close(ind_chan)
}

// pulls indicies permutations, finds the longest
func brute_force_worker(points [][]float64, ind_chan chan []int, results chan []int) {
	min_inds := make([]int, len(points))
	var min_len float64 = math.Inf(1)
	// keep pulling until it's over
	for {
		inds, ok := <-ind_chan
		if !ok {
			break
		}
		cur_len := PathLengthInds(points, inds)
		// find the winner
		if cur_len < min_len {
			min_len = cur_len
			copy(min_inds, inds)
		}
	}

	results <- min_inds
}

func BruteForceMT(points [][]float64) [][]float64 {
	ind_chan := make(chan []int)
	result_ind_chan := make(chan []int)
	go permutation_generator(len(points), ind_chan)
	// start up the workers
	for i := 0; i < runtime.NumCPU(); i++ {
		go brute_force_worker(points, ind_chan, result_ind_chan)
	}
	// fetch the results
	min_inds := <-result_ind_chan
	for i := 0; i < runtime.NumCPU() - 1; i++ {
		cur_inds := <-result_ind_chan
		if PathLengthInds(points, cur_inds) < PathLengthInds(points, min_inds) {
			min_inds = cur_inds
		}
	}
	// output the right order of points
	cur_points := make([][]float64, len(points))
	for i, v := range min_inds {
		cur_points[i] = points[v]
	}
	return cur_points
}
