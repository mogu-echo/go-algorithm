package problem0438

func findAnagrams(s string, p string) []int {
	var res = []int{}

	var target, window [26]int
	for i := 0; i < len(p); i++ {
		target[p[i]-'a']++
	}

	check := func(i int) {
		if window == target {
			res = append(res, i)
		}
	}

	for i := 0; i < len(s); i++ {
		window[s[i]-'a']++
		if i == len(p)-1 {
			check(0)
		} else if len(p) <= i {
			window[s[i-len(p)]-'a']--
			check(i - len(p) + 1)
		}
	}

	return res
}

func findAnagrams2(s string, p string) []int {
	win, need := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	left, right := 0, 0
	match := 0
	ans := make([]int, 0)
	for right < len(s) {
		c := s[right]
		right++
		if need[c] != 0 {
			win[c]++
			if win[c] == need[c] {
				match++
			}
		}

		for right-left >= len(p) {
			if right-left == len(p) && match == len(need) {
				ans = append(ans, left)
			}
			d := s[left]
			left++
			if need[d] != 0 {
				if win[d] == need[d] {
					match--
				}
				win[d]--
			}
		}
	}
	return ans
}
