package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
)

const WORKERS = 5
const URL = "https://rosmetallica.ru/"

func main() {
	maxRequests := 1000

	var successRequestCounter int
	var requestCounter int
	mu := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	for i := 0; i < WORKERS; i++ {
		wg.Add(1)
		go func(mu *sync.Mutex, i int) {
			quit := make(chan int)
			defer wg.Done()
			for {
				select {
				case <-quit:
					return
				default:
					if requestCounter < maxRequests {
						_, err := http.Get(URL)
						mu.Lock()
						requestCounter++
						if err == nil {
							successRequestCounter++
						}
						mu.Unlock()
						fmt.Printf("worker %d send %d request\n", i, requestCounter)
					} else {
						go func() {
							quit <- 1
						}()
					}
				}
				runtime.Gosched()
			}
		}(mu, i)
	}
	wg.Wait()

	defer fmt.Printf("amount requests %d\n", requestCounter)
	defer fmt.Printf("success requests %d\n", successRequestCounter)

}
