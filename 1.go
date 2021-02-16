package main

import (
	"fmt"
	"time"
	"math/rand"
)
func add(x int, y int) int {
	return x+y
}
func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())

	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Println(add(3,4))

}
