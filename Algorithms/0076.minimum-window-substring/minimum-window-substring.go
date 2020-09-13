package problem0076

import "math"

func minWindow(s string, t string) string {
	have := [128]int{}
	need := [128]int{}
	for i := range t {
		need[t[i]]++
	}

	size, total := len(s), len(t)

	min := size + 1
	res := ""

	// s[i:j+1] 就是 window
	// count 用于统计已有的 t 中字母的数量。
	// count == total 表示已经收集完需要的全部字母
	for i, j, count := 0, 0, 0; j < size; j++ {
		if have[s[j]] < need[s[j]] {
			// 出现了 window 中缺失的字母
			count++
		}
		have[s[j]]++

		// 保证 window 不丢失所需字母的前提下
		// 让 i 尽可能的大
		for i <= j && have[s[i]] > need[s[i]] {
			have[s[i]]--
			i++
		}

		width := j - i + 1
		if count == total && min > width {
			min = width
			res = s[i : j+1]
		}

	}

	return res
}

func minWindow2(s string, t string) string {
	win := make(map[byte]int)
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	left, right := 0, 0
	start, end := 0, 0
	match, min := 0, math.MaxInt64
	var c byte
	for right < len(s) {
		c = s[right]
		right++
		if need[c] != 0 {
			win[c]++
			if win[c] == need[c] {
				match++
			}
		}
		for match == len(need) {
			if right-left < min {
				min = right - left
				start = left
				end = right
			}
			c = s[left]
			left++
			if need[c] != 0 {
				if win[c] == need[c] {
					match--
				}
			}
			win[c]--
		}
	}
	if min == math.MaxInt64 {
		return ""
	}
	return s[start:end]
}
