package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const WORKERS = 16

const URL = "http://localhost:8080/ping/"

//const URL = "https://rosmetallica.ru/"

func main() {
	start := time.Now()
	maxRequests := 10000

	var successRequestCounter int
	var requestCounter int
	mu := new(sync.Mutex)
	quit := make(chan int)
	run := make(chan bool)
	for i := 0; i < WORKERS; i++ {
		go func(i int) {
			for {
				select {
				case <-quit:
					return
				case <-run:
					_, err := http.Get(URL)
					mu.Lock()
					requestCounter++
					if err == nil {
						successRequestCounter++
					}
					fmt.Printf("worker %d send %d request\n", i, requestCounter)
					mu.Unlock()
				}
			}
		}(i)
	}

	checkTime := time.After(1 * time.Second)
	for k := maxRequests; k > 0; k-- {
		stop := false
		select {
		case <-checkTime:
			stop = true
		default:
			run <- true
		}
		if stop {
			break
		}
	}
	for j := 0; j < WORKERS; j++ {
		quit <- 1
	}

	dur := time.Since(start)
	fmt.Println("Время:", dur)
	fmt.Printf("amount requests %d\n", requestCounter)
	fmt.Printf("success requests %d\n", successRequestCounter)

}
