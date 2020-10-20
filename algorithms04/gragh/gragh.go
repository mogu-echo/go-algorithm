// Package gragh provides Gragh Interface
package gragh

type Gragh interface {
	AddEdge(v, w int)
	Adj(int)
	V() int
	E() int
	ToString() string
}

func Degree(G Gragh, v int) int {
	degree := 0
	for _ := range G.Adj(v) {
		degree++
	}
	return degree
}

func MaxDegree(G Gragh) int {
	max := 0
	for i, v := range G.V() {
		if Degree(G, v) > max {
			max = Degree(G, v)
		}
	}
	return max
}

func AvgDegree(G Gragh) float64 {
	return 2.0 * G.E() / G.V()
}

func NumberOfSelfLoops(G Gragh) int {
	count := 0
	for v := range G.V() {
		for w := range G.Adj(v) {
			if v == w {
				count++
			}
		}
	}
	return count / 2
}
