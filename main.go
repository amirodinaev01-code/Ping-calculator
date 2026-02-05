package main

import (
	"fmt"
)

func main(){
	var myPing int16 = 145

	fmt.Printf("Current ping: %dms\n", myPing)	

	switch {
	case myPing <= 50:
		fmt.Println("your ping is very good")
	case myPing <=100:
		fmt.Println("your ping is nice")
	case myPing <= 150:
		fmt.Println("Tajikistan, you need for Starlnik")
	default:
		fmt.Println("plise delete the game")
	}
}