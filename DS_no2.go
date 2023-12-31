package main

import "fmt"

func Mapping(slice []string) map[string]int {
	occurrences := make(map[string]int)
	for _, str := range slice {
		occurrences[str]++
	}

	return occurrences
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{}))
}
