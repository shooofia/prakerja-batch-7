package main

import "fmt"

func main() {
	var a, b, t float64

	fmt.Print("Masukkan panjang sisi atas trapesium (a): ")
	fmt.Scan(&a)

	fmt.Print("Masukkan panjang sisi bawah trapesium (b): ")
	fmt.Scan(&b)

	fmt.Print("Masukkan tinggi trapesium (t): ")
	fmt.Scan(&t)

	// Menghitung luas trapesium
	luas := 0.5 * (a + b) * t

	fmt.Println("Luas trapesium adalah:", luas)
}
