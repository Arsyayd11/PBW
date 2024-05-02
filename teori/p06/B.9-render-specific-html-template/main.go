package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	// Handler untuk route "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Parsing file template view.html dengan nama template "index"
		var tmpl = template.Must(template.New("index").ParseFiles("view.html"))
		// Eksekusi template tanpa data dan menulis hasilnya ke response writer
		if err := tmpl.Execute(w, nil); err != nil {
			// Menampilkan pesan error jika terjadi kesalahan saat mengeksekusi template
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Handler untuk route "/test"
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		// Parsing file template view.html dengan nama template "test"
		var tmpl = template.Must(template.New("test").ParseFiles("view.html"))
		// Eksekusi template tanpa data dan menulis hasilnya ke response writer
		if err := tmpl.Execute(w, nil); err != nil {
			// Menampilkan pesan error jika terjadi kesalahan saat mengeksekusi template
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Menampilkan pesan bahwa server telah dimulai
	fmt.Println("server started at localhost:9000")

	// Mengaktifkan server HTTP dan mendengarkan permintaan pada port 9000
	http.ListenAndServe(":9000", nil)
}
