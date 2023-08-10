package ask

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

type QueryParams map[string]any

type Request struct {
	client  *Client
	method  string
	url     *url.URL
	Header  http.Header
	payload *bytes.Buffer
}

func NewRequest(method string, requestUrl string) *Request {
	client := NewClient(context.Background())

	parsedUrl, err := url.Parse(requestUrl)
	if err != nil {
		log.Panicln(err)
	}

	return &Request{
		client:  client,
		method:  method,
		url:     parsedUrl,
		Header:  http.Header{},
		payload: nil,
	}
}

func (request *Request) setClient(client *Client) *Request {
	if client != nil {
		request.client = client
	} else {
		request.client = NewClient(context.Background())
	}
	return request
}

func (request *Request) WithPayloadJson(json []byte) *Request {
	request.Header.Set("Content-Type", "application/json")
	request.payload = bytes.NewBuffer(json)
	return request
}

func (request *Request) AcceptJson() *Request {
	request.Header.Set("Accept", "application/json")
	return request
}

func (request *Request) SendRaw() (*http.Response, error) {
	var req *http.Request
	var err error
	if request.payload != nil {
		req, err = http.NewRequest(request.method, request.url.String(), request.payload)
	} else {
		req, err = http.NewRequest(request.method, request.url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	req.Header = request.Header
	for k, v := range request.client.defaultHeaders {
		req.Header.Set(k, v[0])
	}

	if request.client.httpClient == nil {
		request.setClient(nil)
	}

	if len(request.client.baseUrl) > 0 {
		parsedUrl, err := url.Parse(request.client.baseUrl + request.url.String())
		if err != nil {
			log.Panicln(err)
		}
		req.URL = parsedUrl
	}

	response, err := request.client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if request.client.verbose {
		reqHeader, err := json.Marshal(req.Header)
		if err != nil {
			return nil, err
		}

		resHeader, err := json.Marshal(response.Header)
		if err != nil {
			return nil, err
		}

		var reqBody []byte
		if req.Body != nil {
			reqBody, err = io.ReadAll(req.Body)
			if err != nil {
				return nil, err
			}
		}

		log.Println("Request sent to " + req.URL.String())
		log.Println(string(reqHeader))
		log.Print("Payload: " + string(reqBody))
		log.Println("Response from " + req.URL.String())
		log.Println(string(resHeader))
	}

	return response, nil
}

func (request *Request) Send() (*Response, error) {
	response, err := request.SendRaw()
	if response.Body == nil {
		return &Response{StatusCode: response.StatusCode}, nil
	}
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if request.client.verbose {
		log.Println(string(body))
	}

	res := &Response{StatusCode: response.StatusCode}

	if response.StatusCode >= 200 && response.StatusCode < 300 {
		res.body = body
	} else {
		strBody := string(body)
		if (strBody[0] == '{' && strBody[len(strBody)-1] == '}') || (strBody[0] == '[' && strBody[len(strBody)-1] == ']') {
			err = json.Unmarshal(body, &res.Error)
			if err != nil {
				return nil, err
			}
		} else {
			res.Error = strBody
		}
	}

	return res, nil
}

func (request *Request) SetForm(payload map[string]string) (*Request, error) {
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.payload = bytes.NewBuffer(b)
	return request, nil
}
