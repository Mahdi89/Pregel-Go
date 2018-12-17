/*
 This package provides a Pregel-like implementation of the PageRank algorithm.

 It allocates a random graph of size `NUM_VERTEX` (see config) and runs pagerank on it.

 The implementation is benchmarked against matrix- and streaming- based (single threaded) realisations, PageRank_Matrix and PageRank_Stream, respectively.
*/

package pagerank

import (
	"time"
)

// Sums up incoming values and updates the vertex value.
// and passes the updated value to the neighbouring vertices.
func PageRank(g *Graph, t int) {

	sum := float64(0)
	for i := range g.Vertices[t].Incoming_edges {
		sum += <-g.Vertices[t].Incoming_edges[i]
	}
	g.Vertices[t].Value = 0.15/float64(g.NumNodes) + 0.85*(sum)

	len_out := len(g.Vertices[t].Outgoing_edges)
	for j := range g.Vertices[t].Outgoing_edges {
		g.Vertices[t].Outgoing_edges[j] <- g.Vertices[t].Value / float64(len_out)
	}
}

func PageRank_Pregel() []float64 {

	g := MakeGraph(NUM_VERTEX)
	ret := make([]float64, NUM_VERTEX)

	for i := 0; i < SUPER_STEPS; i++ {
		// Spawns a gopher per vertex,
		// partitioning the graph into set of vertices is TBD.
		for j := range g.Vertices {
			go PageRank(&g, j)
		}
		// Wait for the gophers to stablize.
		time.Sleep(time.Microsecond)
	}
	for i := range g.Vertices {
		ret[i] = g.Vertices[i].Value
	}
	return ret
}
