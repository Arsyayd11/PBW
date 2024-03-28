package main

import (
	"fmt"
	"progo/anonymous"
)

func main() {
	jariJari := 5.0
    luas, keliling := anonymous.HitungLuasKeliling(jariJari)
 
    // Menampilkan hasil 
    fmt.Println("Luas lingkaran:", luas) 
    fmt.Println("Keliling lingkaran:", keliling)
}