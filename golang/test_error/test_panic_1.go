package main

import (
	"fmt"
	"io"
	"os"
)

var user = os.Getenv("USER")

func init() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("init failed:%v\n", err)
		}
	}()

	if user == "" {
		panic("no value for $USER")
	} else {
		fmt.Printf("User=%v\n", user)
	}

	fmt.Println("Exit initMain")
}

func copyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer func() {
		src.Close()
		fmt.Printf("close src\n")
	}()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer func() {
		dst.Close()
		fmt.Printf("close dst\n")
	}()

	return io.Copy(dst, src)
}

func main() {
	copyFile("/home/peter/a.out", "a.in")
	fmt.Println("Exit main")
}
