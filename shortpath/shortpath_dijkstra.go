package shortpath

import (
	d "github.com/albertorestifo/dijkstra"
)

func ShortPath_Dijkstra(src, dist string) ([]string, int) {

	// TODO Allow working on large-scale graphs
	// by either reading in or generating random graphs
	g := d.Graph{
		"a": {"b": 20, "c": 80},
		"b": {"a": 20, "c": 20},
		"c": {"a": 80, "b": 20},
		"d": {"a": 80, "c": 20},
	}

	path, cost, _ := g.Path(src, dist) // skipping error handling

	return path, cost
}
