package analyzer_test

import (
	"container/heap"
	"github.com/opinionated/analyzer-core/analyzer"
	"testing"
)

func TestAnalyzableHeap(t *testing.T) {
	items := map[string]float64{
		"d": 5.0, "c": 7.0, "e": 2.0, "b": 9.0, "g": 1.0,
	}

	mheap := make(analyzer.Heap, len(items))

	i := 0
	for key, value := range items {
		analyzable := analyzer.BuildAnalyzable()
		analyzable.Name = key
		analyzable.Score = value

		mheap[i] = &analyzable
		i++
	}

	heap.Init(&mheap)

	a := analyzer.BuildAnalyzable()
	a.Name = "a"
	a.Score = 6.0
	heap.Push(&mheap, &a)
	mheap.Update(&a, a.Name, 15.0)

	b := analyzer.BuildAnalyzable()
	b.Name = "q"
	b.Score = 6.5
	heap.Push(&mheap, &b)
	mheap.Update(&b, "f", 1.5)

	// pop all the items off, they should be in order
	name := 'g'
	for mheap.Len() > 0 {

		peekedAnalyzable := mheap.Peek()
		nextAnalyzable := heap.Pop(&mheap).(*analyzer.Analyzable)

		if nextAnalyzable.Name != string(name) {
			t.Errorf("heap order error, expected: %s got: %s\n",
				string(name), nextAnalyzable.Name)
		}

		if peekedAnalyzable.Name != nextAnalyzable.Name {
			t.Errorf("heap peek broke! peeked: %s popped: %s\n",
				peekedAnalyzable, nextAnalyzable)
		}

		name--
	}
}
