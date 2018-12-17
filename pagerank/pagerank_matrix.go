package pagerank

import (
	mat "github.com/skelterjohn/go.matrix"
	"math/rand"
)

func PageRank_Matrix() []float64 {

	G := mat.Zeros(NUM_VERTEX, NUM_VERTEX)
	I := mat.Eye(NUM_VERTEX)
	v := make([]Vertex, NUM_VERTEX)

	// Initialize the set of Vertices
	for i := 0; i < NUM_VERTEX; i++ {

		v[i] = Vertex{Id: i, Value: 1.0 / float64(NUM_VERTEX), Active: true, Superstep: 0}
	}

	s := rand.NewSource(123456789)
	r := rand.New(s)
	l := r.Perm(NUM_VERTEX)
	// Just make sure there are enough random vertexIDs
	l = append(l, l[:CONN_DEGREE]...)

	// Assign the out-vertices after init
	for i := 0; i < NUM_VERTEX; i++ {
		for j := 0; j < CONN_DEGREE; j++ {

			v[i].Out_vertices = append(v[i].Out_vertices, v[l[i+j]])
		}
	}
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
