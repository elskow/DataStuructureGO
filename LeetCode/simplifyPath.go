package main

func simplifyPath(path string) string {
	var stack []string
	for i := 0; i < len(path); i++ {
		if path[i] == '/' {
			continue
		}
		var tmp string
		for i < len(path) && path[i] != '/' {
			tmp += string(path[i])
			i++
		}
		if tmp == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if tmp != "." && tmp != "" {
			stack = append(stack, tmp)
		}
	}
	if len(stack) == 0 {
		return "/"
	}
	var res string
	for i := 0; i < len(stack); i++ {
		res += "/" + stack[i]
	}
	return res
}

func main() {
	path := "/home/"
	println(simplifyPath(path))
	path = "/../"
	println(simplifyPath(path))
	path = "/home//foo/"
	println(simplifyPath(path))
}
