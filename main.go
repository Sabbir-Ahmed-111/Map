package main

import "fmt"

func main(){

	 fruits := map[string]int{
        "apple":  10,
        "banana": 5,
        "orange": 7,
	 }

	// value add
   fruits["mango"] = 20

	fmt.Println(fruits)

	// value delete

	delete(fruits, "banana")

	fmt.Println(fruits)
}