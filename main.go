package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type TV struct {
	Text string
	Var  interface{}
}

// printNamedValueAndType
func pvtt(nv TV) {
	fmt.Println("__________________________________________________")
	fmt.Printf("text: %s\nvalue: %v\ntype: %T\n", nv.Text, nv.Var, nv.Var)
	fmt.Println("__________________________________________________")
}

// printValueAndType example: value: 5 type: int
func pvt(v interface{}) {
	fmt.Println("__________________________________________________")
	fmt.Printf("value: %v\ntype: %T\n", v, v)
	fmt.Println("__________________________________________________")

}

//	func main() {
//		file, err := os.Create("userInput.txt")
//		if err != nil {
//			fmt.Println("file open error: ", err)
//			return
//		}
//
//		defer file.Close()
//
//		scan := bufio.NewScanner(os.Stdin)
//		for scan.Scan() {
//			input := scan.Text()
//
//			if input == "" {
//				break
//			}
//
//			_, err := file.WriteString(input + "\n")
//			if err != nil {
//				fmt.Println("file writing error: ", err)
//				return
//			}
//		}
//
//		fileData, err := ioutil.ReadFile("userInput.txt")
//		if err != nil {
//			fmt.Println("file reading error: ", err)
//			return
//		}
//		fmt.Printf("file contents:\n%v\n", string(fileData))
//
// }
func main() {
	fileData, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Printf("read file error:\n%v\n")
	}
	fileString := string(fileData)
	fmt.Printf("data:\n%v\n", string(fileData))

	fileStringUpper := strings.ToUpper(fileString)

	file, err := os.Create("data2.txt")
	if err != nil {
		fmt.Printf("create file error:\n%v\n")
	}
	defer file.Close()

	_, err = file.WriteString(fileStringUpper)
	if err != nil {
		fmt.Printf("write file error:\n%v\n")
	}

	fileData, err = ioutil.ReadFile("data2.txt")
	if err != nil {
		fmt.Printf("read file error:\n%v\n")
	}
	fmt.Printf("data2:\n%v\n", string(fileData))
}
