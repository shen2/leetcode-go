package main

import (
    "fmt"
)

func removeElement(nums []int, val int) int {
    j:=0
    for i:=0; i < len(nums); i++{
        if nums[i] != val{
            nums[j] = nums[i]
            j++
        }
    }
    return j
}

func main() {
    fmt.Println(removeElement([]int {0,0,1,1,1,2,2,3,3,4}, 4))
}