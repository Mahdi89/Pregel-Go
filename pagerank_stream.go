package pagerank

import (
	"math/rand"
	//	"time"
)

type Edge struct {
	Id           int
	Value        float64
	Src_Vertex_  Vertex_
	Dist_Vertex_ Vertex_
}

type Vertex_ struct {
	Id             int
	Value          float64
	Out_vertices   []Vertex_
	Incoming_edges []Edge
	Outgoing_edges []Edge
	Active         bool
	Superstep      int
}

type Graph_ struct {
	NumNodes int
	Vertices []Vertex_
	Edges    []Edge
}

// Sums up incoming values and updates the Vertex_ value.
// and passes the updated value to the neighbouring vertices.
func PageRank_(v *Vertex_) {

	sum := float64(0)
	for i := range v.Incoming_edges {
		sum += v.Incoming_edges[i].Value
	}
	v.Value = 0.15/6 + 0.85*(sum)

	len_out := len(v.Outgoing_edges)
	for j := range v.Outgoing_edges {
		v.Outgoing_edges[j].Value = v.Value / float64(len_out)
	}
}

// Makes a Graph_
// TBD: Enable making random Graph_s
func MakeGraph_(Graph_Size int) Graph_ {

	g := Graph_{}
	g.NumNodes = Graph_Size
	g.Edges = make([]Edge, Graph_Size*4)
	g.Vertices = make([]Vertex_, Graph_Size)

	// Initialize the set of Vertices
	for i := 0; i < Graph_Size; i++ {

		s := rand.NewSource(123)
		r := rand.New(s)
		g.Vertices[i] = Vertex_{Id: i, Value: 1.0 / float64(Graph_Size), Incoming_edges: []Edge{g.Edges[r.Intn(Graph_Size)], g.Edges[r.Intn(Graph_Size)]}, Outgoing_edges: []Edge{g.Edges[r.Intn(Graph_Size)], g.Edges[r.Intn(Graph_Size)]}, Active: true, Superstep: 0}
	}

	return g

}

func PageRank_Stream() []float64 {

	// Produce a random number with a deterministic seed
	s := rand.NewSource(123)
	r := rand.New(s)
	num := r.Intn(1000)

	NUM_VERTEX := num
	g := MakeGraph_(NUM_VERTEX)
	ret := make([]float64, NUM_VERTEX)

	for i := 0; i < SUPER_STEPS; i++ {
		// Streams vertices in sequence.
		for i := range g.Vertices {
			PageRank_(&g.Vertices[i])
		}
	}
	for i := range g.Vertices {
		ret[i] = g.Vertices[i].Value
	}
	return (ret)
}
