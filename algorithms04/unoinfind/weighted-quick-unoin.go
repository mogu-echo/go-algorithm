package unoinfind

type WeightedQuickUnoinUF struct {
	id []int
	sz []int
	count int
}

func NewWeightedQuickUnoinUF(n int) *WeightedQuickUnoinUF {
	id := make([]int, n)
	count = n
	for i := 0; i < n; i++ {
		id[i] = i
	}
	sz := make([]int, n)
	for i := range sz {
		sz[i] = 1
	}
	return &WeightedQuickUnoinUF{id, sz, count}
}

func (qu *WeightedQuickUnoinUF) root(int i) int {
	for qu.id[i] != i {
		qu.id[i] = qu.id[qu.id[i]]  // Path compression
		i = qu.id[i]
	}
	return i
}

func (qu *WeightedQuickUnoinUF) Unoin(p int, q int) {
	int i = root(p)
	int j = root(q)
	if qu.sz[i] < qu.sz[j] {
		qu.id[i] = j
		qu.sz[j] += qu.sz[i]
	} else {
		qu.id[j] = i
		qu.sz[i] += qu.sz[j]
	}
	count--
}

func (qu *WeightedQuickUnoinUF) Connected(p int, q int) {
	return root(p) == root(q)
}

func (qu *WeightedQuickUnoinUF) Find(p int) int {
	return qu.root(p)
}

func (qu *WeightedQuickUnoinUF) Count() int {
	return qu.count
}
