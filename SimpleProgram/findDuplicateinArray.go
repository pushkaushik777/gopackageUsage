	package main

	import "fmt"

	func main() {
		fmt.Println("In the main program")

		var org [5]int = [5]int{1, 4, 5, 4, 5}

		for i, elem1 := range org {
			//fmt.Println("First run .... ", frstRun, elem1)
			for k, elem2 := range org {
				if elem1 == elem2 && k < i {
					fmt.Println("Duplicate Element : ", elem2)
				}
				//fmt.Println("Second run .... ", scndRun, elem2)

			}
		}

	}

