package main

import (
	"fmt"
)

func Binary(n int) int {

	count := 0

	for i := 0; i < len(n); i++ {

		if n & (1 << i) {

			fmt.Println("inside if: ")

			count++

		}

	}

	return count

	//return (n & 1) + countBitofOne(n>>1)

}

func main() {

	i := 4

	fmt.Println("size is ")

	// n := int64(i)

	//fmt.Sprintf("%b", big.NewInt(4))

	fmt.Println("count bits of 1 : ", Binary(i))

}
