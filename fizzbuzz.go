package main

import "fmt"

func main() {
	fmt.Print("Enter limit: ")
	var limit int
	fmt.Scanf("%d", &limit)
	
	for i := 1; i <= limit; i++ {
		if i % 3 == 0 {
			if i % 5 == 0 {
				fmt.Println(i, "fizzbuzz")
			} else {
				fmt.Println(i, "fizz")
			}
		} else if i % 5 == 0 {
			fmt.Println(i, "bazz")
		} else {
			fmt.Println(i)
		}
	}
}
