package median

import (
	"container/heap"
	"fmt"
)

type pqueue struct {
	nodes []*node
	len   int
	isMax bool
	id    int
}

func (pq pqueue) Len() int {
	return pq.len
}
func (pq pqueue) Less(i, j int) bool {
	if pq.isMax {
		return pq.nodes[i].val > pq.nodes[j].val
	}
	return pq.nodes[i].val < pq.nodes[j].val
}

func (pq *pqueue) Swap(i, j int) {
	pq.nodes[i], pq.nodes[j] = pq.nodes[j], pq.nodes[i]
	pq.nodes[i].setQInd(i)
	pq.nodes[j].setQInd(j)
}

func (pq *pqueue) Push(x interface{}) {
	n := x.(*node)
	n.setQInd(pq.len)
	n.setQID(pq.id)
	pq.nodes[pq.len] = n
	pq.len++
}

func (pq *pqueue) Pop() interface{} {
	pq.len--
	return pq.nodes[pq.len]
}

func (pq *pqueue) Front() *node {
	return pq.nodes[0]
}

func (pq pqueue) prt() {
	for i, n := range pq.nodes {
		if i >= pq.len {
			break
		}
		fmt.Print(n.val, " ")
	}
	fmt.Println()
}

func qinit(id, limit int, isMxQ bool) pqueue {
	pq := pqueue{
		isMax: isMxQ,
		nodes: make([]*node, limit),
		id:    id,
		len:   0,
	}

	heap.Init(&pq)

	return pq
}
