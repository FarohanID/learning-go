package main

import (
	"errors"
	"fmt"
)

func cekLogin(username, password string) error {
	if username != "admin" || password != "admin123" {
		return errors.New("username atau password salah")
	}
	return nil
}

func main() {
	err := cekLogin("admin", "admin123")
	if err != nil {
		fmt.Println("Login gagal:", err)
	} else {
		fmt.Println("Login berhasil!")
	}

	err = cekLogin("user", "pass")
	if err != nil {
		fmt.Println("Login gagal:", err)
	} else {
		fmt.Println("Login berhasil!")
	}
}

// Output:
// Login berhasil!
// Login gagal: username atau password salah
