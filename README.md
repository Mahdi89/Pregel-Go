[![Build Status](https://travis-ci.com/Mahdi89/PageRank-pregel.svg?branch=master)](https://travis-ci.com/Mahdi89/PageRank-pregel)

# PageRank-pregel

Pregel based (`Think like a Vertex`) realisation of the PageRank algorithm in Go language.

In this implementation there is one gopher spawned per vertex and the communication between the vertices is done via channels. Partitioning the graph into sub set of vertices is TBD. 

Unlike the matrix based implementation (`see pagerank_matrix.go`), which is used for testing the Pregel implementation of PageRank, the message passing based model of Pregel allows processing of large-scale Grpahs.

Currently the tested structure is a simple 6 node, 11 edge graph and the work for implementing random structures such as ones suggested by graph nets is WIP.  

## References

- [Pregel: A System for Large-Scale Graph Processing](https://kowshik.github.io/JPregel/pregel_paper.pdf)
- [Thinking Like a Vertex: A Survey of Vertex-Centric Frameworks for Large-Scale Distributed Graph Processing](https://dl.acm.org/citation.cfm?id=2818185)
- The matrix-based implementation for testing the results is borrowed from: [Pregel](http://www.michaelnielsen.org/ddi/pregel/)

