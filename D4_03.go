package main

import (
	"fmt"
)

func getMinMax(numbers ...*int) (min int, max int) {
	// Inisialisasi nilai min dan max dengan nilai pertama dalam array
	min = *numbers[0]
	max = *numbers[0]

	// Loop untuk membandingkan nilai-nilai dalam array
	for _, num := range numbers {
		if *num < min {
			min = *num
		}
		if *num > max {
			max = *num
		}
	}

	return min, max
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)

	// Memanggil fungsi getMinMax dengan argumen pointer
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("%d merupakan nilai min.\n ", min)
	fmt.Println("%d merupakan nilai max.\n", max)
}
