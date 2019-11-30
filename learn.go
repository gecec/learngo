package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")

	beyondHello()
}

func beyondHello() {
	var x int
	x = 3
	y := 4
	sum, prod := learnMultiple(x, y)
	fmt.Println("sum:", sum, "prod:", prod)
	learnTypes()
}

func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y
}

func learnTypes() {
	str := "Learn Go!"

	s2 := `A "raw" string literal 
	can include line breaks`

	g := 'ג'
	f := 3.14195
	c := 3 + 4i
	var u uint = 7
	var pi float = 22. / 7
	n := byte('\n')

	var a4 [4]int
	a5 := [...]int{3, 1, 5, 10, 100}

	s3 := []int{4, 5, 9}
	s4 := make([]int, 4)
}
