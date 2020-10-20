// Package st provides Hash Tables
package st

type Node struct {
	K    Key
	V    Value
	Next *Node
}

type SeperateChainingHashST struct {
	M  int
	st []Node
}

func (st *SeperateChainingHashSt) Get(k Key) Value {
	i := hash(k)
	for x := st.st[i]; x != nil; x = x.Next {
		if k.Equals(x.K) {
			return x.V
		}
	}
	return nil
}

func (st *SeperateChainingHashst) Put(k Key, v Value) {
	i := hash(k)
	for x := st.st[i]; x != nil; x = x.Next {
		if k.Equals(x.K) {
			x.V = v
			return
		}
	}
	st.st[i] = &Node{k, v, st.st[i]}
	return
}

type LinearProbingHashST struct {
	M    int
	Vals []Value
	Keys []Key
}

func (st *LinearProbingHashST) Put(k Key, v Value) {
	var i int
	for i = hash[k]; st.Keys[i] != nil; i = (i + 1) % st.M {
		if k.Equals(i) {
			break
		}
	}
	st.Keys[i] = k
	st.Vals[i] = v
}

func (st *LinearProbingHashST) Get(k Key) Value {
	for i := hash(k); st.Keys[i] != nil; i = (i + 1) % st.M {
		if k.Equals(i) {
			return st.Vals[i]
		}
	}
	return nil
}
