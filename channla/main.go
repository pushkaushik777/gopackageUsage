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
}
