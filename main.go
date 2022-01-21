package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type call struct {
	Duration   time.Duration
	StatusCode int
}

func main() {
	calls := map[string]call{
		"A": {
			Duration:   time.Second * 1,
			StatusCode: 503,
		},
		"B": {
			Duration:   time.Second * 2,
			StatusCode: 503,
		},
		"C": {
			Duration:   time.Second * 4,
			StatusCode: 503,
		},
	}

	fmt.Println("Starting waitgroup")
	wg := sync.WaitGroup{}

	for k, v := range calls {
		wg.Add(1)
		go someApiCall(k, v.StatusCode, v.Duration, &wg)
	}

	// wait until completion or first occurring error
	wg.Wait()
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
