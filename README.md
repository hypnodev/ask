## Ask

An HTTP client built on top of `net/http` with support of coroutines and JSON deserialization by default.

### Installation

```shell
$ go get github.com/hypnodev/ask
```

### Usage

```go
package main

import (
	"context"
	"encoding/json"
	"github.com/hypnodev/ask"
	"log"
)

type Post struct {
	Id     int
	Title  string
	Body   string
	UserId int `json:"userId"`
}

func main() {
	log.Println("Call an endpoint")
	callEndpoint()

	log.Println("Call an endpoint async")
	callEndpointAsync()

	log.Println("Call an endpoint with client configuration")
	callEndpointWithClient()

	log.Println("Download a file")
	downloadFile()

	log.Println("Send a form")
	sendForm()
}

func callEndpoint() {
	var post Post
	res, err := ask.GetJson("https://jsonplaceholder.typicode.com/posts/1", &post)
	if err != nil {
		log.Panicln(err)
	}

	log.Println(res.StatusCode, post)
}

func callEndpointAsync() {
	resChan := make(chan ask.Response, 1)
	errorChan := make(chan error, 1)

	var post Post
	go ask.GetJsonAsync("https://jsonplaceholder.typicode.com/posts/1", &post, resChan, errorChan)

	err := <-errorChan
	if err != nil {
		log.Panicln(err)
	}

	res := <-resChan
	log.Println(res, post)
}

func callEndpointWithClient() {
	client := ask.NewClient(context.Background())
	client.SetBaseUrl("https://jsonplaceholder.typicode.com")
	client.AddDefaultHeader("Accept", "application/json")
	client.SetVerbose(true)

	var post Post
	payload := Post{
		Title:  "title",
		Body:   "test body",
		UserId: 1,
	}
	data, _ := json.Marshal(payload)

	ask.SetClient(*client)
	res, err := ask.PostJson("/posts", data, &post)
	if err != nil {
		log.Panicln(err)
	}

	log.Println(res.StatusCode, post)
}

func downloadFile() {
	url := "https://picsum.photos/200/300"

	err := ask.GetFile(url, "test.jpg")
	if err != nil {
		log.Panicln(err)
	}
}

func sendForm() {
	var post Post
	payload := map[string]string{
		"title":  "title",
		"body":   "test body",
		"userId": "1",
	}

	res, err := ask.PostForm("https://jsonplaceholder.typicode.com/posts/1", payload, &post)
	if err != nil {
		log.Panicln(err)
	}

	log.Println(res.StatusCode, post)
}
```

## Contribute
Feel free to push any changes, improve anything and fix stuff.  

## Stay in touch
- Author - [Cristian Cosenza](https://linkedin.com/in/cristiancosenza)
