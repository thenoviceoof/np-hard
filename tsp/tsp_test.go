// currently, test the utils and TSP brute forcer

package tsp

import (
	"fmt"
	"testing"
)

// func Test_read_points_from_stdin() {
// 	p1s := [][]float64{{1, 2}, {2, 3}}
// 	// TODO
// }

func Test_brute_force_tsp(t *testing.T) {
	pts := [][]float64{
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
	pts = brute_force_tsp(pts)
	for i, v := range pts {
		if fmt.Sprintf("%v", v) != fmt.Sprintf("%v", result[i]) {
			t.Errorf("Got %v instead of %v", pts, result)
			break
		}	
	}
}
