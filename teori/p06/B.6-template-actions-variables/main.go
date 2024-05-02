package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Info adalah struktur untuk informasi tambahan tentang seseorang
type Info struct {
	Affiliation string // Afiliasi
	Address     string // Alamat
}

// Person adalah struktur untuk merepresentasikan seseorang
type Person struct {
	Name    string   // Nama
	Gender  string   // Jenis Kelamin
	Hobbies []string // Hobi-hobi
	Info    Info     // Info tambahan
}

// GetAffiliationDetailInfo adalah metode untuk mendapatkan detail informasi afiliasi
func (t Info) GetAffiliationDetailInfo() string {
	return "have 31 divisions" // mengembalikan detail informasi afiliasi
}

func main() {
	// Handler untuk route "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Data person
		var person = Person{
			Name:    "Bruce Wayne",
			Gender:  "male",
			Hobbies: []string{"Reading Books", "Traveling", "Buying things"},
			Info:    Info{"Wayne Enterprises", "Gotham City"},
		}

		// Parsing file template view.html
		var tmpl = template.Must(template.ParseFiles("view.html"))
		// Eksekusi template dengan data person dan menulis hasilnya ke response writer
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
