package main

import (
	"fmt"
	"time"
)

var slc = make([]string, 0)

func fun1() {
	slc = append(slc, "x")
}

func fun2() {
	slc = append(slc, "o")
}

func main() {
	go fun1()
	go fun1()
	go fun1()
	go fun1()
	go fun2()
	go fun2()
	go fun2()
	go fun2()
	time.Sleep(time.Second * 5)
	for _, x := range slc {
		fmt.Println(x)
	}
}
