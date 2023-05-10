package main

func isValid(s string) bool {
	var stack []rune
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		if c == ')' && stack[len(stack)-1] != '(' {
			return false
		}
		if c == ']' && stack[len(stack)-1] != '[' {
			return false
		}
		if c == '}' && stack[len(stack)-1] != '{' {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}

func main() {
	println(isValid("()[]{}"))
	println(isValid("([)]"))
}
