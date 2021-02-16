package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
)
func add(x int, y int) (z int) {
	z = x+y
	return
}

func multiply(x int, y int) int {
	return x*y
}

func TwiningOp(x int, y int) (z int, a int){
	z = (x+y)/2
	a = (x*y)/2
	return
}

var IamBool, Bool2 bool // bool variable for checking true or false
var i int // variablr for global Integer var

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())

	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Println(add(3,4))
	fmt.Println(multiply(3,4))

	fmt.Println("Bye bye!")
	fmt.Printf("Now you have %v problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)

	avg1, avg2 := TwiningOp(5,7)
	fmt.Println(avg1, avg2)

	fmt.Println(IamBool, Bool2, i)
}
