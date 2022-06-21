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
	write(&balance, calc(&balance, amount))
	fmt.Printf("new balance: %d\n", balance)
}

func main() {
	go deposit(10)
	go deposit(10)
	go deposit(10)
	go deposit(10)
	go deposit(10)
	time.Sleep(time.Second * 1)

}
