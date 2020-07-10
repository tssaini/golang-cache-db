package main

import (
	"fmt"
)

var cache map[int]Book


func main() {
	cache = make(map[int]Book)
	for _, b := range books{
		fmt.Println(b)
	}

}


