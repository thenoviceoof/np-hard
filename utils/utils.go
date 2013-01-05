/* utils.go
 * various utilities used by different np-hard citizens
 */

package utils

import (
	"fmt"
	"io"
	"math"
	"os"
)

//------------------------------------------------------------------------------
// IO utils

func ReadPointsFromFile(file io.Reader) ([][]float64, error) {
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

func ReadPointsFromStdin() ([][]float64, error) {
	return ReadPointsFromFile(os.Stdin)
}

//------------------------------------------------------------------------------
// graph utils

// takes a pair of N-dim points, outputs the euclidean distance
func Dist(x, y []float64) float64 {
	var sum float64 = 0
	for i, _ := range x {
		sum += math.Pow(x[i] - y[i], 2)
	}
	return math.Sqrt(sum)
}

// NOT USED CURRENTLY
// go through a path, add up all the distances between the points sequentially
func PathLength(points [][]float64) float64 {
	var cur_len float64 = 0
	for i := 0; i < len(points) - 1; i++ {
		cur_len += Dist(points[i], points[i+1])
	}
	return cur_len
}

// http://en.wikipedia.org/wiki/Permutation#Generation_in_lexicographic_order
func PermuteIndicies(inds []int) []int {
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