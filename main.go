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
	ch1 := make(chan Book)
	ch2 := make(chan Book)

	for i:=0; i< 10; i++{
		randID := r.Intn(10)+1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, rwMutex *sync.RWMutex, ch chan<- Book){
			rwMutex.RLock()
			b, ok := queryCache(id)
			rwMutex.RUnlock()
			if ok {
				ch <- b
			}
			wg.Done()
		}(randID, wg, rwMutex, ch1)

		go func(id int, wg *sync.WaitGroup, rwMutex *sync.RWMutex, ch chan<- Book){
			b, ok := queryDB(id)
			if ok {
				rwMutex.Lock()
				cache[id] = b
				rwMutex.Unlock()
				ch <- b
			}
			wg.Done()
		}(randID, wg, rwMutex, ch2)

		go func(ch1 <-chan Book, ch2 <-chan Book){
			select{
				case b := <-ch1:
					fmt.Println("From cache")
					fmt.Println(b)
					<-ch2
				case b:= <-ch2:
					fmt.Println("From database")
					fmt.Println(b)
			}
		}(ch1, ch2)
		time.Sleep(150 * time.Millisecond)
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
