package main

import "fmt"

func main()  {
	studentRecord := make(map[int]string)
	studentRecord[01] = "Devesh"
	studentRecord[02] = "Niranjan"

	fmt.Println(" Before modifying Entries : ")

	for k,v := range studentRecord {
		fmt.Print(" \n Roll number is : ", k," Student name is : ",v)

	}

	fmt.Println("\n After deleting Entries : ")

	for k,v:= range studentRecord {
		fmt.Print(" \n Roll number is : ", k," Student name is : ",v)
	}

	defer fmt.Println("\n Earlier not able to delete Entries because defer is used. Now Student Entries are : ", studentRecord)

	defer delete(studentRecord, 01)

}
