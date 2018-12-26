package shortpath

import (
	"time"
)

// TODO Allow working on large-scale graphs
// by either reading in or generating random graphs
func TestGraph(graphSize int) Graph {

	g := Graph{}
	g.NumNodes = graphSize

	g.Edges = make([]chan Msg, graphSize*CONN_DEGREE)
	g.Vertices = make([]Vertex, graphSize)

	// Initialize the edges
	for i := 0; i < graphSize*CONN_DEGREE; i++ {

		g.Edges[i] = make(chan Msg, 1)
	}
	g.Edges[0] <- Msg{Value: 0, Weight: 20}
	g.Edges[1] <- Msg{Value: 0, Weight: 80}
	g.Edges[2] <- Msg{Value: INFINITY, Weight: 20}
	g.Edges[3] <- Msg{Value: INFINITY, Weight: 20}
	g.Edges[4] <- Msg{Value: INFINITY, Weight: 80}
	g.Edges[5] <- Msg{Value: INFINITY, Weight: 20}
	g.Edges[6] <- Msg{Value: INFINITY, Weight: 80}
	g.Edges[7] <- Msg{Value: INFINITY, Weight: 20}

	// Initialize the Vertix set
	for i := 0; i < graphSize; i++ {

		g.Vertices[i] = Vertex{Id: i, Value: INFINITY, Active: true, Superstep: 0}
	}

	// Assign Connectivities
	g.Vertices[0].Outgoing_edges = []chan Msg{g.Edges[0], g.Edges[1]}
	g.Vertices[0].Incoming_edges = []chan Msg{g.Edges[2], g.Edges[4], g.Edges[6]}

	g.Vertices[1].Outgoing_edges = []chan Msg{g.Edges[2], g.Edges[3]}
	g.Vertices[1].Incoming_edges = []chan Msg{g.Edges[0], g.Edges[5]}

	g.Vertices[2].Outgoing_edges = []chan Msg{g.Edges[4], g.Edges[5]}
	g.Vertices[2].Incoming_edges = []chan Msg{g.Edges[3], g.Edges[1], g.Edges[7]}

	g.Vertices[3].Outgoing_edges = []chan Msg{g.Edges[6], g.Edges[7]}
	g.Vertices[3].Incoming_edges = []chan Msg{}

	return g
}
func ShortPath_Pregel(src, dist string) ([]string, int) {

	g := TestGraph(4)
	str := []string{}

	// map nodes
	m := make(map[string]int)
	m["a"] = 0
	m["b"] = 1
	m["c"] = 2
	m["d"] = 3

	// Initiate the 'src' node
	g.Vertices[m[src]].Value = 0

	for i := 1; i < g.NumNodes; i++ {
		go func(i int) {
			for j := range g.Vertices[i].Incoming_edges {

				msg := <-g.Vertices[i].Incoming_edges[j]
				tmp := msg.Value + msg.Weight
				if tmp < g.Vertices[i].Value {
					g.Vertices[i].Value = tmp
				}
			}
			for k := range g.Vertices[i].Outgoing_edges {

				msg := <-g.Vertices[i].Outgoing_edges[k]
				g.Vertices[i].Outgoing_edges[k] <- Msg{Value: g.Vertices[i].Value, Weight: msg.Weight}
			}
		}(i)
	}
	time.Sleep(time.Millisecond)

	return str, int(g.Vertices[m[dist]].Value)
}
