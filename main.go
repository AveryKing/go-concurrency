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

	smokeSignal := make(chan bool)

	for _, evilNinja := range evilNinjas {
		go attack(evilNinja, smokeSignal)
	}

	fmt.Println(<-smokeSignal)
}

func attack(target string, defeated chan bool) {
	fmt.Println("Throwing ninja stars at ", target)
	defeated <- true
}
