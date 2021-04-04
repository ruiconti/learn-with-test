package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("vim-go")
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

var TimeoutError = func(a, b string) error {
	errMsg := fmt.Sprintf("Timed out waiting for %s and %s to return", a, b)
	return errors.New(errMsg)
}

var defaultTimeout = 10 * time.Second

func Racer(urlA, urlB string) (winner string, err error) {
	return ConfigurableRacer(urlA, urlB, defaultTimeout)
}

func ConfigurableRacer(urlA, urlB string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(urlA):
		return urlA, nil
	case <-ping(urlB):
		return urlB, nil
	case <-time.After(timeout):
		return "", TimeoutError(urlA, urlB)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
