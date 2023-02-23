package main

import "fmt"

// [1] Concept of Iota
// for each constant block `denoted with const ()`
// iota will count values from 0
// if it's used in another block it will start again from 0
const (
	first = iota
	second
)

const (
	third = iota
	fourth
)

const (
	firstExpr  = iota + 6
	secondExpr = iota << 1 // byte shift by 1 it will basically multiply by 2
)

// constant expresion used with Iota will be used for each initialization till reassigned
const (
	firstExpr2 = iota + 6
	secondExpr2
	thirdExpr2 = iota - 6
)

func main() {
	fmt.Println(first, second, third, fourth)
	fmt.Println(firstExpr, secondExpr)
	fmt.Println(firstExpr2, secondExpr2, thirdExpr2)
}
