package main

import "fmt"

func main() {
	fmt.Println(isHappy(7))
}

func isHappy(n int) bool {
	m := map[int]bool{}
	//for ; n != 1 && !m[n]; n, m[n] = step(n), true {
	//	//
	//	//}
	for n != 1 && !m[n] {
		n = step(n)
		m[n] = true
	}
	//for ;; {}
	//for cond {}
	//for {}
	return n == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n%10) * (n%10)
		n = n/10
	}
	return sum
}
