package sort

func Selection(a []Compareble) {
	N := len(a)
	for i := 0; i < N; i++ {
		min := i
		for j := i + 1; j < N; j++ {
			if a.Less(a[j], a[min]) {
				min = j
			}
		}
		a.Swap(i, min)
	}
}
