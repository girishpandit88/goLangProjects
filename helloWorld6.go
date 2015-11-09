package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Create("samp.txt")

	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("This is some random text")
	file.Close()

	stream, err := ioutil.ReadFile("samp.txt")

	if err != nil {
		log.Fatal(err)
	}

	readString := string(stream)

	fmt.Println(readString)

	randInt := 5

	// randFloat := 10.5

	randString := "100"

	// randString2 := "250.5"

	fmt.Println(float64(randInt))
	newInt, _ := strconv.ParseInt(randString, 0, 64)
	newFloat, _ := strconv.ParseFloat(randString, 64)
	fmt.Println(newInt, newFloat)
}
