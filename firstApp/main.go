package main

import (
	"fmt"
	"math"
	"reflect"
)

const i float32 = 42.8

const (
	_ = iota - 6
	a
	b
)

type Animal struct {
	name string `required max: "100"`
	origin string
}

func main() {
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("name")

	fmt.Println(field.Tag)

	fmt.Printf("%v of type %T\n", i, i)

	var j string
	j = fmt.Sprint(i)
	fmt.Printf("%v of type %T\n", j, j)

	var k bool = false
	fmt.Printf("%v of type %T\n", !k, k)

	var f int = 10
	var g int = 3
	fmt.Println(float32(f / g))

	var h int = 20
	var p = h << 2
	fmt.Println(p)

	var grades = []int{4, 5}
	fmt.Printf("%v of type %T\n", len(grades), grades)

	var students = make([]int, 5, 16)
	fmt.Println(cap(students))

	var complex complex64 = 2i
	fmt.Printf("%v of type %T\n", complex, complex)

	const helloThere float64 = math.Pi
	fmt.Printf("%v of type %T\n", helloThere, helloThere)

	const i = 12
	var l int64 = 32
	fmt.Printf("%v of type %T\n", i+l, i+l)

	fmt.Printf("%v of type %T\n", a, a)
	fmt.Printf("%v of type %T\n", b, b)

	var i1 = []int{4, 5, 7}
	var i2 = []int{6, 1, 2}
	i2 = append(i2, i1...)
	var i3 = i2[:len(i2)-1]
	var i4 = append(i2[:2], i2[3:]...)
	fmt.Println(i2)
	fmt.Println(i3)
	fmt.Println(i4)

	i5 := [...]int{}
	fmt.Println(i5)

	var m = make(map[string]int)
	m = map[string]int{"California": 77777}
	fmt.Printf("%v of type %T\n", m, m)

	var statePopulations = map[string]int{
		"Hello": 1234,
	}

	fmt.Printf("%v of type %T\n", statePopulations, statePopulations)

	fmt.Printf("Value for California: %v\n", m["California"])

	pop, k := statePopulations["Georgia"]

	fmt.Println(pop, k)

	fmt.Println(statePopulations["There"])

	if x:=8>6; x {
		fmt.Println("It works")
	}
	if returnTrue() {
		fmt.Println("It does indeed work")
	}
}

func returnTrue() bool {
	return true;
}
