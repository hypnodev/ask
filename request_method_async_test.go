package ask

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetJsonAsync(t *testing.T) {
	var post Post
	res := make(chan Response, 1)
	err := make(chan error, 1)

	SetClient(mockClient(nil))
	go GetJsonAsync("https://jsonplaceholder.typicode.com/posts/1", &post, res, err)
	e := <-err
	if e != nil {
		t.Fatal(e)
	}

	r := <-res
	assert.Equal(t, "Test title", post.Title)
	assert.Equal(t, 200, r.StatusCode)
}

func TestGetJsonAsyncError(t *testing.T) {
	var post Post
	res := make(chan Response, 1)
	err := make(chan error, 1)
	resError := &ResponseError{StatusCode: 404, Err: errors.New("not found")}

	SetClient(mockClient(resError))
	go GetJsonAsync("https://jsonplaceholder.typicode.com/posts/1", &post, res, err)
	e := <-err
	if e != nil {
		t.Fatal(e)
	}

	r := <-res
	assert.NotEqual(t, "Test title", post.Title)
	assert.Equal(t, 404, r.StatusCode)
	assert.Equal(t, "not found", r.Error)
}

func TestPostJsonAsync(t *testing.T) {
	var post Post
	payload := Post{
		Title:  "Test title",
		Body:   "Test body",
		UserId: 1,
	}
	data, _ := json.Marshal(payload)
	res := make(chan Response, 1)
	err := make(chan error, 1)

	SetClient(mockClient(nil))
	go PostJsonAsync("https://jsonplaceholder.typicode.com/posts/1", data, &post, res, err)
	e := <-err
	if e != nil {
		t.Fatal(e)
	}

	assert.Equal(t, "Test title", post.Title)
	assert.Equal(t, 1, post.Id)
}

func TestPostJsonAsyncError(t *testing.T) {
	var post Post
	payload := Post{
		Title:  "Test title",
		Body:   "Test body",
		UserId: 1,
	}
	data, _ := json.Marshal(payload)
	res := make(chan Response, 1)
	err := make(chan error, 1)

	unprocessableEntityError := map[string]interface{}{
		"Title": "Min 4 length",
	}
	errorData, _ := json.Marshal(unprocessableEntityError)
	resError := &ResponseError{StatusCode: 422, Err: errors.New(string(errorData))}

	SetClient(mockClient(resError))
	go PostJsonAsync("https://jsonplaceholder.typicode.com/posts/1", data, &post, res, err)
	e := <-err
	if e != nil {
		t.Fatal(e)
	}

	r := <-res
	assert.Equal(t, 422, r.StatusCode)
	assert.Equal(t, unprocessableEntityError, r.Error)
}

func TestPatchJsonAsync(t *testing.T) {
	var post Post
	payload := Post{
		Title: "Edited title",
	}
	data, _ := json.Marshal(payload)
	res := make(chan Response, 1)
	err := make(chan error, 1)

	SetClient(mockClient(nil))
	go PatchJsonAsync("https://jsonplaceholder.typicode.com/posts/1", data, &post, res, err)
	e := <-err
	if e != nil {
		t.Fatal(e)
	}

	assert.Equal(t, "Edited title", post.Title)
	assert.Equal(t, 1, post.Id)
}

func TestPatchJsonAsyncError(t *testing.T) {
	var post Post
	payload := Post{
		Title: "nop",
	}
	data, _ := json.Marshal(payload)
	res := make(chan Response, 1)
	err := make(chan error, 1)

	unprocessableEntityError := map[string]interface{}{
		"Title": "Min 4 length",
	}
	errorData, _ := json.Marshal(unprocessableEntityError)
	resError := &ResponseError{StatusCode: 422, Err: errors.New(string(errorData))}

	SetClient(mockClient(resError))
	go PatchJsonAsync("https://jsonplaceholder.typicode.com/posts/1", data, &post, res, err)
	e := <-err
	if e != nil {
		t.Fatal(e)
	}

	r := <-res
	assert.Equal(t, 422, r.StatusCode)
	assert.Equal(t, unprocessableEntityError, r.Error)
}

func TestPutJsonAsync(t *testing.T) {
	var post Post
	payload := Post{
		Title:  "Edited title",
		Body:   "Edited body",
		UserId: 2,
	}
	data, _ := json.Marshal(payload)
	res := make(chan Response, 1)
	err := make(chan error, 1)

	SetClient(mockClient(nil))
	go PutJsonAsync("https://jsonplaceholder.typicode.com/posts/1", data, &post, res, err)
	e := <-err
	if e != nil {
		t.Fatal(e)
	}

	assert.Equal(t, "Edited title", post.Title)
	assert.Equal(t, 1, post.Id)
}

func TestPutJsonAsyncError(t *testing.T) {
	var post Post
	payload := Post{
		Title:  "nop",
		Body:   "Edited body",
		UserId: 2,
	}
	data, _ := json.Marshal(payload)
	res := make(chan Response, 1)
	err := make(chan error, 1)

	unprocessableEntityError := map[string]interface{}{
		"Title": "Min 4 length",
	}
	errorData, _ := json.Marshal(unprocessableEntityError)
	resError := &ResponseError{StatusCode: 422, Err: errors.New(string(errorData))}

	SetClient(mockClient(resError))
	go PutJsonAsync("https://jsonplaceholder.typicode.com/posts/1", data, &post, res, err)
	e := <-err
	if e != nil {
		t.Fatal(e)
	}

	r := <-res
	assert.Equal(t, 422, r.StatusCode)
	assert.Equal(t, unprocessableEntityError, r.Error)
}

func TestDeleteJsonAsync(t *testing.T) {
	res := make(chan Response, 1)
	err := make(chan error, 1)

	SetClient(mockClient(nil))
	go DeleteJsonAsync("https://jsonplaceholder.typicode.com/posts/1", nil, nil, res, err)
	e := <-err
	if e != nil {
		t.Fatal(e)
	}

	r := <-res
	assert.Equal(t, 204, r.StatusCode)
	assert.True(t, nil == r.GetBody())
}

func TestDeleteJsonAsyncError(t *testing.T) {
	res := make(chan Response, 1)
	err := make(chan error, 1)

	resError := &ResponseError{StatusCode: 404, Err: errors.New("not found")}

	SetClient(mockClient(resError))
	go DeleteJsonAsync("https://jsonplaceholder.typicode.com/posts/1", nil, nil, res, err)
	e := <-err
	if e != nil {
		t.Fatal(e)
	}

	r := <-res
	assert.Equal(t, 404, r.StatusCode)
	assert.Equal(t, "not found", r.Error)
}
