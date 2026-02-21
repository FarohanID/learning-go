package main

import "fmt"

func main() {
	nilai := 85

	fmt.Println("Nilai kamu adalah:", nilai)

	if nilai >= 80 {
		fmt.Println("Selamat! Kamu lulus dengan pujian.")
	} else if nilai >= 60 {
		fmt.Println("Kamu lulus standar.")
	} else {
		fmt.Println("Maaf, kamu harus belajar lebih giat lagi.")
	}
}
