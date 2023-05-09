package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var reverse int
	var temp int = x
	for temp != 0 {
		reverse = reverse*10 + temp%10
		temp /= 10
	}
	return reverse == x
}

func main() {
	println(isPalindrome(121))
	println(isPalindrome(-121))
}
