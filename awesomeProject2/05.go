package main

import "fmt"

// defer

func main() {
	//deferTest()

	defer fmt.Println("world")
	fmt.Println("hello")
}

func deferTest(){
	i:= 10
	defer fmt.Println( i)
	i += 5
	fmt.Print( "i=")
}