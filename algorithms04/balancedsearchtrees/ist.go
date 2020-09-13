package bst

type IntervalST interface {
	Put(lo, hi Key, val Value)
	Get(Key) Value
	Delete(lo, hi Key)
	Intersects(lo, hi Key) Iterable
}

func (st *IntervalST) search(ho, hi Key) {
	x := st.root
	for x != nil {
		if x.Interval.Intersects(lo, hi) {
			return x.Interval
		} else if x.Left == nil {
			x = x.Right
		} else if x.Left.max < lo {
			x = x.Right
		} else {
			x = x.Left
		}
	}
	return nil
}
