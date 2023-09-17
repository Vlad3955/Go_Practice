package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	um := make(map[int]int)

	// for i := 0; i < len(nums); i++ {
	// 	compl := target - nums[i]
	// 	if comp, ok := um[compl]; ok {
	// 		return []int{comp, i}
	// 	}
    //     um[nums[i]] = i
	// }
    // return []int{}

	for idx, num := range nums {
		if val, found := um[target - num]; found {
			return []int{val, idx}
		}
		um[num] = idx
	} 
	return nil
}


func main()  {
	nums := []int{3,3}
	target := 6

	fmt.Println(twoSum(nums, target))
}