package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Create("userInput.txt")
	if err != nil {
		fmt.Println("file open error: ", err)
		return
	}

	defer file.Close()

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		input := scan.Text()

		if input == "" {
			break
		}

		_, err := file.WriteString(input + "\n")
		if err != nil {
			fmt.Println("file writing error: ", err)
			return
		}
	}

	fileData, err := ioutil.ReadFile("userInput.txt")
	if err != nil {
		fmt.Println("file reading error: ", err)
		return
	}
	fmt.Printf("file contents:\n%v\n", string(fileData))

}
