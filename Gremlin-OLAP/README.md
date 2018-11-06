
# Vertex Programs in Gremlin 

"PageRank is perhaps the most popular OLAP-oriented graph algorithm. This eigenvector centrality variant was developed by Brin and Page of Google. PageRank defines a centrality value for all vertices in the graph, where centrality is defined recursively where a vertex is central if it is connected to central vertices. PageRank is an iterative algorithm that converges to a steady state distribution. If the pageRank values are normalized to 1.0, then the pageRank value of a vertex is the probability that a random walker will be seen that that vertex in the graph at any arbitrary moment in time. In order to help developers understand the methods of a VertexProgram, the PageRankVertexProgram code is analyzed below."

Try the `pagerank.groovy` program using the ThinkerPop3 gremlin console:

```java
graph = TinkerFactory.createModern()
result = graph.compute().program(PageRankVertexProgram.build().create()).submit().get()
result.memory().runtime
g = result.graph().traversal()
g.V().valueMap()
```
output:
```
==>[gremlin.pageRankVertexProgram.pageRank:[0.11375510357865538],name:[marko],age:[29]]
==>[gremlin.pageRankVertexProgram.pageRank:[0.14598540152719103],name:[vadas],age:[27]]
==>[gremlin.pageRankVertexProgram.pageRank:[0.3047200907912249],name:[lop],lang:[java]]
==>[gremlin.pageRankVertexProgram.pageRank:[0.14598540152719103],name:[josh],age:[32]]
==>[gremlin.pageRankVertexProgram.pageRank:[0.17579889899708231],name:[ripple],lang:[java]]
==>[gremlin.pageRankVertexProgram.pageRank:[0.11375510357865538],name:[peter],age:[35]]
```

## References

- [PageRank vertex program](http://tinkerpop.apache.org/docs/current/reference/#pagerankvertexprogram)
