package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func main() {
	// Menetapkan handler untuk route "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Menggabungkan path ke file template index.html dalam folder "views"
		var filepath = path.Join("views", "index.html")
		// Membaca dan mem-parsing file template
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			// Menampilkan pesan error jika terjadi kesalahan saat mem-parsing template
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Data yang akan dipass ke template
		var data = map[string]interface{}{
			"title": "Learning Golang Web", // Judul halaman
			"name":  "Batman",              // Nama yang akan ditampilkan di halaman
		}

		// Mengeksekusi template dengan data yang diberikan dan menulis hasilnya ke response writer
		err = tmpl.Execute(w, data)
		if err != nil {
			// Menampilkan pesan error jika terjadi kesalahan saat mengeksekusi template
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Menetapkan handler untuk serve file statis dari folder "assets"
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	// Menampilkan pesan bahwa server telah dimulai
	fmt.Println("server started at localhost:9000")

	// Mengaktifkan server HTTP dan mendengarkan permintaan pada port 9000
	http.ListenAndServe(":9000", nil)
}
