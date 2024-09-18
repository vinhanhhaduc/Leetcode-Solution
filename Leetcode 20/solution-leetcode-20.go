type Stack struct {
    items []rune
}

func (s *Stack) Push(item rune) {
    s.items = append(s.items, item)
}

func (s *Stack) Pop() rune {
    if len(s.items) == 0 {
        fmt.Println("Stack is empty")
        return 0
    }
    lastIndex := len(s.items) - 1
    poppedItem := s.items[lastIndex]
    s.items = s.items[:lastIndex]
    return poppedItem
}

func (s *Stack) Peek() rune {
    if len(s.items) == 0 {
        fmt.Println("Stack is empty")
        return 0
    }
    return s.items[len(s.items)-1]
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
    return len(s.items) == 0
}

func isValid(s string) bool {
    stack := Stack{}
    
    for _, char := range s {
        if stack.IsEmpty() {
            stack.Push(char)
        } else if (stack.Peek() == '(' && char == ')') || (stack.Peek() == '{' && char == '}') || (stack.Peek() == '[' && char == ']') {
            stack.Pop()
        } else {
            stack.Push(char)
        }
    }
    
    return stack.IsEmpty()
}