package main

import "fmt"

func main() {
	student := Student{"Ashish", 12, "Army school", "Raturi"}
	student.getStudent()
	student.NewCast()
	fmt.Println(student.Cast)

}

type Student struct {
	Name   string
	Age    int
	School string
	Cast   string
}

func (u Student) getStudent() {
	fmt.Println("student cast is: ", u.Cast)

}

func (u Student) NewCast() {
	u.Cast = "Semwal"
	fmt.Println("cast is : ", u.Cast)
}
