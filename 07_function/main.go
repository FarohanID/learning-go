package main

import "fmt"

// Ini adalah "Mesin" penyapa kita
func sapaOrang(nama string, umur int) {
	fmt.Printf("Halo %s, umur kamu %d tahun ya?\n", nama, umur)
}

func cekGenapGanjil(nomor int) {
	if nomor%2 == 0 {
		fmt.Printf("%d adalah angka genap.\n", nomor)
	} else {
		fmt.Printf("%d adalah angka ganjil.\n", nomor)
	}
}

func main() {
	// Kita panggil mesinnya di sini
	sapaOrang("Farhan", 29)
	fmt.Println("Farhan memiliki nomor 1, jadi kita cek apakah itu genap atau ganjil:")
	cekGenapGanjil(1)
	sapaOrang("Andi", 20)
	fmt.Println("Andi memiliki nomor 0, jadi kita cek apakah itu genap atau ganjil:")
	cekGenapGanjil(0)
	sapaOrang("Budi", 35)
	fmt.Println("Budi memiliki nomor 3, jadi kita cek apakah itu genap atau ganjil:")
	cekGenapGanjil(3)

}
