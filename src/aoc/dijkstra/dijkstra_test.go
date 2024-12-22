package dijkstra

import (
	"aoc/src/testutils"
	"strings"
	"testing"
)

func TestGetEdges(t *testing.T) {
	graph := MakeGraph()
	graph.AddEdge("A", "B", 4)
	graph.AddEdge("A", "C", 5)
	graph.AddEdge("A", "D", 2)

	edges := graph.GetEdges("A")
	testutils.ExpectToMatchString(t, edges[0].Value, "B")
	testutils.ExpectToMatchString(t, edges[1].Value, "C")
	testutils.ExpectToMatchString(t, edges[2].Value, "D")

}

func TestDijkstra(t *testing.T) {
	graph := MakeGraph()
	graph.AddEdge("A", "B", 4)
	graph.AddEdge("A", "C", 5)
	graph.AddEdge("A", "D", 2)
	graph.AddEdge("B", "E", 3)
	graph.AddEdge("C", "F", 3)
	graph.AddEdge("D", "E", 5)
	graph.AddEdge("D", "F", 5)
	graph.AddEdge("D", "B", 1)
	graph.AddEdge("E", "G", 2)
	graph.AddEdge("F", "G", 3)

	cost, path := graph.GetPath("A", "G")
	pathStr := strings.Join(path, "")
	testutils.ExpectToMatchInt(t, cost, 8)
	testutils.ExpectToMatchString(t, pathStr, "ADBEG")

}

func TestDijkstraNotPossible(t *testing.T) {
	graph := MakeGraph()

	graph.AddEdge("A", "B", 4)
	graph.AddEdge("A", "C", 5)
	graph.AddEdge("A", "D", 2)
	graph.AddEdge("B", "E", 3)
	graph.AddEdge("C", "F", 3)
	graph.AddEdge("D", "E", 5)
	graph.AddEdge("D", "F", 5)
	graph.AddEdge("D", "B", 1)
	graph.AddEdge("H", "G", 2) // Isolate "G" by connecting it to "H" only
	graph.AddEdge("H", "I", 3) // Nodes "H" and "I" are disconnected from the rest of the graph

	_, path := graph.GetPath("A", "G")
	testutils.ExpectToMatchInt(t, len(path), 0)

}
