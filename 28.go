package main

import (
    "fmt"
)

func strStr(haystack string, needle string) int {
    for i:=0; i < len(haystack) - len(needle) + 1; i++{
        match := true
        for j:=0; j < len(needle); j++{
            if haystack[i+j] != needle[j]{
                match = false
                break
            }
        }
        if match{
            return i
        }
    }
    return -1
}

func main() {
    fmt.Println(strStr("aaaaa", "bba"))
}