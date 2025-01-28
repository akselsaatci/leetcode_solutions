package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type Stack struct {
	items []rune
}

func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (rune, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}
func isOpeningRune(c rune) (bool, error) {
	if c == '(' || c == '[' || c == '{' {
		return true, nil
	}
	if c == ')' || c == ']' || c == '}' {
		return false, nil
	}
	return false, fmt.Errorf("Char is not (,{,[,),},]. Current Char : %c", c)
}

func matchClosingRune(c rune) rune {
	switch c {
	case '}':
		return '{'
	case ')':
		return '('
	case ']':
		return '['
	default:
		// Should we throw
		return ' '
	}
}

func isValid(s string) bool {
	stack := Stack{}

	for _, v := range s {

		cond, err := isOpeningRune(v)
		if err != nil {
			return false
		}
		if cond {
			stack.Push(v)
		} else {
			c, ok := stack.Pop()
			if !ok {
				return false
			}
			if c != matchClosingRune(v) {
				return false
			}

		}
	}

	return true
}
