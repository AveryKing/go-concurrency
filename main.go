package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	evilNinjas := []string{"Adam", "Avery", "Tom", "Ryan"}

	for _, evilNinja := range evilNinjas {
		go attack(evilNinja)
	}

	time.Sleep(time.Second)
}

func attack(target string) {
	fmt.Println("Throwing ninja stars at ", target)
	time.Sleep(time.Second)
}
