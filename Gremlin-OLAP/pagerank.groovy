# The PageRank Vertex program in groovy based on the example provided in 
# http://tinkerpop.apache.org/docs/current/reference/#pagerankvertexprogram
# This program uses Gremlin-OLAP (Pregel based) traversals to compute the 
# PageRank algorithm for pre constructed graph of 6 vertices.

graph = TinkerFactory.createModern()
result = graph.compute().program(PageRankVertexProgram.build().create()).submit().get()
result.memory().runtime
g = result.graph().traversal()
g.V().valueMap()
