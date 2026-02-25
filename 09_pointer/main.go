package main

import "fmt"

func ulangTahun(umur *int) {
	// Kita bisa mengubah nilai yang ditunjuk oleh pointer umur
	// *umur akan mengakses nilai yang berada di alamat tersebut, jadi kita bisa menambah umur lewat pointer
	*umur = *umur + 1
}

func main() {
	nama := "Farhan"
	umur := 29

	fmt.Printf("Sebelum ulang tahun, nama: %s, umur: %d\n", nama, umur)

	// Kita buat pointer yang menunjuk ke alamat variabel 'nama'
	// &nama akan memberikan alamat memori dari variabel nama
	// & (and) adalah operator untuk mendapatkan alamat memori dari suatu variabel
	alamatNama := &nama

	// Kita buat pointer yang menunjuk ke alamat variabel 'umur'
	alamatUmur := &umur

	// Memanggil fungsi ulangTahun dengan pointer alamatUmur
	ulangTahun(alamatUmur)

	fmt.Printf("Setelah ulang tahun, nama: %s, umur: %d\n", nama, *alamatUmur)

	fmt.Println("Isi variabel nama:", nama)
	fmt.Println("Alamat memori nama:", alamatNama) // Akan muncul kode aneh seperti 0xc000...

	// Mengubah isi lewat pointer
	// *alamatNama akan mengakses nilai yang berada di alamat tersebut, jadi kita bisa mengubah isi variabel nama lewat pointer
	// * (bintang) adalah operator dereference untuk mengakses nilai yang berada di alamat memori yang ditunjuk oleh pointer
	*alamatNama = "Farhan - Informatics Graduate"

	fmt.Println("Isi nama setelah diubah lewat pointer:", nama)
}
