package bst

// A BST such that:
// No node has two red links connected to it.
// Every path from root to null link has the same number of black links. Red links lean left.

const (
	RED   = true
	BLACK = false
)

type Node struct {
	K           Key
	V           Value
	N           int
	Color       bool
	Left, Right *Node
}

func isRed(x *Node) bool {
	if x == nil {
		return false
	}
	return x.Color == RED
}

func rotateLeft(h *Node) *Node {
	x := h.Right
	h.Right = x.Left
	x.Left = h
	x.Color = h.Color
	h.Color = RED
	x.N = h.N
	h.N = 1 + size(h.Left) + size(h.Right)
	return x
}

func rotateRight(h *Node) *Node {
	x := h.Left
	h.Left = x.Right
	x.Right = h
	x.Color = h.Color
	h.Color = RED
	x.N = h.N
	h.N = 1 + size(h.Left) + size(h.Right)
	reutn x
}

func flipColors(h *Node) {
	h.Color = RED
	h.Left.Color = BLACK
	h.Right.Color = BLACK
}
 
type RedBlackBST struct {
	root *Node
}

func (r *RedBlackBST) Size() {
	return size(r.root)
}

func size(x Node) int {
	if x == nil {
		return 0
	}
	return x.Count
}

fucn (r *RedBlackBST) Get(k Key) Value {
	x := r.root
	for x != nil {
		cmp := k.compareTo(x.K)
		if cmp < 0 {
			x = x.Left
		} else if cmp > 0 {
			x = x.Right
		} else {
			return x.V
		}
	}
	return nil
}

func (r *RedBlackBST) Put(k Key, v Value) {
	r.root = put(r.root, k, v)
	r.root.Color = Block
}

func put(h *Node, k Key, v Value) *Node {
	if h == nil {
		return &Node{
			k,v,1,RED,
		}
	}
	cmp := k.compareTo(h.Key)
	if cmp < 0 {
		h.Left = put(h.Left, k, v)
	} else if cmp > 0 {
		h.Right = put(h.Right, k, v)
	} else {
		h.V = v
	}
	if isRed(h.Right) && !isRed(h.Left) {
	   h = rotateLeft(h)
	}
	if isRed(h.Left) && isRed(h.Left.Left) {
		h = rotateRight(h)
	}
	if isRed(h.Left) && isRed(h.Right) {
		flipColors(h)
	}
	h.N = 1 + size(h.Left) + size(h.Right)
	return h
}

func (r *RedBlackBST) Size(lo, hi Key) {
	if r.Contains(hi) {
		return r.Rank(hi) - r.Rank(lo) + 1
	} else {
		return r.Rank(hi) - r.Rank(lo)
	}
}
