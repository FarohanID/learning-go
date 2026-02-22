package main

import "fmt"

func main() {
	// hitung mundur
	for i := 10; i >= 1; i-- {
		fmt.Printf("Hitung mundur: %d\n", i)
	}
	fmt.Println("BOOM! Roket Meluncur!")

	// Penjelasan:
	// for: Kata kunci untuk memulai loop.
	// i := 10: Titik start (mulai dari 10).
	// i >= 1: Syarat (terus jalan selama i lebih besar atau sama dengan 1).
	// i--: Langkah (setiap putaran selesai, kurangi i sebanyak 1).
	// Catatan: Di sini saya pakai fmt.Printf dan %d. Ini adalah cara keren untuk memasukkan angka ke dalam kalimat. %d artinya "taruh angka di sini".

	// hitung naik
	for j := 1; j <= 10; j++ {
		fmt.Printf("Hitung naik: %d\n", j)
	}

	// Penjelasan:
	// j := 1: Titik start (mulai dari 1).
	// j <= 10: Syarat (terus jalan selama j kurang atau sama dengan 10).
	// j++: Langkah (setiap putaran selesai, tambah j sebanyak 1).

	// Looping dengan kondisi
	counter := 0
	for counter < 5 {
		fmt.Printf("Counter: %d\n", counter)
		counter++
	}

	// Penjelasan:
	// counter := 0: Inisialisasi variabel counter dengan nilai awal 0.
	// counter < 5: Syarat (terus jalan selama counter kurang dari 5).
	// counter++: Langkah (setiap putaran selesai, tambah counter sebanyak 1).
}
