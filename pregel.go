package main

import (
	"fmt"
	"time"
)

var NUM_VERTEX = 6
var SUPER_STEPS = 10

type Vertex struct {
	Id             int
	Value          float32
	Out_vertices   []Vertex
	Incoming_edges []chan float32
	Outgoing_edges []chan float32
	Active         bool
	Superstep      int
}

type Graph struct {
	NumNodes int
	Vertices []Vertex
	Edges    []chan float32
}

// Sums up incoming values and updates the vertex value.
// and passes the updated value to the neighbouring vertices.
func PageRank(v *Vertex) {

	sum := float32(0)
	for i := range v.Incoming_edges {
		sum += <-v.Incoming_edges[i]
	}
	v.Value = 0.15/6 + 0.85*(sum)

	len_out := len(v.Outgoing_edges)
	for j := range v.Outgoing_edges {
		v.Outgoing_edges[j] <- v.Value / float32(len_out)
	}
}

// Makes a graph
// TBD: Enable making random graphs
func MakeGraph(graphSize int) Graph {

	g := Graph{}
	g.NumNodes = graphSize
	g.Edges = []chan float32{
		make(chan float32, 1),
		make(chan float32, 1),
		make(chan float32, 1),
		make(chan float32, 1),
		make(chan float32, 1),
		make(chan float32, 1),
		make(chan float32, 1),
		make(chan float32, 1),
		make(chan float32, 1),
		make(chan float32, 1),
		make(chan float32, 1)}

	// Initialize the edges
	g.Edges[0] <- 0.055
	g.Edges[1] <- 0.055
	g.Edges[2] <- 0.055
	g.Edges[3] <- 0.166
	g.Edges[4] <- 0.083
	g.Edges[5] <- 0.083
	g.Edges[6] <- 0.166
	g.Edges[7] <- 0.055
	g.Edges[8] <- 0.055
	g.Edges[9] <- 0.055
	g.Edges[10] <- 0.166

	g.Vertices = []Vertex{
		// Initialize a set of Vertices
		Vertex{Id: 0, Value: 0.166, Incoming_edges: []chan float32{}, Outgoing_edges: []chan float32{g.Edges[0], g.Edges[1], g.Edges[2]}, Active: true, Superstep: 0},
		Vertex{Id: 1, Value: 0.166, Incoming_edges: []chan float32{g.Edges[0], g.Edges[4], g.Edges[7]}, Outgoing_edges: []chan float32{g.Edges[3]}, Active: true, Superstep: 0},
		Vertex{Id: 2, Value: 0.166, Incoming_edges: []chan float32{g.Edges[1], g.Edges[8]}, Outgoing_edges: []chan float32{g.Edges[4], g.Edges[5]}, Active: true, Superstep: 0},
		Vertex{Id: 3, Value: 0.166, Incoming_edges: []chan float32{g.Edges[2], g.Edges[3], g.Edges[5], g.Edges[9]}, Outgoing_edges: []chan float32{g.Edges[6]}, Active: true, Superstep: 0},
		Vertex{Id: 4, Value: 0.166, Incoming_edges: []chan float32{}, Outgoing_edges: []chan float32{g.Edges[7], g.Edges[8], g.Edges[9]}, Active: true, Superstep: 0},
		Vertex{Id: 5, Value: 0.166, Incoming_edges: []chan float32{g.Edges[10], g.Edges[6]}, Outgoing_edges: []chan float32{g.Edges[10]}, Active: true, Superstep: 0}}
	return g

}

func main() {

	g := MakeGraph(6)

	for i := 0; i < SUPER_STEPS; i++ {
		// Spawns a gopher per vertex,
		// partitioning the graph into set of vertices is TBD.
		fmt.Printf("--------SuperStep %d -------------------------\n", i)
		for i := range g.Vertices {
			go PageRank(&g.Vertices[i])
		}
		// Wait for the gophers to stablize.
		time.Sleep(time.Second)
		for i := range g.Vertices {
			fmt.Println(g.Vertices[i].Value)
		}
	}
}