package pagerank

import ()

type Edge struct {
	Id          int
	Value       float64
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
	g.Edges = make([]Edge, 11)

	// Initialize the edges
	g.Edges[0].Value = 0.055
	g.Edges[1].Value = 0.055
	g.Edges[2].Value = 0.055
	g.Edges[3].Value = 0.166
	g.Edges[4].Value = 0.083
	g.Edges[5].Value = 0.083
	g.Edges[6].Value = 0.166
	g.Edges[7].Value = 0.055
	g.Edges[8].Value = 0.055
	g.Edges[9].Value = 0.055
	g.Edges[10].Value = 0.166

	g.Vertices = []Vertex_{
		// Initialize a set of Vertices
		Vertex_{Id: 0, Value: 0.166, Incoming_edges: []Edge{}, Outgoing_edges: []Edge{g.Edges[0], g.Edges[1], g.Edges[2]}, Active: true, Superstep: 0},
		Vertex_{Id: 1, Value: 0.166, Incoming_edges: []Edge{g.Edges[0], g.Edges[4], g.Edges[7]}, Outgoing_edges: []Edge{g.Edges[3]}, Active: true, Superstep: 0},
		Vertex_{Id: 2, Value: 0.166, Incoming_edges: []Edge{g.Edges[1], g.Edges[8]}, Outgoing_edges: []Edge{g.Edges[4], g.Edges[5]}, Active: true, Superstep: 0},
		Vertex_{Id: 3, Value: 0.166, Incoming_edges: []Edge{g.Edges[2], g.Edges[3], g.Edges[5], g.Edges[9]}, Outgoing_edges: []Edge{g.Edges[6]}, Active: true, Superstep: 0},
		Vertex_{Id: 4, Value: 0.166, Incoming_edges: []Edge{}, Outgoing_edges: []Edge{g.Edges[7], g.Edges[8], g.Edges[9]}, Active: true, Superstep: 0},
		Vertex_{Id: 5, Value: 0.166, Incoming_edges: []Edge{g.Edges[10], g.Edges[6]}, Outgoing_edges: []Edge{g.Edges[10]}, Active: true, Superstep: 0}}
	return g

}

func PageRank_Stream() []float64 {

	g := MakeGraph_(6)
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
