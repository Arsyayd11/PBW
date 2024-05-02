package main

import (
	"fmt"
	"net/http"
)

func main() {
	// handlerIndex adalah fungsi anonim yang menangani permintaan HTTP pada route "/" dan "/index"
	handlerIndex := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	}

	// Menetapkan handlerIndex sebagai handler untuk route "/" dan "/index"
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)

	// Menetapkan fungsi anonim sebagai handler untuk route "/data"
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello again"))
	})

	// Menampilkan pesan bahwa server telah dimulai
	fmt.Println("server started at localhost:9000")

	// Mengaktifkan server HTTP dan mendengarkan permintaan pada port 9000
	http.ListenAndServe(":9000", nil)
}
