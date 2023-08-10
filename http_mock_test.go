package ask

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type Post struct {
	Id     int    `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Body   string `json:"body,omitempty"`
	UserId int    `json:"userId"`
}

type ResponseError struct {
	StatusCode int
	Err        error
}

var err *ResponseError

// ClientMock is a mocked Client interface
// It is used to mock http responses
type ClientMock struct {
}

func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	if err != nil {
		e := err
		err = nil

		r := strings.NewReader(e.Err.Error())
		body := io.NopCloser(r)
		return &http.Response{StatusCode: e.StatusCode, Body: body}, nil
	}

	if req.Method == http.MethodDelete {
		return &http.Response{StatusCode: 204, Body: nil}, nil
	}

	post := Post{
		Id:     1,
		Title:  "Test title",
		Body:   "Test body",
		UserId: 1,
	}
	if req.URL.Path == "/posts/1" && (req.Method == http.MethodPatch || req.Method == http.MethodPut) {
		_ = json.NewDecoder(req.Body).Decode(&post)
	}
	data, _ := json.Marshal(post)

	r := bytes.NewReader(data)
	return &http.Response{StatusCode: 200, Body: io.NopCloser(r)}, nil
}

func mockClient(e *ResponseError) Client {
	client := NewClient(context.Background())
	client.httpClient = &ClientMock{}
	err = e

	return *client
}
