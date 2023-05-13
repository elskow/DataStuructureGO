package main

func addBinary(a string, b string) string {
	var res string
	var carry byte
	i, j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0 {
		var sum byte
		if i >= 0 {
			sum += a[i] - '0'
		}
		if j >= 0 {
			sum += b[j] - '0'
		}
		sum += carry
		res = string(sum%2+'0') + res
		carry = sum / 2
		i--
		j--
	}
	if carry == 1 {
		res = "1" + res
	}
	return res
}

func main() {
	a := "11"
	b := "1"
	println(addBinary(a, b))
	a = "1010"
	b = "1011"
	println(addBinary(a, b))
}
