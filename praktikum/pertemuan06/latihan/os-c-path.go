package main

import (
	"fmt"
	"os"
)

func main() {
	var err error
	fileInfo, err := os.Stat("ArsyaYanDuribta")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if fileInfo.IsDir() {
		fmt.Println("ArsyaYanDuribta adalah sebuah direktori.")
	} else {
		fmt.Println("ArsyaYanDuribta adalah sebuah file.")
	}
}
