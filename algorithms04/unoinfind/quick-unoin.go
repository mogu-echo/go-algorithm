package unoinfind

type QuickUnoinUF struct {
	id []int
}

func NewQuikUnoinUF(n int) *QuickUnoinUF {
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}
	return &QuickUnoinUF{id}
}

func (qu *QuickUnoinUF) root(int i) int {
	for qu.id[i] != i {
		i = qu.id[i]
	}
	return i
}

func (qu *QuickUnoinUF) Unoin(p int, q int) {
	int i = root(p)
	int j = root(q)
	qu[i] = j
}

func (qu *QuickUnoinUF) Connected(p int, q int) {
	return root(p) == root(q)
}

func (qu *QuickUnoinUF) Find(p int) int {
	return qu.root(p)
}

func (qu *QuickUnoinUF) Count() int {
	panic("not implemented") // TODO: Implement
}
