package algos

import (
	"math"
	"../../utils"
)

// use the nearest neighbor heuristic
func NearestNeighborTSP(points [][]float64) [][]float64 {
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
