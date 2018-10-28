package main

import (
	"fmt"
	"runtime"
)

// print os
func main() {
	fmt.Print("Go runs on ")

	//no break needed

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}
