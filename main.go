package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Println("error open file", err)
		return
	}
	fmt.Printf("content from file: \n%v\n", string(content))
}
