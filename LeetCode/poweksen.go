package main

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		return myPow(x*x, n/2)
	}
	return x * myPow(x, n-1)
}

func main() {
	println(myPow(2.0, 10))
	println(myPow(2.1, 3))
	println(myPow(2.0, -2))
}
