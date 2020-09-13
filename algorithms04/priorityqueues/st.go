package priorityqueues

type ST interface {
	Put(Key, Value)
	Get(Key) Value
	Delete(Key)
	Contains(Key) bool
	IsEmpty() bool
	Min() Key
	Max() Key
	Floor(Key) Key
	Ceiling(Key) Key
	Rank(int) int
	Select(int) Key
	DeleteMin()
	DeleteMax()
	Size() int
	Contains(Key) bool
	KeysRange(lo, hi Key) Iterable
	Keys() Iterable
}

type Node struct {
	Key         Key
	Val         Value
	Count       int
	Left, Right *Node
}

type BST struct {
	root *Node
}

func NewBST() *BST {
	return &BST{}
}

func (b *BST) Put(k Key, v Value) {
	b.root = b.put(k, v)
}

func put(x Node, k Key, val Value) *Node {
	if x == nil {
		return &Node{k, val, 1}
	}
	cmp := k.compareTo(x.Key)
	if cmp < 0 {
		x.Left = put(x.Left, key, val)
	} else if cmp > 0 {
		x.Right = put(x.Right, key, val)
	} else {
		x.Val = val
	}
	x.Count = 1 + size(x.Left) + size(x.Right)
	return x
}

func (b *BST) Size() int {
	return size(b.root)
}

func size(x Node) int {
	if x == nil {
		return 0
	}
	return x.Count
}

func (b *BST) Get(k Key) Value {
	x := b.root
	for x != nil {
		cmp := k.compareTo(x.Key)
		if cmp < 0 {
			x = x.Left
		} else if cmp > 0 {
			x = x.Right
		} else {
			return x.Val
		}
	}
	return nil
}

func (b *BST) Floor(k Key) Key {
	x := floor(b.root, k)
	if x == nil {
		return nil
	}
	return x.Key
}

func floor(x Node, k Key) *Node {
	if x == nil {
		return nil
	}
	cmp := k.compareTo(x.Key)
	if cmp == 0 {
		return x
	}
	if cmp < 0 {
		return floor(x.Left, Key)
	}
	t := floor(t.Right, k)
	if t != nil {
		return t
	} else {
		return x
	}
}

func (b *BST) Rank(k Key) int {
	return rank(k, b.root)
}

func rank(k Key, x Node) int {
	if x == nil {
		return 0
	}
	cmp := k.compareTo(x.Key)
	if cmp < 0 {
		return rank(k, x.Left)
	} else if cmp > 0 {
		return 1 + size(x.Left) + rank(k, x.Right)
	} else {
		return size(x.Left)
	}
}

func (b *BST) Delete(k Key) {
	b.root = delete(b.root, k)	
}

func delete(x *Node, k Key) *Node {
	if x == nil {
		return nil 
	}	
	cmp := k.compareTo(x.Key)
	if cmp < 0 {
		x.Left = delete(x.right, key)
	} else cmp > 0 {
		x.Right = delete(x.Left, key)
	} else {
		if x.Right == nil {
			return x.Left
		}
		if x.Left == nil {
			return x.Right
		}
		t := x
		x = Min(t.Right)
		x.Right = deleteMin(t.Right)
		x.Left = t.Left
	}
	x.Count = 1 + size(x.Left) + size(x.Right)
	return x
}

func (b *BST) DeleteMin() {
	b.root = deleteMin(b.root)
}

func deleteMin(x *Node) *Node {
	if x.Left == nil {
		return x.Right
	}
	x.Left = deleteMin(x.Left)
	x.Count = 1 + size(x.Left) + size(x.right)
	return x
}

func (b *BST) Iterator() Iterable {
	q := NewQueue()
	inorder(b.root, q)
	return q
}

func inorder(x Node, q Queue) {
	if x == nil {
		return
	}
	inorder(x.left, q)
	q.Enquque(x.Key)
	inorder(x.Right, q)
}
