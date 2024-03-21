package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Membuat scanner untuk input dari konsol
	scanner := bufio.NewScanner(os.Stdin)

	// Input nama
	fmt.Println()
	fmt.Print("Masukkan nama Anda: ")
	scanner.Scan()
	nama := scanner.Text()

	// Input usia
	fmt.Print("Masukkan usia Anda: ")
	scanner.Scan()
	var usia int
	fmt.Sscanf(scanner.Text(), "%d", &usia)

	// Menentukan kategori usia
	var kategoriUsia string
	if usia < 18 {
		kategoriUsia = "Anak-anak"
	} else {
		kategoriUsia = "Dewasa"
	}

	fmt.Println()
	fmt.Printf("Selamat datang, %s! Anda termasuk kategori %s.\n", nama, kategoriUsia)
	fmt.Println()
}
