package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	printCurrentTime()

	//go printChar("a",10)
	//go printChar("b",10)

	//time.Sleep(5*time.Second)

	//Demo create a new channel
	//channelA:=make(chan float64)
	//channelB:=make(chan int)
	//channelC:=make(chan bool)

	//channelA <- 3.14
	//channelB <- 100
	//channelC <- true

	/*message:=make(chan string)
	go sayHello(message,"good morning")
	receivedMessage:=<-message
	fmt.Println(receivedMessage)*/

	//channelA:=make(chan string)
	//channelB:=make(chan string)

	//go abc(channelA)
	//go def(channelB)

	//should print out
	//a
	//d
	//b
	//e
	//c
	//f
	/*fmt.Println(<-channelA)
	fmt.Println(<-channelB)
	fmt.Println(<-channelA)
	fmt.Println(<-channelB)
	fmt.Println(<-channelA)
	fmt.Println(<-channelB)*/

	/*channel:=make(chan string)
	go reportNG(channel,3)
	waitforTime("received message sleeping",1)
	fmt.Println(<-channel)
	fmt.Println(<-channel)*/

	//channel:=make(chan string)
	//go deadlock(channel)
	//fmt.Println(<-channel)

	channel := make(chan page)

	urls := []string{"http://ec2-18-141-10-109.ap-southeast-1.compute.amazonaws.com",
		"http://www.qq.com", "http://www.oracle.com",
	}

	for _, value := range urls {
		go printWebPageSize(channel, value)
	}

	for i := 0; i < len(urls); i++ {
		pages := <-channel
		//fmt.Println(pages.url,pages.size)
		fmt.Printf("%d:\"%s\": [%d] -%d\n", i+1, pages.url, pages.size, pages.statusCode)
	}

	printCurrentTime()
}

type page struct {
	url        string
	size       int
	statusCode int
}

func printWebPageSize(channel chan page, targetUrl string) {

	//fmt.Println(targetUrl)
	response, err := http.Get(targetUrl)
	responseCode := response.StatusCode
	if err == nil {
		//fmt.Println(response.StatusCode)
		//fmt.Println(response.Status)
	} else {
		log.Fatal(err)
	}

	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	} else {
		//fmt.Println(string(content))
		//fmt.Println(len(content))
	}

	channel <- page{url: targetUrl, size: len(content), statusCode: responseCode}
}

func printCurrentTime() {

	current := time.Now()
	fmt.Println(current.Year(), current.Day(), current.Month(), current.Hour(), current.Minute(), current.Second(), current.Nanosecond())

}

func printChar(s string, times int) {
	for i := 1; i <= times; i++ {
		fmt.Print(s)
	}
}

func sayHello(message chan string, word string) {
	message <- word
}

func abc(channel chan string) {
	channel <- "a"
	//fmt.Println("received message","a")
	channel <- "b"
	//fmt.Println("received message","b")
	channel <- "c"
	//fmt.Println("received message","c")
}

func def(channel chan string) {
	channel <- "d"
	//fmt.Println("received message","d")
	channel <- "e"
	//fmt.Println("received message","e")
	channel <- "f"
	//fmt.Println("received message","f")
}

func waitforTime(name string, times int) {
	for i := 0; i < times; i++ {
		fmt.Printf("%s is sleeping\n", name)
		time.Sleep(1 * time.Second)
	}
}

func reportNG(channel chan string, times int) {
	waitforTime("sending message sleeping", times)
	fmt.Println("sending values.......")
	channel <- "a"
	fmt.Println("sent out values.......")
	channel <- "b"
	fmt.Println("sent out values.......")

}

func deadlock(message chan string) {

}
