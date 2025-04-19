package main

import (
    "fmt"
    "net/http"
    "time"
)

// fetchWithRetry attempts to fetch the URL with retries.
func fetchWithRetry(url string, retries int) (*http.Response, error) {
    var resp *http.Response
    var err error

    for i := 0; i < retries; i++ {
        resp, err = http.Get(url)
        if err == nil {
            return resp, nil
		}
        time.Sleep(time.Second * time.Duration(i+1)) // Exponential...technically 'linear'backoff.
		fmt.Println("Retrying...")
    }
    return nil, err
}

func main() {
    url := "https://example.com"
    resp, err := fetchWithRetry(url, 5)
    if err != nil {
        fmt.Println("Failed to fetch URL:", err)
        return
    }
    defer resp.Body.Close()
    fmt.Println("Response status:", resp.Status)
}