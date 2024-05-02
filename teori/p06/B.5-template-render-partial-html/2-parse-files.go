package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// M adalah tipe data untuk menyimpan data yang akan dipass ke template
type M map[string]interface{}

func main() {
	// Menetapkan handler untuk route "/index"
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		// Data yang akan dipass ke template
		var data = M{"name": "Batman"}
		// Mem-parsing file template index.html bersama dengan file-file partial _header.html dan _message.html
		var tmpl = template.Must(template.ParseFiles(
			"views/index.html",
			"views/_header.html",
			"views/_message.html",
		))
		// Mengeksekusi template "index" dengan data yang diberikan dan menulis hasilnya ke response writer
		var err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			// Menampilkan pesan error jika terjadi kesalahan saat mengeksekusi template
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Menetapkan handler untuk route "/about"
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		// Data yang akan dipass ke template
		var data = M{"name": "Batman"}
		// Mem-parsing file template about.html bersama dengan file-file partial _header.html dan _message.html
		var tmpl = template.Must(template.ParseFiles(
			"views/about.html",
			"views/_header.html",
			"views/_message.html",
		))
		// Mengeksekusi template "about" dengan data yang diberikan dan menulis hasilnya ke response writer
		var err = tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			// Menampilkan pesan error jika terjadi kesalahan saat mengeksekusi template
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Menampilkan pesan bahwa server telah dimulai
	fmt.Println("server started at localhost:9000")

	// Mengaktifkan server HTTP dan mendengarkan permintaan pada port 9000
	http.ListenAndServe(":9000", nil)
}
