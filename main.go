package main

import (
	"time"
	"math/rand"
	"fmt"
)

var cache map[int]Book
var seed = rand.NewSource(time.Now().UnixNano())
var r = rand.New(seed)

func main() {
	cache = make(map[int]Book)
	for i:=0; i< 10; i++{
		randID := r.Intn(10)+1

		b, _ := queryDB(randID)
		fmt.Println(b)
	}
}


func queryCache(id int) (Book, bool) {
	val, ok := cache[id]
	if ok{
		return val, true
	}
	return Book{}, false
}

func queryDB(id int) (Book, bool) {
	for _, b := range books{
		if b.ID == id{
			return b, true
		}
	}
	return Book{}, false
}
