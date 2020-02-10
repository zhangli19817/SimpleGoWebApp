package main

import (
	"awesomeProject/DemoFiles/Files"
	"fmt"
)

func main() {
	file := "/Users/luffy/Downloads/filesDemoList"
	_, err := Files.GetFileList(file)
	if err != nil {
		fmt.Println("Err Msg:", err)
	}

}
