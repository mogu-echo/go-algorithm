package sort

func Shell(a []Compareble) {
	N := len(a)
	h := 1
	for h < N/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < N; i++ {
			for j := i; j >= h && a.Less(j, j-h); j -= h {
				a.Swap(i, i-h)
			}
		}
		h /= 3
	}
}
