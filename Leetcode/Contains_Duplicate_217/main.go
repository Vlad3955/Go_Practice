package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
    if len(nums) <= 1 {
		return false
	}

	um := make(map[int]struct{})

	for _, v := range nums {

		if _, ok := um[v]; ok {
			return true
		}

		um[v] = struct{}{}
	}

	return false

}

func main()  {
	nums := []int{1,2,3,1}

	fmt.Println(containsDuplicate(nums))
}