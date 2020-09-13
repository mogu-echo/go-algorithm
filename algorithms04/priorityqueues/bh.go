package priorityqueues

type BinaryHeapMaxPQ struct {
	pq []Key
	n  int
}

func NewBinaryHeapMaxPQ(capacity) *BinaryHeapMaxPQ {
	pq := make([]Compareble, capacity+1)
	return &BinaryHeapMaxPQ{pq}
}

func (b *BinaryHeapMaxPQ) sort() {
	N := b.n
	for k := 0; k < N/2; k++ {
		sink(k)
	}
	for N > 1 {
		b.pq.Swap(1, N)
		N--
		sink(1)
	}
}

func (b *BinaryHeapMaxPQ) Insert(k Key) {
	b.n++
	b.pq[b.n] = x
	swim(b.n)
}

func (b *BinaryHeapMaxPQ) DelMax() Key {
	max = b.pq[1]
	b.pq.Swap(1, b.n-1)
	b.n--
	b.sink(1)
	b.pq[b.n+1] = nil
	return
}

func (b *BinaryHeapMaxPQ) swim(k int) {
	for k > 1 && b.less(k/2, k) {
		b.pq.Swap(k, k/2)
		k /= 2
	}
}

func (b *BinaryHeapMaxPQ) sink(k int) {
	for 2*k <= b.n {
		j := 2 * k
		if j < b.n && b.less(j, j+1) {
			j++
		}
		if !b.less(k, j) {
			break
		}
		b.pq.Swap(k, j)
		k = j
	}
}

func (b *BinaryHeapMaxPQ) IsEmpty() bool {
	return b.n == 0
}

func (b *BinaryHeapMaxPQ) less(i, j int) bool {
	return b.pq[i].compareTo(b.pq[j]) < 0
}

func (b *BinaryHeapMaxPQ) Max() Key {
	panic("not implemented") // TODO: Implement
}

func (b *BinaryHeapMaxPQ) Size() int {
	panic("not implemented") // TODO: Implement
}
