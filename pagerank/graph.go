/*
 This package provides a Pregel-like implementation of the PageRank algorithm.

 It allocates a random graph of size `NUM_VERTEX` (see config) and runs pagerank on it.

 The implementation is benchmarked against matrix- and streaming- based (single threaded) realisations, PageRank_Matrix and PageRank_Stream, respectively.
*/

package pagerank

import (
	"math/rand"
)

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
}

// Makes a random graph
// Set the parameters in config
func MakeGraph(graphSize int) Graph {

	g := Graph{}
	g.NumNodes = graphSize

	g.Edges = make([]chan float64, graphSize*CONN_DEGREE)
	g.Vertices = make([]Vertex, graphSize)

	// Initialize the edges
	for i := 0; i < graphSize*CONN_DEGREE; i++ {

		g.Edges[i] = make(chan float64, 1)
		g.Edges[i] <- float64(1 / (graphSize * CONN_DEGREE))
	}

	// Initialize the Vertix set
	for i := 0; i < graphSize; i++ {

		g.Vertices[i] = Vertex{Id: i, Value: 1.0 / float64(graphSize), Active: true, Superstep: 0}
	}

	s := rand.NewSource(123456789)
	r := rand.New(s)
	l := r.Perm(graphSize)
	// Just make sure there are enough random vertexIDs
	l = append(l, l[:CONN_DEGREE]...)

	// Assign the out-vertices after init
	// Also assign outgoing and incoming channels
	for i := 0; i < graphSize; i++ {
		for j := 0; j < CONN_DEGREE; j++ {

			g.Vertices[i].Out_vertices = append(g.Vertices[i].Out_vertices, g.Vertices[l[i+j]])
			g.Vertices[i].Outgoing_edges = append(g.Vertices[i].Outgoing_edges, g.Edges[i*CONN_DEGREE+j])
			g.Vertices[l[i+j]].Incoming_edges = append(g.Vertices[l[i+j]].Incoming_edges, g.Edges[i*CONN_DEGREE+j])
		}
	}
	return g
}
