package sort

func Quick(a []Compareble) {
	sort(a, 0, len(a)-1)
}

func sort(a []Compareble, lo, hi int) {
	if hi <= lo {
		return
	}
	j := pairtition(a, lo, hi)
	sort(a, lo, j-1)
	sort(a, j+1, hi)
}

func sortK(a []Compareble, k int) Compareble {
	lo, hi := 0, len(a)-1
	for hi > lo {
		j := partition(a, lo, hi)
		if j < k {
			lo = j + 1
		} else if j > k {
			hi = j - 1
		} else {
			return a[k]
		}
	}
	return a[k]
}

func sortDuplicateKeys(a []Compareble, lo, hi int) {
	if hi <= lo {
		return
	}
	lt, gt := lo, hi
	Compareble v = a[lo]
	i := lo
	for i <= gt {
		cmp := a[i].compareTo(v)
		if cmp < 0 {
			a.Swap(lt, i)
			lt++
			i++
		} else if cmp > 0 {
			a.Swap(i, gt)
			gt--
		} else {
			i++
		}
	}
	sortDuplicateKeys(a, lo, lt-1)
	sortDuplicateKeys(a, gt+1, gi)
}

func partition(a []Compareble, lo, hi int) int {
	i, j := lo, hi+1
	for {
		for ; a.Less(i+1, lo); i++ {
			if i == hi {
				break
			}
		}
		for ; a.Less(lo, j-1); j-- {
			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		a.Swap(i, j)
	}
}
