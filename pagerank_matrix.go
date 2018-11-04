// Matrix based calculation of PageRank
package pagerank

import (
	mat "github.com/skelterjohn/go.matrix"
)

func PageRank_Matrix() []float64 {
	G := mat.Zeros(NUM_VERTEX, NUM_VERTEX)
	I := mat.Eye(NUM_VERTEX)
	v := [NUM_VERTEX]Vertex{}

	// Initialize the set of Vertices
	v[0] = Vertex{Id: 0, Value: 0.166, Active: true, Superstep: 0}
	v[1] = Vertex{Id: 1, Value: 0.166, Active: true, Superstep: 0}
	v[2] = Vertex{Id: 2, Value: 0.166, Active: true, Superstep: 0}
	v[3] = Vertex{Id: 3, Value: 0.166, Active: true, Superstep: 0}
	v[4] = Vertex{Id: 4, Value: 0.166, Active: true, Superstep: 0}
	v[5] = Vertex{Id: 5, Value: 0.166, Active: true, Superstep: 0}

	v[0].Out_vertices = []Vertex{v[1], v[2], v[3]}
	v[1].Out_vertices = []Vertex{v[3]}
	v[2].Out_vertices = []Vertex{v[1], v[3]}
	v[3].Out_vertices = []Vertex{v[5]}
	v[4].Out_vertices = []Vertex{v[1], v[2], v[3]}
	v[5].Out_vertices = []Vertex{v[5]}

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
