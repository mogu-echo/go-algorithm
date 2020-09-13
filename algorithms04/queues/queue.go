package main

type ListNode struct {
	Val  string
	Next *ListNode
}

type QueueOfStrings interface {
	Enqueue(string)
	Dequeue() string
	Size() int
	IsEmpty() bool
}

type LinkedQueueOfStrings struct {
	first, last *ListNode
}

func NewLinkedQueueOfStrings() *LinkedQueueOfStrings {
	return &LinkedQueueOfStrings{}
}

func (s *LinkedQueueOfStrings) Enqueue(c string) {
	oldLast := s.last
	s.last = new(ListNode)
	s.last.Val = c
	if s.IsEmpty() {
		s.first = s.last
	} else {
		oldLast.Next = s.last
	}
}

func (s *LinkedQueueOfStrings) Dequeue() (c string) {
	c = s.first.Val
	s.first = s.first.Next
	if s.IsEmpty() {
		s.last = nil
	}
	return
}

func (s *LinkedQueueOfStrings) IsEmpty() bool {
	return s.first == nil
}

func (s *LinkedQueueOfStrings) Size() (n int) {
	return
}
