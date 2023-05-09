package main

func intToRoman(x int) string {
	var result string
	var roman = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var value = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	for i := 0; i < len(value); i++ {
		for x >= value[i] {
			x -= value[i]
			result += roman[i]
		}
	}
	return result
}

func main() {
	println(intToRoman(1994))
}
