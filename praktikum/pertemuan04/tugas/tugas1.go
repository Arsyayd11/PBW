package main

import "fmt"

// Fungsi faktorial
func fact(n int) int {
	// Fungsi dibawah sama saja dengan
	// 7 x (6 x (5 x (4 x (3 x (2 x (1 x 1))))))
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}
