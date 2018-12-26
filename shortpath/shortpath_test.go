package shortpath

import (
	//	"reflect"
	"runtime"
	"testing"
)

func TestShortPath_Dijkstra(t *testing.T) {

	// Expected return vector from Dijkstra algorithm
	_, expected_val := ShortPath_Dijkstra("a", "c")
	// Returned from the X implementation
	_, ret_val := ShortPath_Pregel("a", "c")

	// TODO Disable 'path' return
	/*	if !reflect.DeepEqual(expected_path, ret_path) {
			t.Errorf("expected %v but got %v", expected_path, ret_path)
		}
	*/
	if expected_val != ret_val {
		t.Errorf("expected %d but got %d", expected_val, ret_val)
	}

}

func BenchmarkShortPath_Dijkstra(b *testing.B) {
	runtime.GC()
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		ShortPath_Dijkstra("a", "c")
	}
}
func BenchmarkShortPath_Pregel(b *testing.B) {
	runtime.GC()
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		ShortPath_Pregel("a", "c")
	}
}
