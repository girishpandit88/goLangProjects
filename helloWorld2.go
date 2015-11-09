package main

import "fmt"

func main() {

	defer printTwo() //defers execution until main ends... mainly used for cleanup operations
	printOne()
	fmt.Println(safeDiv(3, 0))
	fmt.Println(safeDiv(3, 2))

	onMyPanic()
	fmt.Println(fact(5))
	fmt.Println(gcd(25, 15))
}

func onMyPanic() {
	defer func() {
		fmt.Println(recover())
	}()

	panic("PANIC!!!")
}

func safeDiv(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()
	solution := num1 / num2
	return solution
}

func printOne() { fmt.Println(1) }
func printTwo() { fmt.Println(2) }

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func fact(num int) int {
	if num == 0 {
		return 1
	}
	return num * fact(num-1)

}
