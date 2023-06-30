package main

func fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	} else if n == 1 {
		return []int{0}
	} else if n == 2 {
		return []int{0, 1}
	}

	fibSequence := make([]int, n)
	fibSequence[0] = 0
	fibSequence[1] = 1

	for i := 2; i < n; i++ {
		fibSequence[i] = fibSequence[i-1] + fibSequence[i-2]
	}

	return fibSequence
}

func main() {
	n := 1000000
	fibonacci(n)
}
