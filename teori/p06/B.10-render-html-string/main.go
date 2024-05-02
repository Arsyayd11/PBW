package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// view adalah string yang berisi kode HTML dari template
const view string = `<!DOCTYPE html>
<html>
	<head>
		<title>Template</title>
	</head>
	<body>
		<h1>Hello</h1>
	</body>
</html>`

func main() {
	// Handler untuk route "/index"
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		// Parsing template dari string view
		var tmpl = template.Must(template.New("main-template").Parse(view))
		// Eksekusi template tanpa data dan menulis hasilnya ke response writer
		if err := tmpl.Execute(w, nil); err != nil {
			// Menampilkan pesan error jika terjadi kesalahan saat mengeksekusi template
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Handler untuk route "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Redirect permintaan ke route "/index"
		http.Redirect(w, r, "/index", http.StatusTemporaryRedirect)
	})

	// Menampilkan pesan bahwa server telah dimulai
	fmt.Println("server started at localhost:9000")

	// Mengaktifkan server HTTP dan mendengarkan permintaan pada port 9000
	http.ListenAndServe(":9000", nil)
}
