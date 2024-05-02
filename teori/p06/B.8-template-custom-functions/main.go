package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// funcMap adalah peta fungsi yang akan digunakan dalam template HTML
var funcMap = template.FuncMap{
	// unescape adalah fungsi untuk mengonversi string ke template.HTML, menghindari escape HTML
	"unescape": func(s string) template.HTML {
		return template.HTML(s)
	},
	// avg adalah fungsi untuk menghitung rata-rata dari beberapa angka
	"avg": func(n ...int) int {
		var total = 0
		for _, each := range n {
			total += each
		}
		return total / len(n)
	},
}

func main() {
	// Handler untuk route "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Parsing file template view.html dengan menggunakan funcMap
		var tmpl = template.Must(template.New("view.html").Funcs(funcMap).ParseFiles("view.html"))
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
