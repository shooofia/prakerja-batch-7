package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	// Langkah 1: Buat map untuk menyimpan jumlah kemunculan setiap angka
	occurrences := make(map[int]int)

	// Langkah 2: Hitung jumlah kemunculan setiap angka pada string input
	for _, char := range angka {
		num, _ := strconv.Atoi(string(char))
		occurrences[num]++
	}

	// Langkah 3: Buat slice untuk menyimpan angka-angka yang hanya muncul sekali
	var result []int

	// Langkah 4: Tambahkan angka-angka yang hanya muncul sekali ke dalam slice
	for num, count := range occurrences {
		if count == 1 {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
