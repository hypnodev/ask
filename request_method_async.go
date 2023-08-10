package ask

import (
	"net/http"
	"os"
)

func GetJsonAsync(url string, v any, res chan Response, error chan error) {
	request := NewRequest(http.MethodGet, url)
	request.setClient(&client)
	response, err := request.AcceptJson().Send()
	if err != nil {
		error <- err
		return
	}

	if response.body != nil {
		err = marshalResponseIfStruct(response.body, &v)
		if err != nil {
			error <- err
			return
		}
	}

	res <- *response
	error <- nil
}

func PostJsonAsync(url string, payload []byte, v any, res chan Response, error chan error) {
	request := NewRequest(http.MethodPost, url)
	request.setClient(&client)
	request.WithPayloadJson(payload)
	response, err := request.AcceptJson().Send()
	if err != nil {
		error <- err
		return
	}

	if response.body != nil {
		err = marshalResponseIfStruct(response.body, &v)
		if err != nil {
			error <- err
			return
		}
	}

	res <- *response
	error <- nil
}

func PutJsonAsync(url string, payload []byte, v any, res chan Response, error chan error) {
	request := NewRequest(http.MethodPut, url)
	request.setClient(&client)
	request.WithPayloadJson(payload)
	response, err := request.AcceptJson().Send()
	if err != nil {
		error <- err
		return
	}

	if response.body != nil {
		err = marshalResponseIfStruct(response.body, &v)
		if err != nil {
			error <- err
			return
		}
	}

	res <- *response
	error <- nil
}

func PatchJsonAsync(url string, payload []byte, v any, res chan Response, error chan error) {
	request := NewRequest(http.MethodPatch, url)
	request.setClient(&client)
	request.WithPayloadJson(payload)
	response, err := request.AcceptJson().Send()
	if err != nil {
		error <- err
		return
	}

	if response.body != nil {
		err = marshalResponseIfStruct(response.body, &v)
		if err != nil {
			error <- err
			return
		}
	}

	res <- *response
	error <- nil
}

func DeleteJsonAsync(url string, payload *[]byte, v any, res chan Response, error chan error) {
	request := NewRequest(http.MethodDelete, url)
	request.setClient(&client)
	if payload != nil {
		request.WithPayloadJson(*payload)
	}
	response, err := request.AcceptJson().Send()
	if err != nil {
		error <- err
		return
	}

	if response.body != nil && v != nil {
		err = marshalResponseIfStruct(response.body, &v)
		if err != nil {
			error <- err
			return
		}
	}

	res <- *response
	error <- nil
}

func GetFileAsync(url string, dest string, error chan error) {
	request := NewRequest(http.MethodGet, url)
	request.setClient(&client)
	response, err := request.Send()
	if err != nil {
		error <- err
		return
	}

	err = os.WriteFile(dest, response.body, 0644)
	if err != nil {
		error <- err
		return
	}

	error <- nil
}
