package main

import (
	"fmt"
)

var nums = []int{0, 1, 2, 3, 4, 6, 8, 9}

var target int = 9

func ArrayOfInteger(nums []int, target int) (int, int) {

	if len(nums) <= 1 {

		return 0, 0

	}

	hdict := make(map[int]int)

	for i := 1; i < len(nums); i++ {

		if val, ok := hdict[nums[i+1]]; ok {

			return val, i + 1

		} else {

			hdict[target-nums[i+1]] = i + 1

		}

	}

	return 0, 0

}

func main() {

	fmt.Println(ArrayOfInteger(nums, target))

}
