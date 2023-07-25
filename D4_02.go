package main

import (
	"fmt"
)

// Struktur Student
type Student struct {
	Name  string
	Score int
}

// Metode untuk menghitung skor rata-rata
func CalculateAverageScore(students []Student) float64 {
	totalScore := 0

	for _, student := range students {
		totalScore += student.Score
	}

	averageScore := float64(totalScore) / float64(len(students))
	return averageScore
}

// Metode untuk mencari siswa dengan skor minimum dan maksimum
func FindMinMaxScoreStudents(students []Student) (string, string) {
	minScoreStudent := students[0]
	maxScoreStudent := students[0]

	for _, student := range students {
		if student.Score < minScoreStudent.Score {
			minScoreStudent = student
		}

		if student.Score > maxScoreStudent.Score {
			maxScoreStudent = student
		}
	}

	return minScoreStudent.Name, maxScoreStudent.Name
}

func main() {
	var students []Student

	// Mengisi data siswa sebanyak 5 siswa
	for i := 0; i < 5; i++ {
		var name string
		var score int

		fmt.Printf("Input %d student's name: ", i+1)
		fmt.Scanln(&name)

		fmt.Printf("Input %d student's score: ", i+1)
		fmt.Scanln(&score)

		students = append(students, Student{Name: name, Score: score})
	}

	// Menghitung skor rata-rata
	averageScore := CalculateAverageScore(students)

	// Mencari siswa dengan skor minimum dan maksimum
	minScoreStudent, maxScoreStudent := FindMinMaxScoreStudents(students)

	// Menampilkan hasil
	fmt.Printf("average score: %.2f\n", averageScore)
	fmt.Printf("min score of student: %s (%d)\n", minScoreStudent, students[0].Score)
	fmt.Printf("max score of student: %s (%d)\n", maxScoreStudent, students[0].Score)
}
