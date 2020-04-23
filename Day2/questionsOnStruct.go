package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type contactInfo struct {
	mobileNum int
}

type student struct {
	stuName string
	contact contactInfo //added stcure as field
}

func updateStudent(name *string, num *int) {

	res := student{
		stuName: *name, //using pointers
		contact: contactInfo{*num},
	}
	fmt.Println("Updated values :", res)

}
func main() {

	wholeStruct := student{stuName: "devesh", contact: contactInfo{9131725221}}
	fmt.Println(wholeStruct) //prinint whole stuct

	newName := "DEVESHTAMHANE"
	newNum := 89281928
	updateStudent(&newName, &newNum)

	f, err := os.Create("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	b, _ := json.Marshal(wholeStruct)
	l, err := f.WriteString(string(b))
	fmt.Println(l)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}
