package main

import "fmt"

// Mengurutkan array dalam satu kali perulangan
func main() {
	fmt.Println("Bubble Short")
	var arrayNumber = [20]int{4, 2, 10, 5, 8, 1, 6, 3, 9, 7, 12, 11, 19, 15, 17, 14, 18, 13, 16, 20}

	fmt.Println("Sebelum dilakukan pengurutan:")
	fmt.Println(arrayNumber)

	i := 0
	for i < len(arrayNumber) {
		if i != (len(arrayNumber) - 1) {
			if arrayNumber[i] > arrayNumber[i+1] {
				x := arrayNumber[i]
				arrayNumber[i] = arrayNumber[i+1]
				arrayNumber[i+1] = x
				if i > 0 {
					i--
				}
			} else {
				i++
			}
		} else {
			if arrayNumber[i] < arrayNumber[i-1] {
				x := arrayNumber[i]
				arrayNumber[i] = arrayNumber[i-1]
				arrayNumber[i-1] = x
				i--
			} else {
				break
			}
		}
	}

	fmt.Println("\nSetelah dilakukan pengurutan.")
	fmt.Println(arrayNumber)
}
