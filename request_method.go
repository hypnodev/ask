package ask

import (
	"net/http"
	"os"
)

func GetJson(url string, v any) (*Response, error) {
	request := NewRequest(http.MethodGet, url)
	request.setClient(&client)
	response, err := request.AcceptJson().Send()
	if err != nil {
		return nil, err
	}

	if response.body != nil {
		err = marshalResponseIfStruct(response.body, &v)
		if err != nil {
			return nil, err
		}
	}

	return response, nil
}

func PostJson(url string, payload []byte, v any) (*Response, error) {
	request := NewRequest(http.MethodPost, url)
	request.setClient(&client)
	request.WithPayloadJson(payload)
	response, err := request.AcceptJson().Send()
	if err != nil {
		return nil, err
	}

	if response.body != nil {
		err = marshalResponseIfStruct(response.body, &v)
		if err != nil {
			return nil, err
		}
	}

	return response, nil
}

func PutJson(url string, payload []byte, v any) (*Response, error) {
	request := NewRequest(http.MethodPut, url)
	request.setClient(&client)
	request.WithPayloadJson(payload)
	response, err := request.AcceptJson().Send()
	if err != nil {
		return nil, err
	}

	if response.body != nil {
		err = marshalResponseIfStruct(response.body, &v)
		if err != nil {
			return nil, err
		}
	}

	return response, nil
}

func PatchJson(url string, payload []byte, v any) (*Response, error) {
	request := NewRequest(http.MethodPatch, url)
	request.setClient(&client)
	request.WithPayloadJson(payload)
	response, err := request.AcceptJson().Send()
	if err != nil {
		return nil, err
	}

	if response.body != nil {
		err = marshalResponseIfStruct(response.body, &v)
		if err != nil {
			return nil, err
		}
	}

	return response, nil
}

func DeleteJson(url string, payload *[]byte, v any) (*Response, error) {
	request := NewRequest(http.MethodDelete, url)
	request.setClient(&client)
	if payload != nil {
		request.WithPayloadJson(*payload)
	}
	response, err := request.AcceptJson().Send()
	if err != nil {
		return nil, err
	}

	if response.body != nil && v != nil {
		err = marshalResponseIfStruct(response.body, &v)
		if err != nil {
			return nil, err
		}
	}

	return response, nil
}

func GetFile(url string, dest string) error {
	request := NewRequest(http.MethodGet, url)
	request.setClient(&client)
	response, err := request.Send()
	if err != nil {
		return err
	}

	err = os.WriteFile(dest, response.body, 0644)
	if err != nil {
		return err
	}

	return nil
}

func PostForm(url string, payload map[string]string, v any) (*Response, error) {
	request := NewRequest(http.MethodPost, url)
	request.setClient(&client)

	_, err := request.SetForm(payload)
	if err != nil {
		return nil, err
	}

	response, err := request.AcceptJson().Send()
	if err != nil {
		return nil, err
	}

	if response.body != nil {
		err = marshalResponseIfStruct(response.body, &v)
		if err != nil {
			return nil, err
		}
	}

	return response, nil
}

func PutForm(url string, payload map[string]string, v any) (*Response, error) {
	request := NewRequest(http.MethodPut, url)
	request.setClient(&client)

	_, err := request.SetForm(payload)
	if err != nil {
		return nil, err
	}

	response, err := request.AcceptJson().Send()
	if err != nil {
		return nil, err
	}

	if response.body != nil {
		err = marshalResponseIfStruct(response.body, &v)
		if err != nil {
			return nil, err
		}
	}

	return response, nil
}

func PatchForm(url string, payload map[string]string, v any) (*Response, error) {
	request := NewRequest(http.MethodPatch, url)
	request.setClient(&client)

	_, err := request.SetForm(payload)
	if err != nil {
		return nil, err
	}

	response, err := request.AcceptJson().Send()
	if err != nil {
		return nil, err
	}

	if response.body != nil {
		err = marshalResponseIfStruct(response.body, &v)
		if err != nil {
			return nil, err
		}
	}

	return response, nil
}
