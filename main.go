package main

import (
	"fmt"
)

func main() {
var num1 int64
var num2 int64
var num3 int64
var num4 int64
var num5 int64
var num6 int64
var num7 int64
var num8 int64
num1 = 0
num2 = 0
num3 = 0
num4 = 0
num5 = 0
num6 = 0
num7 = 0
num8 = 0
for {
	num1 = num1 + 1
	fmt.Println(num8, num7, num6, num5, num4, num3, num2, num1)
	if num1 == 99999999999999999 || num1 < 0 {
		num2 = num2 + 1
		num1 = 0
	}
	if num2 == 99999999999999999 || num2 < 0 {
		num3 = num3 + 1
		num2 = 0
	}
	if num3 == 99999999999999999 || num3 < 0 {
		num4 = num4 + 1
		num3 = 0
	}
	if num4 == 99999999999999999 || num4 < 0 {
		num5 = num5 + 1
		num4 = 0
	}
	if num5 == 99999999999999999 || num5 < 0 {
		num6 = num6 + 1
		num5 = 0
	}
	if num6 == 99999999999999999 || num6 < 0 {
		num7 = num7 + 1
		num6 = 0
	}
	if num7 == 99999999999999999 || num7 < 0 {
		num8 = num8 + 1
		num7 = 0
	}
	
}
}
