# LAPORAN PRAKTIKUM KE-2
# PEMROGRAMAN BERBASIS WEB
# KELAS A
Disusun Untuk Memenuhi Tugas Mata Kuliah Prak. Pemrograman Berbasis Web



Disusun oleh:
Arsya Yan Duribta
4522210117

Dosen Pengampu: 
Adi Wahyu Pribadi, S.Si., M.Kom. 

Program Studi Teknik Informatika 
Fakultas Teknik Universitas Pancasila 
2023/2024



### Link Repository Github: 


### Tugas 1
Source code:
package main


import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // Membuat scanner untuk input dari konsol
    scanner := bufio.NewScanner(os.Stdin)

    // Input nama
    fmt.Println()
    fmt.Print("Masukkan nama Anda: ")
    scanner.Scan()
    nama := scanner.Text()

    // Input usia
    fmt.Print("Masukkan usia Anda: ")
    scanner.Scan()
    var usia int
    fmt.Sscanf(scanner.Text(), "%d", &usia)

    // Menentukan kategori usia
    var kategoriUsia string
    if usia < 18 {
        kategoriUsia = "Anak-anak"
    } else {
        kategoriUsia = "Dewasa"
    }

    fmt.Println()
    fmt.Printf("Selamat datang, %s! Anda termasuk kategori %s.\n", nama, kategoriUsia)
    fmt.Println()
}

Output


### Tugas 2
Source code:
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

Output


### Tugas 3
Jelaskan Source code dibawah:
package main

import "fmt"

func main() {
    panjang := 5
    lebar := 3
    luas := hitungLuasPersegiPanjang(panjang, lebar)
    keliling := hitungKelilingPersegiPanjang(panjang, lebar)
    fmt.Println("Luas persegi panjang:", luas)
    fmt.Println("Keliling persegi panjang:", keliling)
}

func hitungLuasPersegiPanjang(panjang int, lebar int) int {
    return panjang * lebar
}

func hitungKelilingPersegiPanjang(panjang int, lebar int) int {
    return 2 * (panjang + lebar)
}

Jawab:
Source code di atas adalah sebuah program golang sederhana untuk menghitung luas dan keliling persegi panjang berdasarkan panjang dan lebar yang telah ditetapkan. 
1. Package dan Import: 
Program dimulai dengan mendefinisikan package main.
Mengimport package "fmt" yang digunakan untuk melakukan input-output format pada program Go.
2. Fungsi main():
Fungsi main() adalah fungsi utama yang akan dieksekusi pertama kali saat program dijalankan.
Dalam fungsi main(), nilai panjang dan lebar persegi panjang didefinisikan masing-masing dengan nilai panjang = 5 dan lebar = 3.
Kemudian, fungsi hitungLuasPersegiPanjang() dipanggil dengan parameter panjang dan lebar, hasilnya disimpan dalam variabel luas.
Fungsi hitungKelilingPersegiPanjang() juga dipanggil dengan parameter panjang dan lebar, hasilnya disimpan dalam variabel keliling.
Hasil perhitungan luas dan keliling kemudian dicetak menggunakan fungsi fmt.Println("Luas persegi panjang:", luas), dan fmt.Println("Keliling persegi panjang:", keliling) .
3. Fungsi hitungLuasPersegiPanjang():
Fungsi ini memiliki dua parameter, yaitu panjang dan lebar persegi panjang, yang keduanya bertipe data integer.
Mengembalikan hasil perkalian panjang dengan lebar, yang merupakan rumus untuk menghitung luas persegi panjang.
4. Fungsi hitungKelilingPersegiPanjang():
Sama seperti fungsi sebelumnya, memiliki dua parameter, panjang dan lebar persegi panjang serta tipe data integer.
Mengembalikan hasil perhitungan 2 kali (panjang + lebar), sesuai rumus untuk menghitung keliling persegi panjang.
Jadi, keseluruhan program tersebut akan menghasilkan output berupa luas dan keliling persegi panjang berdasarkan nilai panjang dan lebar yang telah ditentukan di dalam fungsi main().

