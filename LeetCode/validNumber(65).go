package main

func isNumber(s string) bool {
	var i int
	for i < len(s) && s[i] == ' ' {
		i++
	}
	if i < len(s) && (s[i] == '+' || s[i] == '-') {
		i++
	}
	var isNum, isDot, isE bool
	for i < len(s) {
		if s[i] >= '0' && s[i] <= '9' {
			isNum = true
		} else if s[i] == '.' && !isDot && !isE {
			isDot = true
		} else if (s[i] == 'e' || s[i] == 'E') && !isE && isNum {
			isE = true
			isNum = false
		} else if s[i] == '+' || s[i] == '-' {
			if s[i-1] != 'e' && s[i-1] != 'E' {
				return false
			}
		} else {
			break
		}
		i++
	}
	for i < len(s) && s[i] == ' ' {
		i++
	}
	return isNum && i == len(s)
}

func main() {
	s := "0"
	println(isNumber(s))
	s = " 0.1 "
	println(isNumber(s))
	s = "abc"
	println(isNumber(s))
}
