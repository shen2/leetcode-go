package main

import (
    "fmt"
)

func romanToInt(str string) int {
    sum := 0

    for i:=0; i < len(str); i++{
        switch string(str[i]) {
        case "I":
            sum += 1
        case "V":
            if i > 0 && string(str[i-1]) == "I" {
                sum += 3
            }else{
                sum += 5
            }
        case "X":
            if i > 0 && string(str[i-1]) == "I" {
                sum += 8
            }else{
                sum += 10
            }
        case "L":
            if i > 0 && string(str[i-1]) == "X" {
                sum += 30
            }else{
                sum += 50
            }
        case "C":
            if i > 0 && string(str[i-1]) == "X" {
                sum += 80
            }else{
                sum += 100
            }
        case "D":
            if i > 0 && string(str[i-1]) == "C" {
                sum += 300
            }else{
                sum += 500
            }
        case "M":
            if i > 0 && string(str[i-1]) == "C" {
                sum += 800
            }else{
                sum += 1000
            }
        }
    }
    return sum
}

func main() {
    fmt.Println(romanToInt("CM"))
}