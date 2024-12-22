package pqueue

import (
	"container/heap"
)

type Item struct {
	Value    string
	Priority int
	index    int
	// Custom field
	Path []string
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

// Less defines the order of the heap. For a min-heap, we use priority comparison.
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

// Swap swaps the elements at indices i and j.
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push adds an item to the heap.
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop removes and returns the smallest item from the heap.
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // Avoid memory leak.
	item.index = -1 // Mark the item as removed.
	*pq = old[0 : n-1]
	return item
}

// Update modifies the priority and value of an item in the queue.
func (pq *PriorityQueue) Update(item *Item, value string, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.index)
}