### Tugas 4
Jelaskan Source code dibawah:
package main
 
import "fmt"
 
func main() {
    panjang := 5
    lebar := 3
 
    luas, keliling := hitungLuasKelilingPersegiPanjang(panjang, lebar)
 
    fmt.Println("Luas persegi panjang:", luas)
    fmt.Println("Keliling persegi panjang:", keliling)
}
 
func hitungLuasKelilingPersegiPanjang(panjang int, lebar int) (luas int,
keliling int) {
    luas = panjang * lebar
    keliling = 2 * (panjang + lebar)
    return
}

Jawab:
Source code di atas merupakan program yang sama dengan program pada Tugas 03, namun menggunakan pengembalian nilai ganda dari sebuah fungsi. Berikut adalah penjelasan dari setiap bagian dari kode tersebut:
1. Package dan Import: 
Program dimulai dengan mendefinisikan package main.
Mengimport package "fmt" yang digunakan untuk melakukan input-output format pada program Go.
2. Fungsi main():
Fungsi main() adalah fungsi utama yang akan dieksekusi pertama kali saat program dijalankan.
Dalam fungsi main(), nilai panjang dan lebar persegi panjang ditetapkan masing-masing dengan nilai panjang = 5 dan lebar = 3.
Kemudian, fungsi hitungLuasKelilingPersegiPanjang() yang memiliki parameter panjang dan lebar, dan mengembalikan dua nilai yaitu luas dan keliling, yang kemudian disimpan dalam variabel luas dan keliling.
Hasil perhitungan luas dan keliling kemudian dicetak menggunakan fungsi fmt.Println("Luas persegi panjang:", luas), dan fmt.Println("Keliling persegi panjang:", keliling).
3. Fungsi hitungLuasKelilingPersegiPanjang():
Fungsi ini memiliki dua parameter, yaitu panjang dan lebar persegi panjang, keduanya bertipe data integer.
Mengembalikan dua nilai sekaligus, yaitu luas dan keliling.
Di dalam fungsi, luas dihitung dengan perkalian panjang dan lebar, sedangkan keliling dihitung dengan rumus 2 kali (panjang + lebar).
Karena dalam deklarasi fungsi sudah menyebutkan bahwa akan mengembalikan luas dan keliling, maka pada bagian return tidak perlu menyebutkan nama variabelnya lagi, cukup dengan return.

### Tugas 5
Jelaskan perbedaan antara Tugas 03 dan Tugas 04!
Jawab:
Perbedaan antara Tugas 03 dan Tugas 04 terletak pada cara pengembalian nilai dari fungsi hitungLuasKelilingPersegiPanjang():
Tugas 03:
Pada Tugas 03, fungsi hitungLuasPersegiPanjang() dan hitungKelilingPersegiPanjang() masing-masing mengembalikan satu nilai, yaitu luas dan keliling persegi panjang.
Di dalam fungsi main(), setelah memanggil kedua fungsi tersebut, hasilnya disimpan dalam variabel luas dan keliling secara terpisah.
Tugas 04:
Pada Tugas 04, fungsi hitungLuasKelilingPersegiPanjang() mengembalikan dua nilai sekaligus, yaitu luas dan keliling persegi panjang.
Di dalam fungsi main(), setelah memanggil fungsi tersebut, kedua hasil perhitungan (luas dan keliling) disimpan dalam satu baris kode menggunakan deklarasi luas, keliling := hitungLuasKelilingPersegiPanjang(panjang, lebar).
Oleh karena itu, perbedaan utama antara Tugas 03 dan Tugas 04 adalah bagaimana nilai fungsi dikembalikan. Tugas 03 menggunakan pengembalian nilai tunggal untuk perhitungan luas dan keliling, sedangkan Tugas 04 menggunakan pengembalian nilai ganda untuk kedua perhitungan sekaligus.


### Link Repository Github: