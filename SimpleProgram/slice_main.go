package main

import "fmt"

func main() {
	fmt.Println("In the main program")
	//Creating Sliceas j
	var j []int = []int{21, 3, 5, 6}
	//Creating Another Slice just copying it..
	k := j

	//Now modifying value in copied slice
	k[2] = 3333

	fmt.Println("Now Slices after change")
	fmt.Println("Original Slice", j)  //Original Slice [21 3 3333 6]
	fmt.Println("Updated Slice", k) //Updated Slice [21 3 3333 6]
		
}
