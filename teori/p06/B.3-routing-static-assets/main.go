package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Menetapkan handler untuk route yang dimulai dengan "/static/"
	// Handler ini akan menangani permintaan untuk file statis di dalam folder "assets"
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	// Menampilkan pesan bahwa server telah dimulai
	fmt.Println("server started at localhost:9000")

	// Mengaktifkan server HTTP dan mendengarkan permintaan pada port 9000
	http.ListenAndServe(":9000", nil)
}
