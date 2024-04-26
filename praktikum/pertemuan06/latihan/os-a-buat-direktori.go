package main

import (
	"fmt"
	"os"
)

func main() {
	// Membuat direktori baru.
	err := os.Mkdir("ArsyaYanDuribta", 0755)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Direktori 'ArsyaYanDuribta' berhasil dibuat.")
}
