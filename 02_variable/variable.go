package main

import "fmt"

func main() {
	// 1. Deklarasi variabel dengan tipe data (Formal)
	var name string = "Farhan"
	var age int = 29
	var isStudent bool = true
	// var tesUnused string = "Variable ini tidak digunakan" // Variabel ini tidak digunakan, akan menghasilkan peringatan

	// 2. Deklarasi variabel tanpa tipe data (Go akan menebak tipe data berdasarkan nilai yang diberikan)
	var city = "Bandung"
	var height = 165.5

	// 3. Deklarasi variabel dengan pendek (Short Variable Declaration)
	country := "Indonesia"
	weight := 55.5

	// Menampilkan nilai variabel
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Student:", isStudent)
	fmt.Println("City:", city)
	fmt.Println("Height:", height)
	fmt.Println("Country:", country)
	fmt.Println("Weight:", weight)
}
