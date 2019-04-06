package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	for index1, num1 := range nums {
		for index2, num2 := range nums {
			if target == num1+num2 && index1 != index2 {
				return []int{index1, index2}
			}
		}
	}
	return nums
}
func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
