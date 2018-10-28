package main

import (
	"fmt"
	"golang.org/x/tour/pic"
)

func main() {
	//arrayExample()
	//sliceExample()
	//makeSlice()
	//appendSlice()
	//rangeSlice()
	pic.Show(Pic)
}


func arrayExample() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func sliceExample()  {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	// build slices
	a := names[0:2]
	b := names[1:3]
	c := names[1:] // bounds are optional
	fmt.Println(a, b)
	fmt.Println(a, c)

	// modify the slice updates the underlying array
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func makeSlice()  {

	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func appendSlice(){
	var s []int
	printIntSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printIntSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printIntSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printIntSlice(s)
}

func rangeSlice(){
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func Pic(dx, dy int) [][]uint8 {

	var a,b uint8

	pow := make([][]uint8, dx)
	for i := range pow {
		pow[i] = make([]uint8, dy)
		for j := range pow[i]{

			a = uint8(i)
			b = uint8(j)

			//pow[i][j] = a + b

			if a> b {
				pow[i][j] = a*a - b*b
			}else{
				pow[i][j] = b*b - a*a
			}
		}
	}
return pow
}

func printIntSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}



func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}