package main

import "fmt"

func main() {
	var start, end int

	fmt.Print("Masukkan nilai awal rentang: ")
	fmt.Scan(&start)

	fmt.Print("Masukkan nilai akhir rentang: ")
	fmt.Scan(&end)

	fmt.Println("Kelipatan 7 antara", start, "dan", end, "adalah:")
	for i := start; i <= end; i++ {
		if i%7 == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
