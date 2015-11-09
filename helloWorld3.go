package main

import "fmt"

func main() {
	x := 0
	changeXVal(x)
	fmt.Println(x)
	changeXValNow(&x)
	fmt.Println(x)
	fmt.Println("Memory Address for x = ", &x)

	yPtr := new(int)

	changeXValNow(yPtr)

	fmt.Println(*yPtr)

	rect1 := Rectangle{0, 50, 10, 10}

	fmt.Println(rect1.width, rect1.leftX)

	fmt.Println("Area of rectangle", rect1.area())
}

// func attached to a struct
func (rect *Rectangle) area() float64 {
	return rect.width * rect.height
}

type Rectangle struct {
	leftX  float64
	topY   float64
	height float64
	width  float64
}

// by value, default
func changeXValNow(x *int) {
	*x = 2
}

//by reference
func changeXVal(x int) {
	x = 2
}
