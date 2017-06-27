package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// generates a random string
func srand(min, max int, readable bool) string {

	var length int
	var char string

	if min < max {
		length = min + rand.Intn(max-min)
	} else {
		length = min
	}

	if readable == false {
		char = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	} else {
		char = "ABCDEFHJLMNQRTUVWXYZabcefghijkmnopqrtuvwxyz23479"
	}

	buf := make([]byte, length)
	for i := 0; i < length; i++ {
		buf[i] = char[rand.Intn(len(char)-1)]
	}
	return string(buf)
}

// For testing only
func main() {
	println(srand(5, 5, true))
	println(srand(5, 5, true))
	println(srand(5, 5, true))
	println(srand(5, 5, false))
	println(srand(5, 7, true))
	println(srand(5, 10, false))
	println(srand(5, 50, true))
	println(srand(5, 10, false))
	println(srand(5, 50, true))
	println(srand(5, 10, false))
	println(srand(5, 50, true))
	println(srand(5, 10, false))
	println(srand(5, 50, true))
	println(srand(5, 4, true))
	println(srand(5, 400, true))
	println(srand(6, 5, true))
	println(srand(6, 5, true))
	println(srand(6, 5, true))
	println(srand(6, 5, true))
	println(srand(6, 5, true))
	println(srand(6, 5, true))
	println(srand(6, 5, true))
	println(srand(6, 5, true))
	println(srand(6, 5, true))
	println(srand(6, 5, true))
	println(srand(6, 5, true))
	println(srand(6, 5, true))
}
