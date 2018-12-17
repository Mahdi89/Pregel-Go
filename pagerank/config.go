/*
 This package provides a Pregel-like implementation of the PageRank algorithm.

 It allocates a random graph of size `NUM_VERTEX` (see config) and runs pagerank on it.

 The implementation is benchmarked against matrix- and streaming- based (single threaded) realisations, PageRank_Matrix and PageRank_Stream, respectively.
*/

package pagerank

// Set high-level meta parameters
const SUPER_STEPS = 100
const CONN_DEGREE = 4
const NUM_VERTEX = 15
