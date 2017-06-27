package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	//for i := 0; i < 5; i++ {
	for {
		//fmt.Println("ping... ")
		msg := "ping"
		c <- msg
		//fmt.Println(msg)
		//time.Sleep(time.Second * 1)
	}
}
func printer(c chan string) {
	for {
		//fmt.Println("printer... ")
		msg := <-c
		fmt.Println(msg, "<-", time.Now())
		time.Sleep(time.Second * 1)
	}
}
func main1() {
	var c = make(chan string, 3)
	go printer(c)
	go pinger(c)
	var input string
	fmt.Scanln(&input)
	fmt.Println("ending main... ")
}
