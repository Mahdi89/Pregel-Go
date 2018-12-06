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
