package main

import (
	"fmt"
	"net/http"
	"sync"
)

const WORKERS = 5
const URL = "http://localhost:8080/ping/"

func main() {
	maxRequests := 100

	var successRequestCounter int
	var requestCounter int
	mu := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	quit := make(chan int)
	for i := 0; i < WORKERS; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case <-quit:
					return
				default:
					mu.Lock()
					if requestCounter >= maxRequests {
						go func() {
							for j := 0; j < WORKERS; j++ {
								quit <- 1
							}
						}()
					} else {
						_, err := http.Get(URL)
						requestCounter++
						if err == nil {
							successRequestCounter++
						}
						fmt.Printf("worker %d send %d request\n", i, requestCounter)
					}
					mu.Unlock()
				}
			}
		}(i)
	}
	wg.Wait()

	defer fmt.Printf("amount requests %d\n", requestCounter)
	defer fmt.Printf("success requests %d\n", successRequestCounter)

}
