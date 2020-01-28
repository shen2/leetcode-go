package main

import (
    "fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int)  {
    i:=m-1
    j:=n-1
    for k:= m + n -1; k >= 0; k--{
        if (i < 0 || j >= 0 && nums1[i] < nums2[j]){
            nums1[k] = nums2[j]
            j--
        }else{
            nums1[k] = nums1[i]
            i--
        }
    }
}

func main() {
    merge([]int {1,2,3,0,0,0}, 3, []int {2,5,6}, 3)
}