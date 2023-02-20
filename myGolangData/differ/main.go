package main

import "fmt"

/// difer usees lifo last in fist out
func main() {
	defer fmt.Println("using differ")
	defer fmt.Println("One two")
	defer fmt.Println("Three four")
	defer fmt.Println("Five six")

	fmt.Println("after difer")
	myDiffer()
}

func myDiffer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
