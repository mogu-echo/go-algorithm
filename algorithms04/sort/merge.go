package sort

import "math"

const CUTOFF = 10

func Merge(a []Compareble) {
	aux := make([]Compareble, len(a))
	sort(a, aux, 0, len(a)-1)
}

func MergeBU(a []Compareble) {
	N := len(a)
	aux := make([]Compareble, N)
	for sz := 1; sz < N; sz += sz {
		for lo := 0; lo < N-sz; lo += sz + sz {
			merge(a, aux, lo, lo+sz-1, math.Min(lo+sz+sz-1, N-1))
		}
	}
}

func sort(a, aux []Compareble, lo, hi int) {
	if hi < lo {
		return
	}
	if hi <= lo+CUTOFF-1 {
		Insertion(a)
		return
	}
	mid := lo + (hi-lo)/2
	sort(a, aux, lo, mid)
	sort(a, aux, mid+1, hi)
	if !a.Less(mid+1, mid) {
		return
	}
	merge(a, aux, lo, mid, hi)
}

func merge(a, aux []Compareble, lo, mid, hi int) {
	for i := range a {
		aux[i] = a[i]
	}
	i, j := lo, hi
	for k := lo; k < hi; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else if aux.Less(j, i) {
			a[k] = aux[j]
			j++
		} else {
			a[k] = aux[i]
			i++
		}
	}
}
