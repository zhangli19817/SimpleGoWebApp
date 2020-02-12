package main

import (
	"awesomeProject/DemoFiles/Files"
	"fmt"
	"log"
)

func main() {
	file := "/Users/luffy/Downloads/filesDemoList"
	//if len(os.Args)<2{
	//	log.Fatal("Pls give a valid file name...")
	//}
	//fileList, err := Files.GetFileList(os.Args[1])
	fileList, err := Files.GetFileList(file)
	if err != nil {
		fmt.Println("Err Msg:", err)
	} else {
		for index, value := range fileList {
			fmt.Printf("%d:[%s]\n", index+1, value)
		}
	}

	//Demo list Directory
	dirs, err1 := Files.ListDir("/Users/luffy/Downloads")
	if err1 == nil {
		for index, value := range dirs {
			fmt.Printf("%d:[\"%s\"]\n", index+1, value)
		}
	} else {
		log.Fatal(err1)
	}

	fmt.Println("---------------------------------")

	total, files, err2 := Files.ListAllFiles("/Users/luffy/Downloads")
	if err2 == nil {
		fmt.Printf("Totally this directory has %d files\n", total)
		for index, file := range files {
			fmt.Printf("%d:[\"%s\"]\n", index+1, file)
		}
	} else {
		log.Fatal(err2)
	}
}
