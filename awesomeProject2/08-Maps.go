package main

// map creates a hashmap / key-value-pair

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	//simpleMap()
	//modifyMap()
	wc.Test(WordCount)
}

func simpleMap() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["Bell"])
}

func modifyMap(){
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func WordCount(s string) map[string]int {
	x:= strings.Fields(s)
	m := make(map[string]int)

	for text, i := range x {
		v, ok := m[i]
		if ok {
			print(v) // must use v for something
			print(text) // must use v for something
			m[i] += m[i]
		}else{
			m[i] = 1
		}
	}

	return m


	//return map[string]int{"x": 1}
}