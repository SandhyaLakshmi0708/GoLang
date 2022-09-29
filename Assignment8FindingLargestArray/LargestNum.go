package main

import "fmt"

func LargestNum() {

	a := []int{1, 2, 3, 5, 6, 9, 10}
	var temp int

	for i := 0; i < len(a); i++ {

		for j := i; j < len(a); j++ {
			if a[i] < a[j] {
				temp = a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}

	}
	fmt.Print("the largest number is:", a[0])

}
