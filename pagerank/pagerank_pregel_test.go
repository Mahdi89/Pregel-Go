/*
 This package provides a Pregel-like implementation of the PageRank algorithm.

 It allocates a random graph of size `NUM_VERTEX` (see config) and runs pagerank on it.

 The implementation is benchmarked against matrix- and streaming- based (single threaded) realisations, PageRank_Matrix and PageRank_Stream, respectively.
*/

package pagerank

import (
	"math"
	"runtime"
	"testing"
)

func TestPageRank_Pregel(t *testing.T) {

	// Expected return vector
	expected := PageRank_Matrix()
	// Calculated pregel vector
	ret := PageRank_Pregel()

	expected_sum := float64(0)
	ret_sum := float64(0)

	for i := range expected {

		expected_sum += expected[i]
		ret_sum += ret[i]
	}
	if math.Abs(ret_sum-expected_sum) > 0.01 {
		t.Errorf("Expected %f got %f", expected, ret)
	}
}

func TestPageRank_Stream(t *testing.T) {

	// Expected return vector
	expected := PageRank_Matrix()
	// Calculated pregel vector
	ret := PageRank_Stream()

	expected_sum := float64(0)
	ret_sum := float64(0)

	for i := range expected {

		expected_sum += expected[i]
		ret_sum += ret[i]
	}
	if math.Abs(ret_sum-expected_sum) > 0.01 {
		t.Errorf("Expected %f got %f", expected, ret)
	}
}

func BenchmarkPageRank_Pregel(b *testing.B) {
	runtime.GC()
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		PageRank_Pregel()
	}
}

func BenchmarkPageRank_Matrix(b *testing.B) {
	runtime.GC()
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		PageRank_Matrix()
	}
}

func BenchmarkPageRank_Stream(b *testing.B) {
	runtime.GC()
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		PageRank_Stream()
	}
}
