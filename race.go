package main

import (
	"fmt"
	"time"
)

var balance = 0

func calc(toModify *int, amount int) int {
	return *toModify + amount
}

func write(toWrite *int, value int) {
	*toWrite += value
}

func deposit(amount int) {
	newVal := calc(&balance, amount)
	go write(&balance, newVal)
	fmt.Printf("new balance: %d\n", balance)
}

func race() {
	go deposit(10)
	go deposit(10)
	go deposit(10)
	go deposit(10)
	time.Sleep(time.Second * 1)
	fmt.Println(balance)
}
