package Files

import (
	"bufio"
	"fmt"
	"os"
)

func OpenFile(fileName string) (*os.File, error) {
	fmt.Println("Opening file...>>", fileName)
	return os.Open(fileName)
}

func CloseFile(file *os.File) {
	fmt.Println("Closing Files...>>")
	file.Close()
}

func GetFileList(fileName string) ([]string, error) {
	var numbers []string
	file, err := OpenFile(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var value string = scanner.Text()
		//fmt.Println(value)
		//value = "\"" + strings.TrimSpace(value) + "\""
		value = "\"" + value + "\""
		numbers = append(numbers, value)
	}

	//fmt.Println(numbers)

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	defer CloseFile(file)
	return numbers, nil
}

func SocialMedia() error {

	defer fmt.Println("Goodbye!")
	fmt.Println("Hello World!")
	return fmt.Errorf("I don't want to talk anymore!!!")

}
