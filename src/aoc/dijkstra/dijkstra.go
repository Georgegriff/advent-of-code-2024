package dijkstra

import (
	pqueue "aoc/src/aoc/priority_queue"
	"container/heap"
	"math"
)

type Edge struct {
	Value  string
	Weight int
}

type Graph struct {
	Nodes map[string][]*Edge
}

func MakeGraph() *Graph {
	return &Graph{
		Nodes: make(map[string][]*Edge),
	}
}

func (g *Graph) AddEdge(origin, destination string, weight int) {
	g.Nodes[origin] = append(g.Nodes[origin], &Edge{Value: destination, Weight: weight})
}

func (g *Graph) GetEdges(node string) []*Edge {
	return g.Nodes[node]
}

func (g *Graph) GetPath(origin, end string) (int, []string) {
	pq := make(pqueue.PriorityQueue, 0)
	heap.Init(&pq)

	heap.Push(&pq, &pqueue.Item{Value: origin, Priority: 0, Path: []string{origin}})

	for pq.Len() > 0 {
		topNode := heap.Pop(&pq).(*pqueue.Item)

		if topNode.Value == end {
			return topNode.Priority, topNode.Path
		}

		for _, edge := range g.GetEdges(topNode.Value) {
			newPath := make([]string, len(topNode.Path))
			copy(newPath, topNode.Path)
			newPath = append(topNode.Path, edge.Value)

			heap.Push(&pq, &pqueue.Item{Value: edge.Value, Priority: edge.Weight + topNode.Priority, Path: newPath})
		}
	}

	return int(math.Inf(1)), nil
}
