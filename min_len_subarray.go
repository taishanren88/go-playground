package main

import (
		"fmt"
		"math"
)

// Max returns the larger of x or y.
func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
    if x > y {
        return y
    }
    return x
}

func minSubArrayLen(s int, nums []int) int {
	// Sliding window  problem
	// Loop from "right" : 0 : n
	// if sum is < s , keep moving "right"
	// Once it it is >= , keep incrementing left and update minimum size (under same loop)
	// Once left is < , move to the next iteration of the loop
	var n int = len(nums)
	var sum int = 0
	var left int = 0
	var len int = math.MaxInt32;
	for right:= 0 ; right < n; right++ {
		if sum < s {
			sum += nums[right];
		}
		for sum >= s {
			len = Min(right - left + 1, len)
			sum -= nums[left];
			left ++;
		}
	}
    if len == math.MaxInt32 {
    	return 0
    } else  {
    	return len;
    }
}

func main() {
	var TestSlice = []int {2, 3,1,2,4,3}
	fmt.Println("done")
	fmt.Println(minSubArrayLen(7, TestSlice))
}