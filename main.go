package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)
var IamBool, Bool2 bool // bool variable for checking true or false
var i int = 100 // variable for global Integer var


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

// to declare more than one variables in same var clause
var (
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	anotherComp = complex(1,7)
)

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

	fmt.Println(z, "It seems fine",anotherComp)
	fmt.Println(z*anotherComp)
	fmt.Println(real(z), imag(z*anotherComp))

	// Type Conversion
	var f float64 = (3*4)
	var tt uint = uint(f)
	fmt.Println(f,tt)

	fmt.Printf("f is of type %T\n", f)

	// constant --- const don't use shortcut operator
	const x = true
	fmt.Println(x)

	// We can have untyped constants (Constant are very high precision)
	const (
		Big = 1 << 100
		Small = Big >> 99
	)
	smallInt := Small*10 + 1
	fmt.Printf("f is of type %T\n", smallInt)

	// for loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	// continued for loop
	for ; sum<1000 ; {
		// 	or we can write for sum<1000{
		//	} --- looks like while loop
		sum += sum
	}
	// forever loop
	// for {
	// }
	// break; can be used here
	fmt.Println(sum)

	// if {
	// } --- like for loop => without () but {} are required

	// if with short statement
	if v:= 4; v < 5 {
		fmt.Println("we are in if with short statement")
	}

	// switch can also be used as short statement like if
	// switch x:= 100; z {
	// case 1:
	//  fmt.Println("1")
	// case 2:
	// 	fmt.Println("2")
	// default:
	// 	fmt.Println(x)
	//}

	// switch can be used without any condition like
	// switch {
	// case __any condition__:
	// case __any condition__:
	// } --- It will act like if-then-else

	// defer and stacking defers
	// implemented immediately and executed at last r = append(answer, converter(value))while returning from function
	// execution of defer will be poped from stack in case of multiple defer statement



}
