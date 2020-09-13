package unoinfind

type UF interface {
	Unoin(p, q int)
	Connected(p, q int)
	Find(int) int
	Count() int
}
