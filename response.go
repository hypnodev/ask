package ask

type Response struct {
	body       []byte
	StatusCode int
	Error      interface{}
}

func (response Response) GetBody() *[]byte {
	if len(response.body) == 0 {
		return nil
	}

	return &response.body
}
