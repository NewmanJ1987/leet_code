package main

import (
	"fmt"

	"github.com/newmanJ1987/leet_code/problems"
)

var stack Stack

type Stack struct {
	items []string
}

func main() {
	// test:= "hello"
	// problems.Push(test)
	problems.InitStack()
	fmt.Println(problems.ValidateParantheses("{[]}"))
	// m := make(map[string]int)
	// m["I"] = 1
	// m["V"] = 5
	// m["X"] = 10
	// m["L"] = 50
	// m["C"] = 100
	// m["D"] = 500
	// m["M"] = 1000

	// input := "LXIV"

	// fmt.Println(problems.EfficentAlgorithmRecursive(m, input, "M", 0))
	// fmt.Println(problems.LongestCommonPrefix(([]string{"apple", "app", "appium"})))

}
