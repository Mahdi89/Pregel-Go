package shortpath

import (
	"reflect"
	"runtime"
	"testing"
)

func TestShortPath_Dijkstra(t *testing.T) {

	// Expected return vector from Dijkstra algorithm
	expected_path, expected_val := ShortPath_Dijkstra()
	// Returned from the X implementation
	ret_path := []string{"a", "b", "c"}
	ret_val := 40

	if !reflect.DeepEqual(expected_path, ret_path) {
		t.Errorf("expected %v but got %v", expected_path, ret_path)
	}
	if expected_val != ret_val {
		t.Errorf("expected %d but got %d", expected_val, ret_val)
	}

}

func BenchmarkShortPath_Dijkstra(b *testing.B) {
	runtime.GC()
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		ShortPath_Dijkstra()
	}
}
