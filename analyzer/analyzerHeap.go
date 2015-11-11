package analyzer

import (
	"container/heap"
)

// Heap implements an Analyzable min heap (item with lowest score at the top)
type Heap []*Analyzable

func (pq Heap) Peek() *Analyzable {
	if pq.Len() == 0 {
		return nil
	}
	ret := (pq)[0]
	return ret
}

func (pq Heap) Len() int { return len(pq) }

func (pq Heap) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, Score so we use greater than here.
	return pq[i].Score < pq[j].Score
}

func (pq Heap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *Heap) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Analyzable)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *Heap) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the Score and value of an Analyzable in the queue.
func (pq *Heap) Update(item *Analyzable, value string, Score float64) {
	item.Name = value
	item.Score = Score
	heap.Fix(pq, item.index)
}

var _ heap.Interface = (*Heap)(nil)
