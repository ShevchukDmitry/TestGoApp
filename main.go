package main

import "fmt"

func main(){

// create variables
var one string 
one = "One" // kyeword var + type
two := "Two" // 
three := "Three"
var four = "Four" // keyword + init

fmt.Printf("%s %s %s %s \n", one, two, three, four)
fmt.Println(RetNums(one))

// slice
nums := []string{one,two,three,four}
fmt.Println(nums)
nums = append(nums, "Five")
fmt.Println(nums)

for t, nums := range nums {
	fmt.Println(t, nums)
	}
}

func RetNums(num string) string{
	return "We have: "+num
}


