[![Build Status](https://travis-ci.com/Mahdi89/PageRank-pregel.svg?branch=master)](https://travis-ci.com/Mahdi89/PageRank-pregel)

# Pregel-Go

Pregel based (`Think like a Vertex`) realisation of a set of graph processing algorithms in Go language.

e.g in PageRank's implementation there is one gopher spawned per vertex and the communication between the vertices is done via channels. It is benchmarked against single threaded- (pipelined) and matrix- based implementations. Partitioning the graph into sub set of vertices is TBD. 
Unlike the matrix based implementation (`see pagerank_matrix.go`), which is used for testing the Pregel implementation of PageRank, the message passing based model of Pregel is considered to allow processing of large-scale Grpahs.

## Test & Benchmark 

Simply try `make bench && cat ./pagerank/bench.out` from `$project_root` to check performance of the implemented versions using Go's test machinary.  

```
goos: linux
goarch: amd64
BenchmarkPageRank_Pregel-4           100          16272149 ns/op         1087194 B/op       4267 allocs/op
BenchmarkPageRank_Matrix-4          5000            213861 ns/op           38080 B/op         84 allocs/op
BenchmarkPageRank_Stream-4           200           6080573 ns/op          928618 B/op       3618 allocs/op
PASS
ok      _/PageRank-pregel    8.118s
```

*Above figures correspond to a machine with following config:

```
processor       : 0-3
vendor_id       : GenuineIntel
cpu family      : 6
model           : 58
model name      : Intel(R) Core(TM) i7-3667U CPU @ 2.00GHz
stepping        : 9
microcode       : 0x20
cpu MHz         : 2596.236
cache size      : 4096 KB

```

## Plots

Currently `graph.go` is capable of generating random graphs with `Size` and `Degree of Sparsity` parameters as input. Later, we are going to support external graph import. To investigate performnace and memory behaviour of the implemented algorithms (namely pregel, matrix and stream) we have incorporated a plotting mechanism as following:

`$project_root/bin/bench.sh [SIZE_ITERATIONS] [CONNECTIVITY_DEGREE]` (where `SIZE_ITERATIONS` indicates range of graph sizes to be considered for benchmarking, e.g. `SIZE_ITERATIONS = 7` will produce graphs of size 5, 10, ... 35).

`$project_root/bin/plot.sh [<DIR_to_bench.out>] [SIZE_ITERATIONS]` yields plots stored in `$project_root/plot`.

### Performance 

<p align="center">
  <img src="https://github.com/Mahdi89/PageRank-pregel/blob/master/plot/bench6_5.out.png" width="280" title="Size_it= 6 and degree 5">
  <img src="https://github.com/Mahdi89/PageRank-pregel/blob/master/plot/bench7_3.out.png" width="280" title="Size_it= 7 and degree 3">
  <img src="https://github.com/Mahdi89/PageRank-pregel/blob/master/plot/bench7_4.out.png" width="280" title="Size_it= 7 and degree 4">
</p>

### Memory footprint

<p align="center">
  <img src="https://github.com/Mahdi89/PageRank-pregel/blob/master/plot/bench_mem_6_5.out.png" width="280" title="Size_it= 6 and degree 5">
  <img src="https://github.com/Mahdi89/PageRank-pregel/blob/master/plot/bench_mem_7_5.out.png" width="280" title="Size_it= 7 and degree 3">
  <img src="https://github.com/Mahdi89/PageRank-pregel/blob/master/plot/bench_mem_7_6.out.png" width="280" title="Size_it= 7 and degree 4">
</p>

### Energy usage

TBD




## References

- [Pregel: A System for Large-Scale Graph Processing](https://kowshik.github.io/JPregel/pregel_paper.pdf)
- [Thinking Like a Vertex: A Survey of Vertex-Centric Frameworks for Large-Scale Distributed Graph Processing](https://dl.acm.org/citation.cfm?id=2818185)
- The matrix-based implementation for testing the results is borrowed from: [Pregel](http://www.michaelnielsen.org/ddi/pregel/)

