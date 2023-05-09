package main

func myAtoi(s string) int {
	var result int
	var sign int = 1
	var i int
	for i = 0; i < len(s); i++ {
		if s[i] != ' ' {
			break
		}
	}
	if i == len(s) {
		return 0
	}
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		result = result*10 + int(s[i]-'0')
		if result > 2147483647 {
			if sign == 1 {
				return 2147483647
			} else {
				return -2147483648
			}
		}
	}
	return sign * result
}

func main() {
	println(myAtoi("42"))
}
