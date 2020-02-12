package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"path/filepath"
)

func main() {

	//sayHello("Welcome to join Google Inc...")
	showNumber(1, 2)
	err, totalNUmber := printFiles("/Users/luffy/DemoFiles")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The above directory has %d files\n", totalNUmber)

	//fmt.Println(string(os.PathSeparator))

	//panic("oh my god, Go panic happened!!! Fuck you!!")
	for i := 1; i <= 10; i++ {
		fmt.Print(i, ":")
		awardPrize()
	}

	//fmt.Println(recover())
	//Demo Go recover built-in function
	defer demoRecover()
	panic(p)

}

func sayHello(message string) {
	fmt.Println(message)
	sayHello(message)
}

func showNumber(start int, end int) {
	fmt.Printf("%d:%d\n", start, end)
	if start < end {
		showNumber(start+1, end)
	}
	fmt.Printf("calling showNumber(%d,%d)returned\n", start, end)

	//panic("oh my god, Go panic happened!!! Fuck you!!")
}

func printFiles(path string) (error, int) {

	items, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("Error:scanning directory \"%s\" was failed\n", path)
		log.Fatal(err)
	}

	total := 0

	for _, value := range items {

		if value.IsDir() {
			currentPath := filepath.Join(path, value.Name())
			err, number := printFiles(currentPath)
			if err != nil {
				fmt.Printf("Error:scanning directory \"%s\" was failed\n", currentPath)
				log.Fatal(err)
			}
			total += number
		} else {
			if value.Mode().IsRegular() {
				fmt.Println(path, value.Name())
				total += 1
			}

		}
	}

	//fmt.Printf("This directort totally has %d files",total)

	return nil, total
}

func awardPrize() {

	value := rand.Intn(3) + 1

	if value == 1 {
		fmt.Println("Congratulations! You win the prize!!")
	} else if value == 2 {
		fmt.Println("good,try again...")
	} else if value == 3 {
		fmt.Println("closed to the correct value, try again...")
	} else {
		panic("not a valid value...")
	}

}

func demoPanicWithRecover() string {
	defer fmt.Println("exited normally!!")
	defer demoRecover()
	panic("throw panic error message!!!")

}

func demoRecover() {
	recover()
}
