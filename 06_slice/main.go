package main

import "fmt"

func main() {
	// Membuat Slice berisi hobi
	hobi := []string{"Coding", "Lari", "Membaca"}

	// Menambah isi Slice menggunakan 'append'
	hobi = append(hobi, "Traveling")

	// Menampilkan isi Slice
	fmt.Println("Daftar Hobi saya:")

	// Menggabungkan ilmu Looping dan Slice
	for i, data := range hobi {
		fmt.Printf("%d. %s\n", i+1, data)
	}

	// penjelasan:
	// 1. Kita membuat Slice dengan tipe data string yang berisi beberapa hobi.
	// 2. Kita menggunakan fungsi 'append' untuk menambahkan hobi baru ke dalam Slice.
	// 3. Kita menggunakan Looping 'for' dengan 'range' untuk menampilkan setiap hobi beserta indeksnya.
	// i+1 digunakan untuk menampilkan nomor urut mulai dari 1, karena indeks dalam Slice dimulai dari 0.
	// Output yang dihasilkan akan menampilkan daftar hobi dengan nomor urutnya.
}
