package metric

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

// Metric computes distance between two given points.
type Metric interface {
	Dist(x, y float64) float64
}

// MetricFunc
type MetricFunc func(x, y float64) float64

func (f MetricFunc) Dist(x, y float64) float64 {
	return f(x, y)
}

var (
	Euclidean = MetricFunc(func(x, y float64) float64 { return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2)) })
)

// A DistanceMatrix stores the distances between every point in a set.
//
// Note: distance matricies are symmetric so M = transpose(M).
//
type DistanceMatrix struct {
	dists []float64
	vals  []float64
}

// Compute
func Compute(stream <-chan float64, d Metric) *DistanceMatrix {
	m := new(DistanceMatrix)
	for x := range stream {
		m.vals = append(m.vals, x)
		for _, y := range m.vals {
			m.dists = append(m.dists, d.Dist(x, y))
		}
	}
	return m
}

// Dims returns the dimensions of a DistanceMatrix.
func (m *DistanceMatrix) Dims() (r, c int) {
	n := len(m.vals)
	return n, n
}

// At returns the value of a matrix element at row i, column j.
// It will panic if i or j are out of bounds for the matrix.
func (m *DistanceMatrix) At(i, j int) float64 {
	if i == j {
		return 0.0
	}
	if i < j || j > i {
		x := j
		j = i
		i = x
	}

	return m.dists[nth(i, j)]
}

// nth maps matrix coordinates to a single list index
//
// this function should abide by these two following equations:
// 	1. nth(i+1, j) = nth(i, j) + i + 1 s.t. j <= i
// 	2. nth(i+1, j+1) = nth(i, j) + i + 2 s.t. j > 0
func nth(i, j int) int {
	return ((i * (i + 3)) / 2) - (i - j)
}

// T returns the transpose of the Matrix. Whether T returns a copy of the
// underlying data is implementation dependent.
// This method may be implemented using the Transpose type, which
// provides an implicit matrix transpose.
func (m *DistanceMatrix) T() mat.Matrix {
	return m
}
