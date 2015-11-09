package main

import "fmt"

func main() {

	var pi = 3.1415

	pi = pi + 1

	fmt.Printf("Hello World %.3f \n", pi)

	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}
	age := make(map[string]int)

	age["girish"] = 27
	age["rashmi"] = 26

	fmt.Println(age["girish"])

	listNums := []float64{1, 2, 3, 4, 5}

	num1, num2 := next2Vals(5)
	fmt.Println(num1, num2)

	fmt.Println(subtractThem(1, 2, 3, 4, 5))

	fmt.Println("Sum: ", addThemUp(listNums))

	num3 := 3

	doubleNum := func() int {
		num3 *= 2

		return num3
	}

	fmt.Println(doubleNum())
	fmt.Println(doubleNum())

	var x (
		varA = 2
		varB = 3
	)
	fmt.Println(x)

}

func subtractThem(args ...int) int {
	finalValue := 0
	for _, value := range args {
		finalValue -= value
	}
	return finalValue
}

func next2Vals(number int) (int, int) {
	return number + 1, number + 2
}

func addThemUp(numbers []float64) float64 {
	sum := 0.0

	for _, value := range numbers {
		sum += value
	}
	return sum
}
