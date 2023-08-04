package main

func factorial(n int) int {
	res := 1

	for n > 0 {
		res *= n
		n -= 1
	}

	return res
}
