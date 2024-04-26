package main

import (
	"fmt"
	"os"
)

func main() {
	var err error
	err = os.Chmod("ArsyaYanDuribta", 0117)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Izin 'ArsyaYanDuribta' telah diubah menjadi 0117.")
}
