package main

import "fmt"

func main() {
	fmt.Println(1, 2, 3)

	var (
		x int
		y int
	)

	fmt.Print("輸入兩個數字，用空格隔開:")
	fmt.Scanln(&x, &y)
	var result int = x + y
	fmt.Println("相加結果:", result)
}
