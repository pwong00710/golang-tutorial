package main

import "fmt"

func main() {
	var ptr *int

	if ptr == nil {
		fmt.Printf("ptr is null pointer!\n")
	} else {
		fmt.Printf("The value of ptr is : %x\n", ptr)
	}
}
