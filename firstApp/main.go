package main

import (
	"fmt"
	"math"
)

const i float32 = 42.8

const (
	_ = iota - 6
	a
	b
)

func main() {
	fmt.Printf("%v of type %T\n", i, i)

	var j string
	j = fmt.Sprint(i)
	fmt.Printf("%v of type %T\n", j, j)

	var k bool = false
	fmt.Printf("%v of type %T\n", !k, k)

	var f int = 10
	var g int = 3
	fmt.Println(float32(f/g))

	var h int = 20
	var p = h << 2
	fmt.Println(p)

	var grades = []int{4,5}
	fmt.Printf("%v of type %T\n", len(grades), grades)

	var students = make([]int, 5, 16)
	fmt.Println(cap(students))

	var complex complex64 = 2i
	fmt.Printf("%v of type %T\n", complex, complex)

	const helloThere float64 = math.Pi
	fmt.Printf("%v of type %T\n", helloThere, helloThere)

	const i = 12
	var l int64 = 32
	fmt.Printf("%v of type %T\n", i + l, i + l)

	fmt.Printf("%v of type %T\n", a, a)
	fmt.Printf("%v of type %T\n", b, b)
}
