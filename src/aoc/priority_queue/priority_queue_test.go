package pqueue

import (
	"aoc/src/testutils"
	"container/heap"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	// Create a priority queue and add some items.
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	heap.Push(&pq, &Item{Value: "Task 1", Priority: 3})
	heap.Push(&pq, &Item{Value: "Task 2", Priority: 1})
	heap.Push(&pq, &Item{Value: "Task 3", Priority: 2})

	item := heap.Pop(&pq).(*Item)
	testutils.ExpectToMatchString(t, item.Value, "Task 2")
	item = heap.Pop(&pq).(*Item)
	testutils.ExpectToMatchString(t, item.Value, "Task 3")
	item = heap.Pop(&pq).(*Item)
	testutils.ExpectToMatchString(t, item.Value, "Task 1")

}
