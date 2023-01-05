package main

import (
	"fmt"
)

func main() {

	var m map[string]int = map[string]int{
		"apples":  5,
		"pears":   7,
		"oranges": 8,
	}

// check  if  a value is in a map
   val, ok := m["apples"]
   fmt.Println(val , ok)

	//access a map
       fmt.Println(m["apples"])
	   //add in a map
	   m["avocado"] = 900
	   //delete in a map
	   delete(m , "pears")
	fmt.Println(m)
}
