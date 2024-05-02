package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Superhero adalah struktur untuk merepresentasikan karakter superhero
type Superhero struct {
	Name    string   // Nama
	Alias   string   // Alias
	Friends []string // Daftar teman
}

// SayHello adalah metode untuk superhero memberi salam
func (s Superhero) SayHello(from string, message string) string {
	return fmt.Sprintf("%s said: \"%s\"", from, message) // Mengembalikan salam dari superhero
}

func main() {
	// Handler untuk route "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Data superhero
		var person = Superhero{
			Name:    "Bruce Wayne",
			Alias:   "Batman",
			Friends: []string{"Superman", "Flash", "Green Lantern"},
		}

		// Parsing file template view.html
		var tmpl = template.Must(template.ParseFiles("view.html"))
		// Eksekusi template dengan data superhero dan menulis hasilnya ke response writer
		if err := tmpl.Execute(w, person); err != nil {
			// Menampilkan pesan error jika terjadi kesalahan saat mengeksekusi template
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Menampilkan pesan bahwa server telah dimulai
	fmt.Println("server started at localhost:9000")

	// Mengaktifkan server HTTP dan mendengarkan permintaan pada port 9000
	http.ListenAndServe(":9000", nil)
}
