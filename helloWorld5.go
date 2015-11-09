package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	sampString := "Hello World"

	fmt.Println(strings.Contains(sampString, "lo"))
	fmt.Println(strings.Index(sampString, "lo"))
	fmt.Println(strings.Count(sampString, "l"))
	fmt.Println(strings.Replace(sampString, "l", "x", 3))

	csvString := "1,2,3,4"

	fmt.Println(strings.Split(csvString, ","))

	listOfLetters := []string{"c", "z", "a"}

	sort.Strings(listOfLetters)
	fmt.Println(listOfLetters)

	listOfNum := strings.Join([]string{"3", "2"}, ", ")
	fmt.Println(listOfNum)
}
