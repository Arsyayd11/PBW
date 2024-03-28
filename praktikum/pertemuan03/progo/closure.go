package main

import (
    "fmt"
    "progo/closure"
)

func main() {
    // Menentukan nilai huruf untuk beberapa nilai ujian 
    nilaiUjian := []int{85, 75, 95, 55, 65}
    for _, nilai := range nilaiUjian {
        fmt.Println("Nilai ujian", nilai, "=", closure.GetNilaiHuruf(nilai))
    }
}
