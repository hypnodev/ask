package ask

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetJson(t *testing.T) {
	var post Post

	SetClient(mockClient(nil))
	_, err := GetJson("https://jsonplaceholder.typicode.com/posts/1", &post)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Test title", post.Title)
}

func TestGetJsonError(t *testing.T) {
	var post Post
	e := &ResponseError{StatusCode: 404, Err: errors.New("not found")}

	SetClient(mockClient(e))
	response, err := GetJson("https://jsonplaceholder.typicode.com/posts/1", &post)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEqual(t, "Test title", post.Title)
	assert.Equal(t, 404, response.StatusCode)
	assert.Equal(t, "not found", response.Error)
}

func TestPostJson(t *testing.T) {
	var post Post
	payload := Post{
		Title:  "Test title",
		Body:   "Test body",
		UserId: 1,
	}
	data, _ := json.Marshal(payload)

	SetClient(mockClient(nil))
	_, err := PostJson("https://jsonplaceholder.typicode.com/posts", data, &post)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Test title", post.Title)
	assert.Equal(t, 1, post.Id)
}

func TestPostJsonError(t *testing.T) {
	var post Post
	payload := Post{
		Title:  "",
		Body:   "Test body",
		UserId: 1,
	}
	data, _ := json.Marshal(payload)

	unprocessableEntityError := map[string]interface{}{
		"Title": "Min 4 length",
	}
	errorData, _ := json.Marshal(unprocessableEntityError)
	e := &ResponseError{StatusCode: 422, Err: errors.New(string(errorData))}

	SetClient(mockClient(e))
	response, err := PostJson("https://jsonplaceholder.typicode.com/posts", data, &post)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 422, response.StatusCode)
	assert.Equal(t, unprocessableEntityError, response.Error)
}

func TestPatchJson(t *testing.T) {
	var post Post
	payload := Post{
		Body: "Edited body",
	}
	data, _ := json.Marshal(payload)

	SetClient(mockClient(nil))
	_, err := PatchJson("https://jsonplaceholder.typicode.com/posts/1", data, &post)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Test title", post.Title)
	assert.Equal(t, "Edited body", post.Body)
}

func TestPatchJsonError(t *testing.T) {
	var post Post
	payload := Post{
		Title: "nop",
	}
	data, _ := json.Marshal(payload)

	unprocessableEntityError := map[string]interface{}{
		"Title": "Min 4 length",
	}
	errorData, _ := json.Marshal(unprocessableEntityError)
	e := &ResponseError{StatusCode: 422, Err: errors.New(string(errorData))}

	SetClient(mockClient(e))
	response, err := PatchJson("https://jsonplaceholder.typicode.com/posts/1", data, &post)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEqual(t, "Test title", post.Title)
	assert.Equal(t, 422, response.StatusCode)
	assert.Equal(t, unprocessableEntityError, response.Error)
}

func TestPutJson(t *testing.T) {
	var post Post
	payload := Post{
		Title:  "Edited title",
		Body:   "Edited body",
		UserId: 2,
	}
	data, _ := json.Marshal(payload)

	SetClient(mockClient(nil))
	_, err := PutJson("https://jsonplaceholder.typicode.com/posts/1", data, &post)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, post.Id)
	assert.Equal(t, "Edited title", post.Title)
	assert.Equal(t, "Edited body", post.Body)
	assert.Equal(t, 2, post.UserId)
}

func TestPutJsonError(t *testing.T) {
	var post Post
	payload := Post{
		Title:  "nop",
		Body:   "Test body",
		UserId: 1,
	}
	data, _ := json.Marshal(payload)

	unprocessableEntityError := map[string]interface{}{
		"Title": "Min 4 length",
	}
	errorData, _ := json.Marshal(unprocessableEntityError)
	e := &ResponseError{StatusCode: 422, Err: errors.New(string(errorData))}

	SetClient(mockClient(e))
	response, err := PatchJson("https://jsonplaceholder.typicode.com/posts/1", data, &post)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEqual(t, "Test title", post.Title)
	assert.Equal(t, 422, response.StatusCode)
	assert.Equal(t, unprocessableEntityError, response.Error)
}

func TestDeleteJson(t *testing.T) {
	SetClient(mockClient(nil))
	response, err := DeleteJson("https://jsonplaceholder.typicode.com/posts/1", nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 204, response.StatusCode)
	assert.True(t, nil == response.GetBody())
}

func TestDeleteJsonError(t *testing.T) {
	e := &ResponseError{StatusCode: 404, Err: errors.New("not found")}

	SetClient(mockClient(e))
	response, err := DeleteJson("https://jsonplaceholder.typicode.com/posts/1", nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 404, response.StatusCode)
	assert.Equal(t, "not found", response.Error)
}
