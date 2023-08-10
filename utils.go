package ask

import "encoding/json"

func marshalResponseIfStruct(response []byte, parsedBody any) error {
	switch parsedBody.(type) {
	case []byte:
		parsedBody = response
	default:
		err := json.Unmarshal(response, &parsedBody)
		if err != nil {
			return err
		}
	}

	return nil
}
