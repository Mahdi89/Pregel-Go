package pagerank

import (
	"math/rand"
	"time"
)

// Set high-level meta parameters
const SUPER_STEPS = 100
const CONN_DEGREE = 5
const NUM_VERTEX = 100

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

// Makes a random graph
// Set the parameters above
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
			g.Vertices[i].Outgoing_edges = append(g.Vertices[i].Outgoing_edges, g.Edges[i+j])
			g.Vertices[l[i+j]].Incoming_edges = append(g.Vertices[l[i+j]].Incoming_edges, g.Edges[i+j])
		}
	}
	return g
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
