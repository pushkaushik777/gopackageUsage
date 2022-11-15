package main

import (
	"fmt"

	"github.com/newGo/channl"
)

func main() {
	fmt.Println("This is the result valure")
	g := channl.GetSum(2)
	fmt.Println("Value is ", g)
	fmt.Println("Value is ", channl.T)

	p := channl.Piyush{Name: "Kaushik", Company: "HCL", Salary: 10}
	fmt.Println(p.Name)
	fmt.Println(p.Company)
	fmt.Println(p.Salary)

	//Testing Slices & Array
	var arr [7]int = [7]int{12, 33, 34, 5, 2, 3, 2}
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
	var s []int = []int{3: 3}
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

}
