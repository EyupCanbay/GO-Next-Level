package main

import (
	"errors"
	"fmt"
)

/*
// last in first out

type Stack struct {
	cap   uint64
	depth uint64
	list  []int
}

func NewStack(cap uint64) *Stack {
	return &Stack{
		cap:   cap,
		depth: 0,
		list:  make([]int, 0, cap),
	}
}

func (s *Stack) Push(v int) {
	s.list[s.depth] = v
	s.depth++
}

func (s *Stack) Pop() int {
	if s.depth <= 0 {
		return -1
	}
	s.depth--
	return s.list[s.depth+1]
}

func main() {
	s := NewStack(5)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	println(s.Pop)
}
*/

type Stack struct {
	cap   uint64
	depth uint64
	list  []int
}

func NewStack(cap uint64) *Stack {
	return &Stack{
		cap:   cap,
		depth: 0,
		list:  make([]int, 0, cap),
	}
}

func (s *Stack) Push(v int) error {
	if s.depth >= s.cap {
		return errors.New("stack overflow: stack is full")
	}
	s.list = append(s.list, v)
	s.depth++
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.depth == 0 {
		return 0, errors.New("stack underflow: stack is empty")
	}
	s.depth--
	value := s.list[s.depth]
	s.list = s.list[:s.depth]
	return value, nil
}

func (s *Stack) Peek() (int, error) {
	if s.depth == 0 {
		return 0, errors.New("stack is empty")
	}
	return s.list[s.depth-1], nil
}

func (s *Stack) IsEmpty() bool {
	return s.depth == 0
}

func (s *Stack) Size() uint64 {
	return s.depth
}

func main() {
	s := NewStack(5)

	if err := s.Push(1); err != nil {
		fmt.Println(err)
	}
	if err := s.Push(2); err != nil {
		fmt.Println(err)
	}
	if err := s.Push(3); err != nil {
		fmt.Println(err)
	}
	if err := s.Push(4); err != nil {
		fmt.Println(err)
	}
	if err := s.Push(5); err != nil {
		fmt.Println(err)
	}
	if err := s.Push(6); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Yığın boyutu:", s.Size()) // 5

	val, err := s.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Pop edilen değer:", val) // 5
	}

	val, err = s.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Pop edilen değer:", val) // 4
	}

	fmt.Println("Yığın boyutu:", s.Size()) // 3

	val, err = s.Peek()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("En üstteki eleman (Peek):", val) // 3
	}

	for !s.IsEmpty() {
		val, err := s.Pop()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Pop edilen değer:", val)
		}
	}

	val, err = s.Pop()
	if err != nil {
		fmt.Println(err) // stack underflow
	} else {
		fmt.Println("Pop edilen değer:", val)
	}

	fmt.Println("Yığın boş mu?", s.IsEmpty()) // true
}
