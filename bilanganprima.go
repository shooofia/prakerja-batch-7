package main

import (
	"fmt"
	"math"
)

// Fungsi untuk mengecek apakah sebuah bilangan prima atau bukan
func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	// Mencari faktor-faktor bilangan
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var x int
	fmt.Println("Masukkan batas bilangan awal: ")
	fmt.Scan(&x)
	var y int
	fmt.Print("Masukkan batas bilangan akhir: ")
	fmt.Scan(&y)

	fmt.Println("Bilangan prima antara", x, "dan", y, "adalah:")
	for i := x; i <= y; i++ {
		if isPrime(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
