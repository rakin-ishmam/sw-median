package median

type node struct {
	val  int
	qind int
	qid  int
}

func (n *node) setQInd(i int) {
	n.qind = i
}

func (n *node) setQID(id int) {
	n.qid = id
}
