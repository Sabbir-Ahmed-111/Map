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

	//Map থেকে ডাটা রিড করা
	V := fruits["apple"]
    fmt.Println(V) // Output: 10


	// Map-এ কোনো key আছে কিনা চেক করা

	 val, ok := fruits["banana"]
        if ok {
          fmt.Println("Found:",val)
        } else {
          fmt.Println("Not found")
        }

		//Map এর উপর loop চালানো

		for key, value := range fruits {
           fmt.Println(key, value)
           }

	
	 //Map এর length জানা

	   fmt.Println(len(fruits))
	   



}