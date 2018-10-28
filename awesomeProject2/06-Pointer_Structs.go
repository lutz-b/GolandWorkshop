package main

import "fmt"

type vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	w := Vertex{3, 3}
	p := &v
	p.X = 1e9
	p = &w
	(*p).Y = 99
	fmt.Println(v)
	fmt.Println(w)
	fmt.Println(p)
	fmt.Println(*p)
}