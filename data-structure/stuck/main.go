package main

import "fmt"

//Stack safsfsfd
type Stack struct {
	items []int
}

// Push stack
func (s *Stack) push(i int) {
	s.items = append(s.items, i)
}

func (s Stack) printStackData() {
	for _, item := range s.items {
		fmt.Printf("%d ", item)
	}
	fmt.Printf("\n")
}

func (s *Stack) pop() int {
	lastIndex := len(s.items) - 1
	toRemove := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return toRemove
}

func main() {
	stack := Stack{}
	fmt.Println(stack)
	stack.push(100)
	stack.push(200)
	stack.push(300)
	stack.printStackData()
	//fmt.Println(myStack)
	stack.pop()
	//fmt.Println(myStack)
}
