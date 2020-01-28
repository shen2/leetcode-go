package main

import (
    "fmt"
)

func isPair(a,b byte) bool{
	switch string(a){
	case "{":
		if string(b) == "}"{
			return true
		}
	case "[":
		if string(b) == "]"{
			return true
		}
	case "(":
		if string(b) == ")"{
			return true
		}
	default:
		return false
	}
	return false
}


func isValid(s string) bool {
	var stack []byte
    for i:=0; i < len(s); i++{
    	if string(s[i]) == "(" || string(s[i]) == "[" || string(s[i]) == "{"{
    		stack = append(stack, s[i])
    	}else{

    		if len(stack) > 0 && isPair(stack[len(stack) - 1], s[i]){
    			stack = stack[0:len(stack) -1]
	    	}else{
	    		return false
	    	}
	    }
    }
    return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()[{]}"))
}