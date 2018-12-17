/*
 This package provides a Pregel-like implementation of the PageRank algorithm.

 It allocates a random graph of size `NUM_VERTEX` (see config) and runs pagerank on it.

 The implementation is benchmarked against matrix- and streaming- based (single threaded) realisations, PageRank_Matrix and PageRank_Stream, respectively.
*/

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
