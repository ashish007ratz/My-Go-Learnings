package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to file in go lang")
	content := "this is a file data and to test demo of a project"
	file, err := os.Create("./myfirstfile.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("lenght is: ", length)
	defer file.Close()
	readFile("./myfirstfile.txt")

}

func readFile(fileName string) {
	dataFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Println("data bytes in a file is below: \n", dataFile)
    fmt.Println("data in a file is below: \n", string(dataFile))
}
