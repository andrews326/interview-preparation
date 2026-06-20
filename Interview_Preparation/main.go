package main

import (
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Result struct {
	URL        string
	StatusCode int
	Err        error
}

func main() {
	start := time.Now()
	urls := []string{
		"https://jsonplaceholder.typicode.com/users",
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://dummyjson.com/products",
		"https://dummyjson.com/users/1",
		"https://reqres.in/api/users?page=2",
		"https://fakestoreapi.com/products",
		"https://opentdb.com/api.php?amount=10",
		"https://restcountries.com/v3.1/all",
		"https://randomuser.me/api/",
		"https://pokeapi.co/api/v2/pokemon/ditto",
		"https://api.jikan.moe/v4/top/anime",
		"https://api.thecatapi.com/v1/images/search?limit=5",
	}

	count := 3
	workers := make(chan Result, len(urls))
	wg := sync.WaitGroup{}

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := http.Client{Timeout: 10 * time.Second}
			for url := range workers {
				result := fetchURL(client, url.URL)
				println("URL:", result.URL, "Status Code:", result.StatusCode, "Error:", result.Err)
			}
		}()
	}

	for _, url := range urls {
		workers <- Result{URL: url}
	}
	close(workers)

	wg.Wait()

	elapsed := time.Since(start)
	println("Elapsed time:", elapsed.String())
}

func urlMustParse(rawurl string) *url.URL {
	parsedURL, err := url.Parse(rawurl)
	if err != nil {
		panic(err)
	}
	return parsedURL
}

func fetchURL(client http.Client, url string) Result {
	req := http.Request{
		Method: http.MethodGet,
		URL:    urlMustParse(url),
	}
	resp, err := client.Do(&req)
	if err != nil {
		return Result{URL: url, Err: err}
	}
	resp.Body.Close()
	return Result{URL: url, StatusCode: resp.StatusCode}
}
