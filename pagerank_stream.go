package pagerank

func PageRank_Stream() []float64 {

	g := MakeGraph(NUM_VERTEX)
	ret := make([]float64, NUM_VERTEX)

	for i := 0; i < SUPER_STEPS; i++ {
		// Stream vertices through
		// partitioning the graph into set of vertices is TBD.
		go func() {
			for j := range g.Vertices {
				go PageRank(&g, j)
			}
		}()
	}
	for i := range g.Vertices {
		ret[i] = g.Vertices[i].Value
	}
	return ret

}
