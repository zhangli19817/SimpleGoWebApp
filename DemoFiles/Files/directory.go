package Files

import (
	"fmt"
	"io/ioutil"
)

func ListDir(dir string) ([]string, error) {

	var dirs []string
	dirList, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	} else {
		var number int = 0
		for _, value := range dirList {
			if value.IsDir() {
				//fmt.Printf("%d:[name:%s]\n",number,value.Name())
				dirs = append(dirs, value.Name())
				number++
			}
		}
		fmt.Printf("###Totally %d directories under this folder.###\n", number)
		//fmt.Println(dirs)
	}

	return dirs, nil
}

func ListAllFiles(dir string) (int, []string, error) {

	var files []string

	fileList, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Failed to read dir", err)
		return 0, nil, err
	}

	var total int = 0

	for _, value := range fileList {
		if value.IsDir() {
			//fmt.Printf("Now scanning directory %s\n", value.Name())
			num, lists, _ := ListAllFiles(dir + "/" + value.Name())
			for _, file := range lists {
				files = append(files, file)
			}
			//fmt.Printf("File %s was scanned\n", value.Name())
			total += num
		} else if value.Mode().IsRegular() {
			total += 1
		}
	}

	return total, files, err
}
