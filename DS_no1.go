package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	// Buat map untuk melacak elemen yang sudah ada dalam array hasil
	seen := make(map[string]bool)
	result := make([]string, 0)

	// Tambahkan elemen dari arrayA ke array hasil
	for _, str := range arrayA {
		if !seen[str] {
			result = append(result, str)
			seen[str] = true
		}
	}

	// Tambahkan elemen dari arrayB ke array hasil (jika belum ada dalam map)
	for _, str := range arrayB {
		if !seen[str] {
			result = append(result, str)
			seen[str] = true
		}
	}
	return result
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yohimitsu"}, []string{"devil jin", "yohimitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))

}
