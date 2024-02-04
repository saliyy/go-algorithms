package main

/*
https://leetcode.com/problems/add-digits/
*/
import "fmt"

func main() {
	res := addDigits(38)
	fmt.Println(res)
}

func divMod(x, y int) (int, int) {
	return x / y, x % y
}

func addDigits(num int) int {
	for num > 10 {
		cur := num
		new_num := 0
		for cur != 0 {
			d := 0
			cur, d = divMod(cur, 10)
			new_num += d
		}
		num = new_num
	}
	return num
}
