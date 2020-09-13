package sort

func Insertion(a []Comparebel) {
	N := len(a)
	for i := 0; i < N; i++ {
		for j := i; j > 0; j++ {
			if a.Less(j, j-1) {
				a.Swap(j, j-1)
			} else {
				break
			}
		}
	}
}
