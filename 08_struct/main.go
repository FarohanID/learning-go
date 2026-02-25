package main

import "fmt"

type Smartphone struct {
	Merk  string
	Model string
	Harga int
}

func main() {
	hpSaya := Smartphone{
		Merk:  "Samsung",
		Model: "S23",
		Harga: 12000000,
	}

	// Menggunakan %v (Hanya isi)
	fmt.Printf("Isi hpSaya: %v\n", hpSaya) // Menggunakan %v akan menampilkan isi struct tanpa nama kolomnya
	fmt.Println("HP saya: ", hpSaya.Merk)  // Menggunakan dot notation untuk mengakses field tertentu

	// Menggunakan %+v (Ada nama kolomnya - SANGAT BERGUNA!)
	fmt.Printf("Detail hpSaya: %+v\n", hpSaya)

	// Cara manual yang rapi
	fmt.Printf("Saya punya %s %s seharga Rp %d\n", hpSaya.Merk, hpSaya.Model, hpSaya.Harga)

}
