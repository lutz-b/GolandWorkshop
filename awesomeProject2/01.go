package main

import (
"fmt"
"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(swapShort("Hello", "World"))

	// one return value not relevant
	a, _ := swapShort("Hello", "World")
	fmt.Println(a)

	fmt.Println(namedReturnValues("Hello", "World"))
	variableDeclaration()

	main()
}


// short parameter type
//multiple return values
func swapShort(x, y string) (string, string) {
	return y, x
}

// short parameter type
//multiple return values
func namedReturnValues(x, y string) (a,b string) {
	b = x
	a = y
	return
}

func variableDeclaration(){
	var c, python, java bool
	var i, j int = 1, 2
	k := 3
	var l int
	fmt.Println("variableDeclaration:")
	fmt.Println(i, j, k, l, c, python, java)
	fmt.Printf("%T, %T,%T,%T,%T,%T,%T \n,", i, j, k, l, c, python, java)
	fmt.Printf("%v, %v, %v, %v, %v, %v, %v \n",i, j, k, l, c, python, java)


}
