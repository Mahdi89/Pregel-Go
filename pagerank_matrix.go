// Matrix based calculation of PageRank
package pagerank

import (
	// 	"fmt"
	mat "github.com/skelterjohn/go.matrix"
	"math/rand"
	//	"time"
)

//var NUM_VERTEX = 6
//var SUPER_STEPS = 5
/*
type Vertex struct {
	Id             int
	Value          float64
	Out_vertices   []Vertex
	Incoming_edges []chan float64
	Outgoing_edges []chan float64
	Active         bool
	Superstep      int
}

type Graph struct {
	NumNodes int
	Vertices []Vertex
	Edges    []chan float64
}*/

func PageRank_Matrix() []float64 {
	//func main() {

	// Produce a random number with a deterministic seed
	s := rand.NewSource(123)
	r := rand.New(s)
	num := r.Intn(1000)

	NUM_VERTEX := num

	G := mat.Zeros(NUM_VERTEX, NUM_VERTEX)
	I := mat.Eye(NUM_VERTEX)
	v := make([]Vertex, NUM_VERTEX)

	// Initialize the set of Vertices
	for i := 0; i < NUM_VERTEX; i++ {

		v[i] = Vertex{Id: i, Value: 1.0 / float64(NUM_VERTEX), Out_vertices: []Vertex{v[r.Intn(NUM_VERTEX)], v[r.Intn(NUM_VERTEX)], v[r.Intn(NUM_VERTEX)], v[r.Intn(NUM_VERTEX)]}, Active: true, Superstep: 0}
	}
	// Assign the out-vertices after init
	for i := 0; i < NUM_VERTEX; i++ {

		v[i].Out_vertices = []Vertex{v[r.Intn(NUM_VERTEX)], v[r.Intn(NUM_VERTEX)]}
	}

	/*	// Pretty print the adjacency list
		for i := 0; i < NUM_VERTEX; i++ {
			for j := range v[i].Out_vertices {
				fmt.Print(v[i].Out_vertices[j].Id)
			}
			fmt.Println()
		}
	*/
	for i := 0; i < NUM_VERTEX; i++ {
		num_out_vertices := len(v[i].Out_vertices)
		for j := range v[i].Out_vertices {
			G.Set(v[i].Out_vertices[j].Id, v[i].Id, 1.0/float64(num_out_vertices))
		}
	}
	O := mat.Ones(NUM_VERTEX, 1)
	for i := 0; i < NUM_VERTEX; i++ {
		O.Set(i, 0, (1.0 / float64(NUM_VERTEX)))
	}
	for i := 0; i < NUM_VERTEX; i++ {
		for j := 0; j < NUM_VERTEX; j++ {
			G.Set(i, j, G.Get(i, j)*0.85)
		}
	}
	_ = I.SubtractDense(G)
	part, _ := I.Inverse()
	for i := 0; i < NUM_VERTEX; i++ {
		for j := 0; j < NUM_VERTEX; j++ {
			part.Set(i, j, part.Get(i, j)*0.15)
		}
	}
	ret, _ := part.TimesDense(O)
	return ret.Array()
}
