package main

import (
    "fmt"
    "time"
)

var timeout = 3 * time.Minute

type coffee string

var maxExpresso coffee = "Expresso double double shot"

func main(){
	cup := make(chan coffee)
	go gonsterBuyCoffee(cup)
	waitForCoffee(cup)
}

func waitForCoffee(cup chan coffee){
	select {
	case coffee := <-cup:
		fmt.Printf("Receive coffee \"%s\"\n", coffee)
	case <- time.After(3 * time.Second):
		fmt.Printf("Time out")
	}
}

func gonsterBuyCoffee(cup chan coffee) {
	time.Sleep(timeout)
	cup <- maxExpresso
}

