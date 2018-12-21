package shortpath

import (
	d "github.com/albertorestifo/dijkstra"
)

func ShortPath_Dijkstra() ([]string, int) {

	g := d.Graph{
		"a": {"b": 20, "c": 80},
		"b": {"a": 20, "c": 20},
		"c": {"a": 80, "b": 20},
	}

	path, cost, _ := g.Path("a", "c") // skipping error handling

	return path, cost
}
