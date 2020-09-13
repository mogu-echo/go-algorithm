package unoinfind

type QuikFindUF struct {
	id []int
}

func NewQuikFindUF(n int) *QuikFindUF {
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}
	return &QuikFindUF{id}
}

func (qf *QuikFindUF) Unoin(p int, q int) {
	pid := qf.id[p]
	qid := qf.id[q]
	for i := 0; i < len(qf.id); i++ {
		if qf.id[i] == pid {
			qf.id[i] = qid
		}
	}
}

func (qf *QuikFindUF) Connected(p int, q int) {
	return qf.id[p] == qf.id[q]
}

func (qf *QuikFindUF) Find(p int) int {
	return qf.id[p]
}

func (qf *QuikFindUF) Count() int {
	panic("not implemented") // TODO: Implement
}
