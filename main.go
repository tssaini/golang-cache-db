package main

import (
	"time"
	"math/rand"
	"fmt"
	"sync"
)

var cache map[int]Book
var seed = rand.NewSource(time.Now().UnixNano())
var r = rand.New(seed)

func main() {
	cache = make(map[int]Book)
	wg := &sync.WaitGroup{}
	rwMutex := &sync.RWMutex{}

	for i:=0; i< 10; i++{
		randID := r.Intn(10)+1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, rwMutex *sync.RWMutex){
			rwMutex.RLock()
			b, ok := queryCache(id)
			rwMutex.RUnlock()
			if ok {
				fmt.Println("From cache")
				fmt.Println(b)
			}
			wg.Done()
		}(randID, wg, rwMutex)

		go func(id int, wg *sync.WaitGroup, rwMutex *sync.RWMutex){
			b, ok := queryDB(id)
			if ok {
				rwMutex.Lock()
				cache[id] = b
				rwMutex.Unlock()
				fmt.Println("From database")
				fmt.Println(b)
			}
			wg.Done()
		}(randID, wg, rwMutex)
		time.Sleep(150 * time.Millisecond)
		// fmt.Printf("Book with ID %v not found\n", randID)
	}
	wg.Wait()
}


func queryCache(id int) (Book, bool) {
	val, ok := cache[id]
	if ok{
		return val, true
	}
	return Book{}, false
}

func queryDB(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books{
		if b.ID == id{
			return b, true
		}
	}
	return Book{}, false
}
