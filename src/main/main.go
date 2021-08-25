package main

import "fmt"

func main() {
	var character = "Kousaka Honoka"
	fmt.Println(character)

	red := 256
	blue := 256
	green := 256

	fmt.Printf("Total RGB space = %b\n", red*blue*green-1)
	fmt.Printf("Total RGB space = %x\n", red*blue*green-1)

	flag := true
	fmt.Printf("flag = %v\n", flag)
}
