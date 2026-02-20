package main

import "fmt"

func main() {
	// Konstanta: nilainya tidak bisa diubah
	const namaToko = "Toko Farhan"
	const pi = 3.14

	// Variabel: nilainya bisa diubah
	var hargaBarang = 50000
	var jumlahBeli = 3

	// Operasi Matematika sedernaha
	totalHarga := hargaBarang * jumlahBeli

	// Output
	fmt.Println("Selamat datang di", namaToko)
	fmt.Println("Harga barang:", hargaBarang)
	fmt.Println("Jumlah beli:", jumlahBeli)
	fmt.Println("Total harga:", totalHarga)

	// namaToko = "Toko Berkah" // Ini akan menyebabkan error karena namaToko adalah konstanta
}
