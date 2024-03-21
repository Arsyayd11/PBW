package main

import "fmt"

type Mahasiswa struct {
	Nama    string
	NPM     string
	Jurusan string
}

func main() {
	// Membuat map untuk menyimpan data mahasiswa
	dataMahasiswa := make(map[string]Mahasiswa)

	// Memasukkan data mahasiswa ke dalam map
	dataMahasiswa["4522210117"] = Mahasiswa{"Arsya", "4522210117", "Teknik Informatika"}
	dataMahasiswa["4522210128"] = Mahasiswa{"Gibran", "4522210128", "Teknik Informatika"}
	dataMahasiswa["4522210133"] = Mahasiswa{"Hotra", "4522210133", "Teknik Informatika"}

	// Menampilkan daftar nama mahasiswa dengan perulangan
	fmt.Println()
	fmt.Println("Daftar Nama Mahasiswa:")
	for _, mahasiswa := range dataMahasiswa {
		fmt.Println(mahasiswa.NPM, mahasiswa.Nama)
	}

	// Mencari data mahasiswa berdasarkan NPM dengan inputan
	var npm string
	fmt.Println()
	fmt.Print("\nMasukkan NPM mahasiswa yang ingin dicari: ")
	fmt.Scanln(&npm)

	mahasiswa, cari := dataMahasiswa[npm]
	if cari {
		fmt.Println("\nData Mahasiswa:")
		fmt.Println("Nama:", mahasiswa.Nama)
		fmt.Println("NPM:", mahasiswa.NPM)
		fmt.Println("Jurusan:", mahasiswa.Jurusan)
	} else {
		fmt.Println("\nMahasiswa dengan NPM tersebut tidak ditemukan.")

	}
	fmt.Println()
}
