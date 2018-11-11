package pagerank

import (
	"math/rand"
	"time"
)

const NUM_VERTEX = 6
const SUPER_STEPS = 5

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

// Sums up incoming values and updates the vertex value.
// and passes the updated value to the neighbouring vertices.
func PageRank(v *Vertex) {

	sum := float64(0)
	for i := range v.Incoming_edges {
		sum += <-v.Incoming_edges[i]
	}
	v.Value = 0.15/6 + 0.85*(sum)

	len_out := len(v.Outgoing_edges)
	for j := range v.Outgoing_edges {
		v.Outgoing_edges[j] <- v.Value / float64(len_out)
	}
}

// Makes a graph
// TBD: Enable making random graphs
func MakeGraph(graphSize int) Graph {

	g := Graph{}
	g.NumNodes = graphSize

	g.Edges = make([]chan float64, graphSize*4)
	g.Vertices = make([]Vertex, graphSize)

	// Initialize the set of Vertices
	for i := 0; i < graphSize; i++ {

		s := rand.NewSource(123)
		r := rand.New(s)
		g.Vertices[i] = Vertex{Id: i, Value: 1.0 / float64(graphSize), Incoming_edges: []chan float64{g.Edges[r.Intn(graphSize)], g.Edges[r.Intn(graphSize)]}, Outgoing_edges: []chan float64{g.Edges[r.Intn(graphSize)], g.Edges[r.Intn(graphSize)]}, Active: true, Superstep: 0}
	}

	return g

}

func PageRank_Pregel() []float64 {

	// Produce a random number with a deterministic seed
	s := rand.NewSource(123)
	r := rand.New(s)
	num := r.Intn(1000)

	NUM_VERTEX := num
	g := MakeGraph(NUM_VERTEX)
	ret := make([]float64, NUM_VERTEX)

	for i := 0; i < SUPER_STEPS; i++ {
		// Spawns a gopher per vertex,
		// partitioning the graph into set of vertices is TBD.
		for i := range g.Vertices {
			go PageRank(&g.Vertices[i])
		}
		// Wait for the gophers to stablize.
		time.Sleep(time.Nanosecond)
	}
	for i := range g.Vertices {
		ret[i] = g.Vertices[i].Value
	}

	return ret
}
