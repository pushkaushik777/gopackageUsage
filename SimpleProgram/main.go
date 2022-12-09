package main

import (
	"fmt"
)

func main() {
	fmt.Println("In the main PAckage")

	var array = []int{1, 4, 6, 8, -4, 2}
	Print(array)
	var t string = "Hello"
	Print(t)

	//out, err := exec.Command("pwd").Output()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(out[:]))

}

func Print(x interface{}) {
	switch y := x.(type) {
	case string:
		fmt.Println("string", y)
	case int:
		fmt.Println("Integer ", y)
	case []int:
		fmt.Println("Slice ", x)
	default:
		fmt.Println("Type not matches")
	}
}
