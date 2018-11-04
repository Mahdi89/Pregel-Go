package pagerank

import (
	"math"
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
