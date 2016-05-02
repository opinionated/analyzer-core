package analyzer

import (
	"container/heap"
)

// Heap implements an Analyzable min heap (item with lowest score at the top)
type Heap []*Analyzable

// Peek into the next item
func (pq Heap) Peek() *Analyzable {
	if pq.Len() == 0 {
		return nil
	}
	ret := (pq)[0]
	return ret
}

// Len of the queue
func (pq Heap) Len() int { return len(pq) }

// Less is the comparator
func (pq Heap) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, Score so we use greater than here.
	return pq[i].Score < pq[j].Score
}

// Swap off
func (pq Heap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push on
func (pq *Heap) Push(x interface{}) {
	item := x.(*Analyzable)
	*pq = append(*pq, item)
}

// Pop off
func (pq *Heap) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

var _ heap.Interface = (*Heap)(nil)
