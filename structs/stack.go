package structs

// Stack abstraction of a stack
type Stack struct {

	stack []int
}

// Push an int item into the stack
func (s *Stack) Push(item int){
	s.stack = append(s.stack, item)
}

// Pop and return the top item in the stack
func (s* Stack) Pop() int {

	lastIndex := len(s.stack) - 1
	item := s.stack[lastIndex]

	s.stack = s.stack[:lastIndex]
	return item
}