package main

import "fmt"

func main() {
	buah := []string{"apel", "jeruk", "mangga"} // membuat slice dengan literal
	buah = append(buah, "nanas")                // menambahkan elemen baru ke dalam slice
	semuaBuah := buah                           // membuat slice baru yang merujuk ke data yang sama dengan slice buah
	for i, data := range semuaBuah {
		fmt.Printf("%d. %s\n", i+1, data)
	}

	fmt.Println("Jumlah buah yang saya punya adalah", len(buah))   // mencetak jumlah elemen dalam slice
	fmt.Println("Tapi buah yang paling saya suka adalah", buah[3]) // mencetak elemen ke-4 dari slice buah
	fmt.Println("Sebenarnya saya tidak suka buah", semuaBuah[1])   // mencetak elemen ke-2 dari slice semuaBuah
	fmt.Println("Tapi saya suka buah", semuaBuah[0:2])             // mencetak elemen ke-1 sampai ke-2 dari slice semuaBuah
}
