package metric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func toStream(xs []float64) <-chan float64 {
	xsChan := make(chan float64)
	go func() {
		defer close(xsChan)
		for _, x := range xs {
			xsChan <- x
		}
	}()
	return xsChan
}

func TestCompute(t *testing.T) {
	stream := toStream([]float64{3.0, 4.0, 5.0})

	// Compute distance matrix based on stream of data
	m := Compute(stream, Euclidean)

	// Assert expected dimensions
	rows, cols := m.Dims()
	if !assert.Equal(t, 3, rows) {
		return
	}
	if !assert.Equal(t, 3, cols) {
		return
	}

	// Assert expected distances
	expected := [][]float64{
		{0.0, 5.0, 5.830951894845301},
		{5.0, 0.0, 6.4031242374328485},
		{5.830951894845301, 6.4031242374328485, 0.0},
	}
	for i, row := range expected {
		for j, d := range row {
			if !assert.Equal(t, d, m.At(i, j)) {
				return
			}
		}
	}

	// Assert m = mT since distance matricies are symmetric
	mT := m.T()
	for i, row := range expected {
		for j, d := range row {
			if !assert.Equal(t, d, mT.At(i, j)) {
				return
			}
		}
	}
}

func FuzzNth(f *testing.F) {
	// seed corpus with 10,000 matrix coordinate pairs
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			f.Add(i, j)
		}
	}

	// fuzz on nth function laws
	f.Fuzz(func(t *testing.T, i, j int) {
		if i < 0 || j < 0 {
			t.Skip()
			return
		}
		if j <= i {
			if !assert.Equal(t, nth(i+1, j), nth(i, j)+i+1) {
				return
			}
		}
		if j > 0 {
			if !assert.Equal(t, nth(i+1, j+1), nth(i, j)+i+2) {
				return
			}
		}
	})
}

func BenchmarkNth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := nth(i, i)
		if i > 0 && n != nth(i-1, i-1)+i+1 {
			b.Fail()
			return
		}
	}
}
