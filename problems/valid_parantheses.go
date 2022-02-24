package problems

import "errors"

var stack Stack

type Stack struct {
	items []string
}

func InitStack() {
	stack = Stack{}
}

func Push(item string) {
	stack.items = append(stack.items, item)
}

func Pop() (string, error) {

	if len(stack.items) == 0 {
		return "", errors.New("Stack empty")
	}
	to_return := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]
	return to_return, nil
}

func Peak() (string, error) {
	if len(stack.items) == 0 {
		return "", errors.New("Stack empty")
	}
	return stack.items[len(stack.items)-1], nil
}

func Size() int {
	return len(stack.items)
}

func ValidateParantheses(item string) bool {
	for _, char := range item {
		if Size() == 0 && (string(char) == "]" || string(char) == "}" || string(char) == ")") {
			return false
		}
		if string(char) == "[" || string(char) == "{" || string(char) == "(" {
			Push(string(char))
		} else if (string(char)) == "}" {
			if compare, err := Pop(); err == nil && compare != "{" {
				return false
			}
		} else if (string(char)) == "]" {
			if compare, err := Pop(); err == nil && compare != "[" {
				return false
			}
		} else if (string(char)) == ")" {
			if compare, err := Pop(); err == nil && compare != "(" {
				return false
			}
		}
	}
	return Size() == 0
}
