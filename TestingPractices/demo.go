package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime()
	channel := make(chan int)
	go calNumbers(channel, 0, 100000000000)
	go calNumbers(channel, 0, 100000000000)
	go calNumbers(channel, 0, 100000000000)

	var total int = 0
	total += <-channel
	total += <-channel
	total += <-channel

	fmt.Printf("with go routines:The total number is    %d\n", total)
	currentTime()

	//fmt.Println(calNumbers1(1,100))
	value := calNumbers1(0, 100000000000)
	value += calNumbers1(0, 100000000000)
	value += calNumbers1(0, 100000000000)
	fmt.Printf("without go routines:The total number is %d\n", value)
	currentTime()

}

func calNumbers(channel chan int, start int, end int) {
	total := 0
	i := 0
	i = start
	for i <= end {
		total += i
		i++
	}

	channel <- total
}

func calNumbers1(start int, end int) int {
	total := 0
	i := 0
	i = start
	for i <= end {
		total += i
		i++
	}

	return total
}

func currentTime() {

	current := time.Now()
	fmt.Println(current.Hour(), current.Minute(), current.Second())
}
