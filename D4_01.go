package main

import (
	"fmt"
)

// Struktur Mobil
type Mobil struct {
	Type   string
	FuelIn float64
}

// Method untuk menghitung perkiraan jarak yang bisa ditempuh
func (c Mobil) CalculateEstimatedDistance() float64 {
	fuelEfficiency := 1.5 // L / mill
	estimatedDistance := c.FuelIn * fuelEfficiency
	return estimatedDistance
}

func main() {
	// Membuat objek Mobil
	mobil := Mobil{
		Type:   "Sedan",
		FuelIn: 30, // Diisi dengan jumlah bahan bakar dalam liter
	}

	// Memanggil method untuk menghitung perkiraan jarak
	estimatedDistance := mobil.CalculateEstimatedDistance()

	// Menampilkan hasil perkiraan jarak
	fmt.Printf("Perkiraan jarak yang bisa ditempuh dengan %s adalah %.2f km\n", mobil.Type, estimatedDistance)
}
