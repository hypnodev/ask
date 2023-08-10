package ask

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient(t *testing.T) {
	var post Post

	client := NewClient(context.Background())
	client.SetBaseUrl("https://jsonplaceholder.typicode.com")
	client.AddDefaultHeader("Accept", "application/json")
	client.SetVerbose(true)
	client.httpClient = &ClientMock{}

	SetClient(*client)
	res, err := GetJson("/posts/1", &post)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, "Test title", post.Title)
}

func TestClientPrintBody(t *testing.T) {
	var post Post
	payload := Post{
		Title:  "Test title",
		Body:   "Test body",
		UserId: 1,
	}
	data, _ := json.Marshal(payload)

	client := NewClient(context.Background())
	client.SetBaseUrl("https://jsonplaceholder.typicode.com")
	client.AddDefaultHeader("Accept", "application/json")
	client.SetVerbose(true)
	client.httpClient = &ClientMock{}

	SetClient(*client)
	res, err := PostJson("/posts/1", data, &post)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, "Test title", post.Title)
}
