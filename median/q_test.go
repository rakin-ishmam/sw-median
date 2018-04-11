package median

import (
	"container/heap"
	"testing"
)

func TestMaxPQ(t *testing.T) {
	q := qinit(1, 100, true)
	tt := []struct {
		name    string
		action  func()
		exFront int
	}{
		{"push 1", func() {
			heap.Push(&q, &node{val: 1})
		}, 1},
		{"push 10", func() {
			heap.Push(&q, &node{val: 10})
		}, 10},
		{"push 4", func() {
			heap.Push(&q, &node{val: 4})
		}, 10},
		{"push 8", func() {
			heap.Push(&q, &node{val: 8})
		}, 10},
		{"pop 10", func() {
			heap.Pop(&q)
		}, 8},
		{"pop 8", func() {
			heap.Pop(&q)
		}, 4},
		{"push 40", func() {
			heap.Push(&q, &node{val: 40})
		}, 40},
	}

	for _, st := range tt {
		t.Run(st.name, func(t *testing.T) {
			st.action()
			frnt := q.Front().val
			if frnt != st.exFront {
				t.Errorf("expected front value is %v but got %v", st.exFront, frnt)
			}
		})
	}
}

func TestMinPQ(t *testing.T) {
	q := qinit(1, 100, false)
	tt := []struct {
		name    string
		action  func()
		exFront int
	}{
		{"push 1", func() {
			heap.Push(&q, &node{val: 1})
		}, 1},
		{"push 10", func() {
			heap.Push(&q, &node{val: 10})
		}, 1},
		{"push 4", func() {
			heap.Push(&q, &node{val: 4})
		}, 1},
		{"push 8", func() {
			heap.Push(&q, &node{val: 8})
		}, 1},
		{"pop 1", func() {
			heap.Pop(&q)
		}, 4},
		{"pop 4", func() {
			heap.Pop(&q)
		}, 8},
		{"push 40", func() {
			heap.Push(&q, &node{val: 40})
		}, 8},
	}

	for _, st := range tt {
		t.Run(st.name, func(t *testing.T) {
			st.action()
			frnt := q.Front().val
			if frnt != st.exFront {
				t.Errorf("expected front value is %v but got %v", st.exFront, frnt)
			}
		})
	}
}
