package priorityqueues

type Key = Compareble

type MaxPQ interface {
	Insert(k Key)
	DelMax() Key
	IsEmpty() bool
	Max() Key
	Size() int
}

type UnorderedMaxPQ struct {
	pq []Key
	n  int
}

func NewUnorderedMaxPQ(capacity int) *UnorderedMaxPQ {
	pq := make([]Compareble, capacity)
	return &UnorderedMaxPQ{pq}
}

func (u *UnorderedMaxPQ) Insert(k Key) {
	u.pq[n] = k
	u.n++
}

func (u *UnorderedMaxPQ) DelMax() Key {
	max := 0
	for i := 0; i < u.n; i++ {
		if u.pq.Less(max, i) {
			max = i
		}
	}
	u.pq.Swap(max, u.n-1)
	k := u.pq[u.n-1]
	u.n--
	return k
}

func (u *UnorderedMaxPQ) IsEmpty() bool {
	return u.n == 0
}

func (u *UnorderedMaxPQ) Max() Key {
	panic("not implemented") // TODO: Implement
}

func (u *UnorderedMaxPQ) Size() int {
	panic("not implemented") // TODO: Implement
}
