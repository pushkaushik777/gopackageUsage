package main

import "fmt"

func main() {
	fmt.Println("In the main program")

	var org [5]int = [5]int{1, 4242, 5555, 43324, 666}
	var max int
	for i, elem1 := range org {
		if i == 4 {
			break
		}
		fmt.Println("Max value comparison b/n:", elem1, org[i+1], max)

		if elem1 > org[i+1] && elem1 >= max {
			max = elem1
		} else {
			if org[i+1] > max {
				max = org[i+1]
			}
		}
		fmt.Println("Max value is :", max)
	}
	fmt.Println("Max value is :", max)
}

