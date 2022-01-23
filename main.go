package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	// start the waitgroup
	fmt.Println("Starting waitgroup")
	wg := sync.WaitGroup{}

	// declare and start a ticker in a background process
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for range ticker.C {
			fmt.Println("tock")
		}
	}()

	wg.Add(1)
	go someApiCall("A", 503, time.Second*1, &wg)

	wg.Add(1)
	go someApiCall("B", 503, time.Second*2, &wg)

	wg.Add(1)
	go someApiCall("C", 503, time.Second*4, &wg)

	// wait until completion or first occurring error
	wg.Wait()
	defer ticker.Stop()
}

func someApiCall(name string, statusCode int, d time.Duration, wg *sync.WaitGroup) http.Response {
	defer wg.Done()

	fmt.Println(fmt.Sprintf("Starting call %s", name))
	time.Sleep(d)
	fmt.Println(fmt.Sprintf("Call %s responding with status code %d after %v second(s)", name, statusCode, d.Seconds()))

	return http.Response{
		StatusCode: statusCode,
	}
}
