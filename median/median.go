package median

import (
	"container/heap"

	"github.com/rakin-ishmam/sw-median/ds/queue"
)

const (
	minqid int = 1
	maxqid     = 2
)

type median struct {
	minq  pqueue
	maxq  pqueue
	q     queue.Queue
	wsize int
}

func (m *median) init() {
	m.minq = qinit(minqid, m.wsize/2+1, false)
	m.maxq = qinit(maxqid, m.wsize/2+1, true)
	m.q = queue.New()
}

func (m *median) balance() {
	for {
		if m.minq.Len() > m.maxq.Len() {
			n := heap.Pop(&m.minq).(*node)
			heap.Push(&m.maxq, n)
		}
		break
	}

	for {
		if m.maxq.Len()-1 > m.minq.Len() {
			n := heap.Pop(&m.maxq).(*node)
			heap.Push(&m.minq, n)
		}
		break
	}
}

func (m *median) removeFst() {
	v, _ := m.q.Front()
	m.q.Pop()
	n := v.(*node)

	if n.qid == minqid {
		heap.Remove(&m.minq, n.qind)
		return
	}

	heap.Remove(&m.maxq, n.qind)
}

func (m *median) AddDelay(d int) {
	if m.q.Size() == m.wsize {
		m.removeFst()
		m.balance()
	}

	n := node{
		val: d,
	}

	m.q.Push(&n)

	if m.maxq.Len() == 0 {
		heap.Push(&m.maxq, &n)
		return
	}

	mf := m.maxq.Front()
	if d <= mf.val {
		heap.Push(&m.maxq, &n)
	} else {
		heap.Push(&m.minq, &n)
	}

	m.balance()

	// fmt.Println("dealy", d)
	// fmt.Println("mxq")
	// m.maxq.prt()
	// fmt.Println("mmq")
	// m.minq.prt()
	// fmt.Println()

}

func (m median) GetMedian() float64 {
	if m.q.Size() < 2 {
		return -1
	}

	if m.q.Size()%2 == 1 {
		return float64(m.maxq.Front().val)
	}

	return float64(m.maxq.Front().val+m.minq.Front().val) / 2.0
}

// New returns new instance for Medianer
func New(windowSize int) Medianer {
	m := median{
		wsize: windowSize,
	}

	m.init()

	return &m
}
