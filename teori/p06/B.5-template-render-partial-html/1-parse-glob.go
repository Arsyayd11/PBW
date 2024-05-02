package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// M adalah tipe data untuk menyimpan data yang akan dipass ke template
type M map[string]interface{}

func main() {
	// Mem-parsing semua file template dalam folder "views"
	var tmpl, err = template.ParseGlob("views/*")
	if err != nil {
		// Menampilkan pesan error jika terjadi kesalahan saat mem-parsing template
		panic(err.Error())
		return
	}

	// Menetapkan handler untuk route "/index"
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		// Data yang akan dipass ke template
		var data = M{"name": "Batman"}
		// Mengeksekusi template "index" dengan data yang diberikan dan menulis hasilnya ke response writer
		err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			// Menampilkan pesan error jika terjadi kesalahan saat mengeksekusi template
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Menetapkan handler untuk route "/about"
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		// Data yang akan dipass ke template
		var data = M{"name": "Batman"}
		// Mengeksekusi template "about" dengan data yang diberikan dan menulis hasilnya ke response writer
		err = tmpl.ExecuteTemplate(w, "about", data)
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
