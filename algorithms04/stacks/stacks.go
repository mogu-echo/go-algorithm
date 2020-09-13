package main

import (
	"bufio"
	"fmt"
	"os"
)

type ListNode struct {
	Val  string
	Next *ListNode
}

type StackOfStrings interface {
	Push(string)
	Pop() string
	IsEmpty() bool
	Size() int
}

type Iterator interface {
	hasNext() bool
	Next()
	// Remove()
}

type LinkedStackOfStrings struct {
	first *ListNode
}

func NewLinkedStackOfStrings() *LinkedStackOfStrings {
	return &LinkedStackOfStrings{}
}

func (s *LinkedStackOfStrings) Push(c string) {
	oldNode := s.first
	s.first = new(ListNode)
	s.first.Val = c
	s.first.Next = oldNode
}

func (s *LinkedStackOfStrings) Pop() (c string) {
	c = s.first.Val
	s.first = s.first.Next
	return
}

func (s *LinkedStackOfStrings) Size() (n int) {
	return
}

func (s *LinkedStackOfStrings) IsEmpty() bool {
	return s.first == nil
}

type ListIterator struct {
	current *ListNode
}

func (iter *ListIterator) HasNext() bool {
	return iter.current != nil
}

func (iter *ListIterator) Next() (c string) {
	c = iter.current.Val
	iter.current = iter.current.Next
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var s StackOfStrings = NewLinkedStackOfStrings()
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) == 0 {
			break
		}
		if text == "-" {
			fmt.Println(s.Pop())
		} else {
			s.Push(text)
		}
	}
}
