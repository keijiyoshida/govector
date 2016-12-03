package vector

import "errors"

// Errors
var (
	ErrDifferentLength = errors.New("lengths of the input vectors are different")
)

// DotProduct returns the dot product of the two input vectors.
func DotProduct(x, y []float64) (float64, error) {
	if len(x) != len(y) {
		return 0.0, ErrDifferentLength
	}

	ch := make(chan float64, numConcurrency)

	n := len(x)
	d := divUp(n, numConcurrency)

	for i := 0; i < numConcurrency; i++ {
		from := d * i
		to := min(d*(i+1), n)

		go dotProduct(x, y, from, to, ch)
	}

	sum := 0.0

	for i := 0; i < numConcurrency; i++ {
		sum += <-ch
	}

	return sum, nil
}

// dotProduct returns the dot product of the two input vectors.
func dotProduct(x, y []float64, from, to int, ch chan<- float64) {
	sum := 0.0

	for i := from; i < to; i++ {
		sum += x[i] * y[i]
	}

	ch <- sum
}
