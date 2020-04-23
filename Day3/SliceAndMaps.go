package main

import "fmt"

var slizzee =[]int{2,3,4,5,3}
func main()  {

	//Slice Manipulations
fmt.Println("------------Slice manipulation----------------")
	fmt.Println("Length of slice: ",len(slizzee)) //prints length

	slizeeUpdate(344,2) //update values

	var slizze2 []int = slizzee
	fmt.Println("Slice1",slizzee,"   Slice2",slizze2) //yes because slices are work with pointers
	fmt.Println("---------------------------------------------")
	//Map Manipulations
	fmt.Println("------------Map manipulation----------------")
	studentRecord := make(map[int]string)
	studentRecord[01] = "Devesh"
	studentRecord[02]="Niranjan"

	fmt.Println("Before modifications")
	for k,v := range studentRecord{
		fmt.Println("Roll No",k,"Name: ",v)
	}

	fmt.Println("After modification")
	studentRecord[02]="Tushar"

	for k,v := range studentRecord{
		fmt.Println("Roll No",k,"Name: ",v)
	}

	}
//Function to update size
func slizeeUpdate(value int, index int) {

	slizzee[index]=value
	fmt.Println("After update of slice ",slizzee)
}

