package main

import (
	"fmt"
	"net/http"
	"time"
)

// handlerIndex adalah fungsi untuk menangani permintaan HTTP pada route "/"
func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Welcome"  // Pesan yang akan ditampilkan sebagai respon
	w.Write([]byte(message)) // Mengirimkan pesan sebagai respons ke client
}

// handlerHello adalah fungsi untuk menangani permintaan HTTP pada route "/hello"
func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello world!" // Pesan yang akan ditampilkan sebagai respon
	w.Write([]byte(message))     // Mengirimkan pesan sebagai respons ke client
}

func main() {
	// Menetapkan handler untuk masing-masing route
	http.HandleFunc("/", handlerIndex)      // Menangani permintaan pada route "/"
	http.HandleFunc("/index", handlerIndex) // Menangani permintaan pada route "/index"
	http.HandleFunc("/hello", handlerHello) // Menangani permintaan pada route "/hello"

	var address = "localhost:9000"                // Alamat server yang akan digunakan
	fmt.Printf("server started at %s\n", address) // Menampilkan pesan bahwa server telah dimulai

	server := new(http.Server)             // Membuat instance baru dari server HTTP
	server.Addr = address                  // Menetapkan alamat server
	server.ReadTimeout = time.Second * 10  // Menetapkan waktu maksimum untuk membaca request dari client
	server.WriteTimeout = time.Second * 10 // Menetapkan waktu maksimum untuk menulis response ke client

	err := server.ListenAndServe() // Mendengarkan permintaan pada alamat server yang ditetapkan

	// Menampilkan pesan error jika terjadi kesalahan saat menjalankan server
	if err != nil {
		fmt.Println(err.Error())
	}
}
