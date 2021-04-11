package main 

import "fmt"

func main() {
	nameAgeMap := map[string]int {
		"Anwesh" : 28,
		"Ajay" : 21,
		"Deep": 26,
	}

	fmt.Printf("%#v\n", nameAgeMap)

	nameIdMap := make(map[string]int)
	nameIdMap["Anwesh"] = 1
	nameIdMap["Ajay"] = 2
	nameIdMap["Deep"] = 3 
	fmt.Printf("%#v\n", nameIdMap)
}